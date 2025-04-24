package client

import (
	"encoding/hex"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestSendHexData(t *testing.T) {
	// 目标地址和端口
	target := "192.168.3.234:17613"

	// 建立TCP连接
	conn, err := net.Dial("tcp", target)
	if err != nil {
		t.Fatalf("无法连接到 %s: %v", target, err)
	}
	defer conn.Close()

	// 设置读取超时
	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		t.Fatalf("设置读取超时失败: %v", err)
	}

	// 定义要发送的十六进制数据
	hexData := "CC55CC5501000100010008001401040000000000"

	// 将十六进制字符串转换为字节数组
	data, err := hex.DecodeString(hexData)
	if err != nil {
		t.Fatalf("解析十六进制数据失败: %v", err)
	}

	fmt.Printf("正在发送数据: %X\n", data)

	// 发送数据
	_, err = conn.Write(data)
	if err != nil {
		t.Fatalf("发送数据失败: %v", err)
	}

	// 接收响应
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		t.Fatalf("接收响应失败: %v", err)
	}

	// 打印接收到的响应
	fmt.Printf("接收到 %d 字节响应:\n", n)
	fmt.Printf("十六进制: %X\n", buffer[:n])
}

// 手动执行测试的工具函数
func SendHexCommand(address string, hexData string) error {
	// 建立TCP连接
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("无法连接到 %s: %v", address, err)
	}
	defer conn.Close()

	// 设置读取超时
	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return fmt.Errorf("设置读取超时失败: %v", err)
	}

	// 将十六进制字符串转换为字节数组
	data, err := hex.DecodeString(hexData)
	if err != nil {
		return fmt.Errorf("解析十六进制数据失败: %v", err)
	}

	fmt.Printf("正在发送数据: %X\n", data)

	// 发送数据
	_, err = conn.Write(data)
	if err != nil {
		return fmt.Errorf("发送数据失败: %v", err)
	}

	// 接收响应
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return fmt.Errorf("接收响应失败: %v", err)
	}

	// 打印接收到的响应
	fmt.Printf("接收到 %d 字节响应:\n", n)
	fmt.Printf("十六进制: %X\n", buffer[:n])

	return nil
}
