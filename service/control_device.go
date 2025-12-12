package service

import (
	"context"
	"ebcp-service/model"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/dapr-platform/common"
)

func ControlDeviceCommand(deviceIP string, devicePort int32, command string, cmdType string, deviceId string, deviceType string) error {

	// 验证输入参数
	if deviceIP == "" {
		return fmt.Errorf("设备IP地址不能为空")
	}
	if devicePort <= 0 || devicePort > 65535 {
		return fmt.Errorf("设备端口号无效: %d", devicePort)
	}
	if command == "" {
		return fmt.Errorf("控制命令不能为空")
	}
	if cmdType == "" {
		return fmt.Errorf("控制命令类型不能为空")
	}
	if deviceId == "" {
		return fmt.Errorf("设备ID不能为空")
	}
	if deviceType == "" {
		return fmt.Errorf("设备类型不能为空")
	}

	// 发送UDP命令到设备
	if err := SendUDPCommand(deviceIP, devicePort, command); err != nil {
		return fmt.Errorf("向设备 %s:%d 发送命令失败: %v", deviceIP, devicePort, err)
	}

	if deviceType == "control_device" {
		// 查询中控设备
		controlDevice, err := common.DbGetOne[model.Ebcp_control_device](context.Background(), common.GetDaprClient(), model.Ebcp_control_deviceTableInfo.Name, model.Ebcp_control_device_FIELD_NAME_id+"="+deviceId)
		if err != nil {
			return fmt.Errorf("查询中控设备失败: %v", err)
		}
		if cmdType == "start" {
			// 启动中控设备
			controlDevice.Status = ItemStatusStart
			controlDevice.UpdatedTime = common.LocalTime(time.Now())
			common.DbUpsert[model.Ebcp_control_device](context.Background(), common.GetDaprClient(), *controlDevice, model.Ebcp_control_deviceTableInfo.Name, "id")
		} else if cmdType == "stop" {
			// 停止中控设备
			controlDevice.Status = ItemStatusStop
			controlDevice.UpdatedTime = common.LocalTime(time.Now())
			common.DbUpsert[model.Ebcp_control_device](context.Background(), common.GetDaprClient(), *controlDevice, model.Ebcp_control_deviceTableInfo.Name, "id")
		} else if cmdType == "pause" {
			// 暂停中控设备
			controlDevice.Status = ItemStatusPause
			controlDevice.UpdatedTime = common.LocalTime(time.Now())
			common.DbUpsert[model.Ebcp_control_device](context.Background(), common.GetDaprClient(), *controlDevice, model.Ebcp_control_deviceTableInfo.Name, "id")
		}
	} else if deviceType == "item" {
		// 查询展项
		exhibitionItem, err := common.DbGetOne[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(), model.Ebcp_exhibition_itemTableInfo.Name, model.Ebcp_exhibition_item_FIELD_NAME_id+"="+deviceId)
		if err != nil {
			return fmt.Errorf("查询展项失败: %v", err)
		}
		controlDevices, err := common.DbQuery[model.Ebcp_control_device](context.Background(), common.GetDaprClient(), model.Ebcp_control_deviceTableInfo.Name, model.Ebcp_control_device_FIELD_NAME_item_id+"="+deviceId)
		if err != nil {
			return fmt.Errorf("查询中控设备失败: %v", err)
		}

		if cmdType == "start" {
			// 启动展项
			exhibitionItem.Status = ItemStatusStart
			exhibitionItem.UpdatedTime = common.LocalTime(time.Now())
			common.DbUpsert[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(), *exhibitionItem, model.Ebcp_exhibition_itemTableInfo.Name, "id")
			for _, controlDevice := range controlDevices {
				if controlDevice.Status == ItemStatusStart {
					continue
				}
				controlDevice.Status = ItemStatusStart
				controlDevice.UpdatedTime = common.LocalTime(time.Now())
				common.DbUpsert[model.Ebcp_control_device](context.Background(), common.GetDaprClient(), controlDevice, model.Ebcp_control_deviceTableInfo.Name, "id")
			}
		} else if cmdType == "stop" {
			// 停止展项
			exhibitionItem.Status = ItemStatusStop
			exhibitionItem.UpdatedTime = common.LocalTime(time.Now())
			common.DbUpsert[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(), *exhibitionItem, model.Ebcp_exhibition_itemTableInfo.Name, "id")
			for _, controlDevice := range controlDevices {
				if controlDevice.Status == ItemStatusStop {
					continue
				}
				controlDevice.Status = ItemStatusStop
				controlDevice.UpdatedTime = common.LocalTime(time.Now())
				common.DbUpsert[model.Ebcp_control_device](context.Background(), common.GetDaprClient(), controlDevice, model.Ebcp_control_deviceTableInfo.Name, "id")
			}
		} else if cmdType == "pause" {
			// 暂停展项
			exhibitionItem.Status = ItemStatusPause
			exhibitionItem.UpdatedTime = common.LocalTime(time.Now())
			common.DbUpsert[model.Ebcp_exhibition_item](context.Background(), common.GetDaprClient(), *exhibitionItem, model.Ebcp_exhibition_itemTableInfo.Name, "id")
			for _, controlDevice := range controlDevices {
				if controlDevice.Status == ItemStatusPause {
					continue
				}
				controlDevice.Status = ItemStatusPause
				controlDevice.UpdatedTime = common.LocalTime(time.Now())
				common.DbUpsert[model.Ebcp_control_device](context.Background(), common.GetDaprClient(), controlDevice, model.Ebcp_control_deviceTableInfo.Name, "id")
			}
		}
	}

	common.Logger.Infof("向中控设备 %s:%d 发送控制命令成功: %s", deviceIP, devicePort, command)
	return nil
}

