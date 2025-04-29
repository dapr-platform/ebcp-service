#!/usr/bin/env python3
import json
import uuid
import requests
from datetime import datetime

def get_holiday_data(year):
    url = f"https://timor.tech/api/holiday/year/{year}/"
    
    # 模拟浏览器请求头
    headers = {
        "authority": "timor.tech",
        "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
        "accept-encoding": "gzip, deflate, br, zstd",
        "accept-language": "zh-CN,zh;q=0.9",
        "cache-control": "no-cache",
        "sec-ch-ua": '"Google Chrome";v="135", "Not-A.Brand";v="8", "Chromium";v="135"',
        "sec-ch-ua-mobile": "?0",
        "sec-ch-ua-platform": '"macOS"',
        "sec-fetch-dest": "document",
        "sec-fetch-mode": "navigate",
        "sec-fetch-site": "none",
        "sec-fetch-user": "?1",
        "upgrade-insecure-requests": "1",
        "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36"
    }

    try:
        response = requests.get(url, headers=headers)
        # 打印响应状态和内容，用于调试
        print(f"Response status code: {response.status_code}")
        
        if response.status_code == 304:
            print("收到304响应，重试请求...")
            # 移除缓存相关的头部，重新请求
            headers.pop("cache-control", None)
            response = requests.get(url, headers=headers)
            print(f"重试后状态码: {response.status_code}")
        
        # 检查响应状态码
        response.raise_for_status()
        
        # 如果响应内容为空，返回None
        if not response.text.strip():
            print("响应内容为空")
            return None
            
        print(f"Response content: {response.text[:200]}...")  # 只打印前200个字符
        
        # 尝试解析JSON
        return response.json()
    except requests.exceptions.RequestException as e:
        print(f"请求失败: {e}")
        return None
    except json.JSONDecodeError as e:
        print(f"JSON解析失败: {e}")
        print(f"响应内容: {response.text}")
        return None

def determine_holiday_type(holiday_info):
    """
    确定假日类型：
    1: 法定节假日
    2: 调休工作日
    3: 周末调休
    4: 闭馆日（暂不使用）
    """
    if holiday_info["holiday"]:
        return 1  # 法定节假日
    elif "target" in holiday_info:
        if holiday_info.get("after", False):
            return 2  # 调休工作日（节后补班）
        else:
            return 3  # 周末调休（节前补班）
    return 3  # 默认为周末调休

def generate_sql(year):
    data = get_holiday_data(year)
    if data is None:
        print(f"获取{year}年节假日数据失败")
        return ""
    
    if not isinstance(data, dict) or "code" not in data or "holiday" not in data:
        print(f"数据格式错误: {data}")
        return ""

    if data["code"] != 0:
        print(f"API返回错误代码: {data['code']}")
        return ""

    sql_lines = []
    sql_lines.append("-- +goose Up")
    sql_lines.append("-- +goose StatementBegin")
    sql_lines.append("")
    
    # 系统用户ID
    system_user = "admin"
    current_time = datetime.now().strftime("%Y-%m-%d %H:%M:%S")

    try:
        for date_str, holiday_info in data["holiday"].items():
            # 生成UUID
            id = str(uuid.uuid4()).replace("-", "")
            
            # 确定假日类型
            holiday_type = determine_holiday_type(holiday_info)
            
            # 生成备注信息
            remarks = "节假日" if holiday_info["holiday"] else (
                f"{holiday_info['target']}{'后' if holiday_info.get('after', False) else '前'}补班"
            )
            
            # 构建SQL插入语句
            sql = f"""INSERT INTO o_ebcp_holiday_date (
    id, created_by, created_time, updated_by, updated_time,
    date, name, type, year, remarks
) VALUES (
    '{id}', '{system_user}', '{current_time}', '{system_user}', '{current_time}',
    '{holiday_info["date"]} 00:00:00', '{holiday_info["name"]}', {holiday_type}, {year},
    '{remarks}'
);"""
            sql_lines.append(sql)

    except Exception as e:
        print(f"生成SQL时发生错误: {e}")
        return ""

    sql_lines.append("")
    sql_lines.append("-- +goose StatementEnd")
    sql_lines.append("")
    sql_lines.append("-- +goose Down")
    sql_lines.append("-- +goose StatementBegin")
    sql_lines.append(f"DELETE FROM o_ebcp_holiday_date WHERE year = {year};")
    sql_lines.append("")
    sql_lines.append("-- +goose StatementEnd")
    
    return "\n".join(sql_lines)

def main():
    year = 2025
    sql_content = generate_sql(year)
    
    if sql_content:
        # 写入SQL文件
        try:
            with open("101_ebcp_holiday_init.sql", "w", encoding="utf-8") as f:
                f.write(sql_content)
            print("SQL文件生成成功！")
        except Exception as e:
            print(f"写入文件时发生错误: {e}")
    else:
        print("由于错误，未生成SQL文件")

if __name__ == "__main__":
    main()