func SendUDPCommand(ip string, port int32, command string) error {
	// 构建目标地址
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return fmt.Errorf("解析UDP地址失败: %v", err)
	}

	// 将十六进制字符串转换为字节数组
	commandBytes, err := hexStringToBytes(command)
	if err != nil {
		return fmt.Errorf("十六进制命令转换失败: %v", err)
	}

	// 建立UDP连接
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return fmt.Errorf("建立UDP连接失败: %v", err)
	}
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			common.Logger.Errorf("关闭UDP连接失败: %v", closeErr)
		}
	}()

	// 设置发送超时
	if err := conn.SetWriteDeadline(time.Now().Add(5 * time.Second)); err != nil {
		return fmt.Errorf("设置发送超时失败: %v", err)
	}

	// 发送命令
	n, err := conn.Write(commandBytes)
	if err != nil {
		return fmt.Errorf("发送UDP数据失败: %v", err)
	}

	// 验证发送的数据长度
	if n != len(commandBytes) {
		return fmt.Errorf("发送数据不完整: 期望 %d 字节，实际发送 %d 字节", len(commandBytes), n)
	}

	common.Logger.Debugf("UDP命令发送成功: %s (%d bytes) -> %s:%d", command, n, ip, port)
	return nil
}

func hexStringToBytes(hexStr string) ([]byte, error) {
	// 去除首尾空格并按空格分割
	hexStr = strings.TrimSpace(hexStr)
	if hexStr == "" {
		return nil, fmt.Errorf("十六进制字符串不能为空")
	}

	hexParts := strings.Split(hexStr, " ")
	var result []byte

	for i, part := range hexParts {
		// 去除每个部分的空格
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		// 验证十六进制格式
		if len(part) != 2 {
			return nil, fmt.Errorf("第 %d 个十六进制数据格式错误: '%s'，应为2位十六进制数", i+1, part)
		}

		// 转换为字节
		b, err := strconv.ParseUint(part, 16, 8)
		if err != nil {
			return nil, fmt.Errorf("第 %d 个十六进制数据转换失败: '%s', %v", i+1, part, err)
		}

		result = append(result, byte(b))
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("未解析到有效的十六进制数据")
	}

	return result, nil
}
