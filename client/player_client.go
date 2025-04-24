package client

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const (
	// Protocol constants
	ProtocolHeader  uint32 = 0x55CC55CC
	ProtocolVersion uint16 = 0x0001

	// Command tags
	TagGetProgramList uint16 = 129
	TagFadeProgram    uint16 = 131
	TagCutProgram     uint16 = 132
	TagPauseProgram   uint16 = 133
	TagPlayProgram    uint16 = 271
	TagStopProgram    uint16 = 272

	// New command tags
	TagMuteMediaSound       uint16 = 270 // 关闭指定节目中指定媒体的声音
	TagPlayLayerMedia       uint16 = 273 // 播放当前节目中指定图层媒体
	TagPauseLayerMedia      uint16 = 274 // 暂停当前节目中指定图层媒体
	TagGetAllProgramMedia   uint16 = 276 // 获取节目中所有媒体
	TagControlLayerProgress uint16 = 283 // 控制图层的播放进度
	TagQueryLayerProgress   uint16 = 293 // 查询图层的播放进度

	// Sound control tags
	TagOpenGlobalSound  uint16 = 262 // 0x0106
	TagCloseGlobalSound uint16 = 263 // 0x0107
	TagSetGlobalVolume  uint16 = 264 // 0x0108
	TagIncreaseVolume   uint16 = 328 // 0x0148
	TagDecreaseVolume   uint16 = 329 // 0x0149

	// Connection constants
	connectTimeout = 5 * time.Second
	readTimeout    = 10 * time.Second

	bufferSize = 655360
)

// Program represents a program in the player
type Program struct {
	ID      uint32
	Name    string
	IsEmpty bool
}

// Media represents a media item in a program
type Media struct {
	ID   uint32
	Name string
}

// MediaListResponse represents response of GetAllProgramMedia
type MediaListResponse struct {
	ProgramID uint32
	Media     []Media
}

// LayerProgressResponse represents the response of QueryLayerProgress
type LayerProgressResponse struct {
	Success    bool
	LayerIndex uint16
	RemainTime uint32
	TotalTime  uint32
}

// ProgramListResponse represents response of GetProgramList
type ProgramListResponse struct {
	TotalCount int
	Programs   []Program
}

// PlayerClient represents a client to control media player
type PlayerClient struct {
	addr     string
	sequence uint16
	mu       sync.Mutex
}

// NewTCPClient creates a new TCP client
func NewTCPClient(addr string) (*PlayerClient, error) {
	client := &PlayerClient{
		addr: addr,
	}
	return client, nil
}

// buildPacket builds a protocol packet
func (c *PlayerClient) buildPacket(tag uint16, data []byte) []byte {
	contentLen := len(data)
	if data == nil {
		contentLen = 0
	}

	// Calculate total packet size
	packetSize := 12 + 4 + contentLen // Header(12) + TLV header(4) + data
	packet := make([]byte, packetSize)

	// 获取当前序列号并增加它，需要锁保护
	c.mu.Lock()
	currentSequence := c.sequence
	c.sequence++
	c.mu.Unlock()

	// Write header
	binary.LittleEndian.PutUint32(packet[0:], ProtocolHeader)        // 0x55CC55CC
	binary.LittleEndian.PutUint16(packet[4:], 1)                     // Packet type should be 1
	binary.LittleEndian.PutUint16(packet[6:], ProtocolVersion)       // Protocol version
	binary.LittleEndian.PutUint16(packet[8:], currentSequence)       // Sequence number
	binary.LittleEndian.PutUint16(packet[10:], uint16(4+contentLen)) // Content length = TLV header(4) + data

	// Write TLV
	binary.LittleEndian.PutUint16(packet[12:], tag)                // Tag
	binary.LittleEndian.PutUint16(packet[14:], uint16(contentLen)) // Length

	// Write data if any
	if contentLen > 0 {
		copy(packet[16:], data) // Value
	}

	return packet
}

// bytesToHexString converts bytes to hex string with spaces
func bytesToHexString(data []byte) string {
	hexStr := hex.EncodeToString(data)
	var parts []string
	for i := 0; i < len(hexStr); i += 2 {
		parts = append(parts, hexStr[i:i+2])
	}
	return strings.ToUpper(strings.Join(parts, " "))
}

// isValidPacketHeader checks if the given position in data is a valid packet header
func isValidPacketHeader(data []byte, offset int) bool {
	if offset+4 > len(data) {
		return false
	}
	// 0x55CC55CC in little-endian is CC 55 CC 55
	return data[offset] == 0xCC && data[offset+1] == 0x55 &&
		data[offset+2] == 0xCC && data[offset+3] == 0x55
}

// sendCommand sends command and returns response
func (c *PlayerClient) sendCommand(tag uint16, data []byte) ([]byte, error) {
	// 创建连接
	conn, err := net.DialTimeout("tcp", c.addr, connectTimeout)
	if err != nil {
		return nil, fmt.Errorf("connect failed: %v", err)
	}
	defer conn.Close()

	// 为大型响应设置更长的读取超时
	var maxWaitTime time.Duration
	if tag == TagGetProgramList || tag == TagGetAllProgramMedia {
		// 程序列表和媒体列表可能很大，设置更长的超时时间
		maxWaitTime = 30 * time.Second
	} else {
		maxWaitTime = readTimeout
	}

	err = conn.SetReadDeadline(time.Now().Add(maxWaitTime))
	if err != nil {
		return nil, fmt.Errorf("set read deadline failed: %v", err)
	}

	// 构建数据包
	packet := c.buildPacket(tag, data)
	fmt.Printf("Send packet: %s\n", bytesToHexString(packet))

	// 发送数据
	_, err = conn.Write(packet)
	if err != nil {
		return nil, fmt.Errorf("write failed: %v", err)
	}

	// 对于大型响应命令，预分配更大的缓冲区
	var buffer []byte
	if tag == TagGetProgramList || tag == TagGetAllProgramMedia {
		buffer = make([]byte, bufferSize*2) // 为大型响应分配更大的缓冲区
	} else {
		buffer = make([]byte, bufferSize)
	}

	totalRead := 0
	lastDataTime := time.Now()
	receiveStartTime := time.Now()

	// 设置绝对超时，防止无限等待
	absoluteTimeout := time.Now().Add(maxWaitTime)

	// 读取响应数据，可能需要多次读取
	for time.Now().Before(absoluteTimeout) {
		// 根据距离上次接收数据的时间来调整读取超时
		timeLeft := absoluteTimeout.Sub(time.Now())
		if timeLeft <= 0 {
			break // 已达到绝对超时时间
		}

		// 计算本次读取的超时时间
		var readTimeout time.Duration
		if totalRead == 0 {
			// 第一次读取给予更长的等待时间
			readTimeout = 5 * time.Second
		} else if time.Since(lastDataTime) > 2*time.Second {
			// 如果距离上次读取数据超过2秒，使用短超时
			readTimeout = 1 * time.Second
		} else {
			// 正常读取中使用中等超时
			readTimeout = 2 * time.Second
		}

		// 超时不应超过剩余的绝对超时时间
		if readTimeout > timeLeft {
			readTimeout = timeLeft
		}

		// 设置本次读取的截止时间
		err = conn.SetReadDeadline(time.Now().Add(readTimeout))
		if err != nil {
			return nil, fmt.Errorf("set read deadline failed: %v", err)
		}

		n, err := conn.Read(buffer[totalRead:])
		currentTime := time.Now()

		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				// 超时错误处理
				if totalRead > 0 {
					// 如果已经有数据且满足以下条件之一，认为传输完成:
					// 1. 距离第一个字节接收时间已超过最大等待时间的80%
					// 2. 距离上次接收数据已超过3秒
					if time.Since(receiveStartTime) > maxWaitTime*4/5 ||
						time.Since(lastDataTime) > 3*time.Second {
						fmt.Printf("Read timeout after receiving %d bytes, considering complete\n", totalRead)
						break
					}
				}
				// 否则继续尝试读取
				continue
			}
			// 非超时错误
			return nil, fmt.Errorf("read failed: %v", err)
		}

		if n == 0 {
			// 连接关闭
			break
		}

		// 更新计时器
		if totalRead == 0 {
			receiveStartTime = currentTime // 第一个字节的接收时间
		}
		lastDataTime = currentTime

		totalRead += n
		fmt.Printf("Received %d bytes, total %d bytes\n", n, totalRead)

		// 检查缓冲区容量，需要时扩展
		if totalRead > len(buffer)*3/4 {
			// 当缓冲区使用超过75%时扩展
			newBuffer := make([]byte, len(buffer)*2)
			copy(newBuffer, buffer)
			buffer = newBuffer
			fmt.Printf("Expanded buffer to %d bytes\n", len(buffer))
		}

		// 特殊命令可能要接收大量数据，需要特殊处理
		if tag == TagGetProgramList || tag == TagGetAllProgramMedia {
			// 这些命令可能返回多个包，需要等待更长时间
			// 检查是否有足够的完整包
			if isEnoughDataForCommand(buffer[:totalRead], tag) {
				fmt.Printf("Received enough data for command %d\n", tag)
				break
			}
			continue
		}

		// 检查是否收到完整响应
		if totalRead >= 16 {
			if isValidPacketHeader(buffer, 0) {
				contentLength := binary.LittleEndian.Uint16(buffer[10:12])
				if totalRead >= int(contentLength)+12 { // 12是头部长度
					// 尝试再读取一次，看是否有更多数据
					extraTimeout := 500 * time.Millisecond
					if extraTimeout > timeLeft {
						extraTimeout = timeLeft
					}

					conn.SetReadDeadline(time.Now().Add(extraTimeout))
					time.Sleep(100 * time.Millisecond)

					n, err = conn.Read(buffer[totalRead:])
					if err != nil {
						if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
							// 超时且没有更多数据，认为传输完成
							break
						}
						if n == 0 {
							// 连接已关闭，完成读取
							break
						}
						// 只有在非超时且非连接关闭的情况下返回错误
						return nil, fmt.Errorf("final read failed: %v", err)
					}

					if n > 0 {
						// 还有更多数据，继续读取
						totalRead += n
						lastDataTime = time.Now()
						fmt.Printf("Additional data received: %d bytes, total %d bytes\n", n, totalRead)
						continue
					}

					// 没有更多数据，完成读取
					break
				}
			}
		}
	}

	// 如果完全没有读到数据，返回错误
	if totalRead == 0 {
		return nil, fmt.Errorf("no data received after %v", time.Since(receiveStartTime))
	}

	// 验证响应
	if totalRead < 16 {
		return nil, fmt.Errorf("response too short: received only %d bytes", totalRead)
	}

	// 验证帧头
	if !isValidPacketHeader(buffer, 0) {
		return nil, errors.New("invalid packet header")
	}

	// 提取响应信息
	packetHeader := binary.LittleEndian.Uint32(buffer[0:4])
	packetType := binary.LittleEndian.Uint16(buffer[4:6])
	protocolVersion := binary.LittleEndian.Uint16(buffer[6:8])
	sequence := binary.LittleEndian.Uint16(buffer[8:10])
	contentLength := binary.LittleEndian.Uint16(buffer[10:12])
	tlvTag := binary.LittleEndian.Uint16(buffer[12:14])
	tlvLength := binary.LittleEndian.Uint16(buffer[14:16])

	fmt.Printf("Response details:\n")
	fmt.Printf("  Header: 0x%08X\n", packetHeader)
	fmt.Printf("  Packet Type: %d\n", packetType)
	fmt.Printf("  Protocol Version: 0x%04X\n", protocolVersion)
	fmt.Printf("  Sequence: %d\n", sequence)
	fmt.Printf("  Content Length: %d\n", contentLength)
	fmt.Printf("  TLV Tag: %d (0x%04X) (Expected: %d)\n", tlvTag, tlvTag, tag)
	fmt.Printf("  TLV Length: %d\n", tlvLength)
	fmt.Printf("  Total bytes received: %d\n", totalRead)
	fmt.Printf("  Total receive time: %v\n", time.Since(receiveStartTime))

	// 特殊情况处理：
	// 1. 某些服务器实现可能返回TLV标签为0
	// 2. QueryLayerProgress命令(293/0x0125)可能返回标签28(0x001C)
	if tlvTag != tag && tlvTag != 0 {
		// 如果是QueryLayerProgress并且返回标签是28，这是正常的
		if tag == TagQueryLayerProgress && tlvTag == 28 {
			fmt.Printf("  Accepting tag 28(0x001C) as valid response for QueryLayerProgress\n")
		} else {
			// 其他情况下标签不匹配视为错误
			return nil, fmt.Errorf("unexpected TLV tag: got %d (0x%04X), want %d (0x%04X)",
				tlvTag, tlvTag, tag, tag)
		}
	}

	return buffer[:totalRead], nil
}

// isEnoughDataForCommand 检查是否接收到足够的数据来处理特定命令
func isEnoughDataForCommand(data []byte, tag uint16) bool {
	if len(data) < 16 {
		return false // 数据不足以包含头部
	}

	// 对于程序列表，检查是否包含至少一个完整程序
	if tag == TagGetProgramList {
		// 从第一个响应包中获取程序数量
		if !isValidPacketHeader(data, 0) {
			return false
		}

		if len(data) < 18 {
			return false // 数据不足以包含程序数量
		}

		programCount := int(binary.LittleEndian.Uint16(data[16:18]))
		fmt.Printf("Program count in response: %d\n", programCount)

		// 寻找至少3个有效的程序包或所有程序（如果程序总数小于3）
		minPrograms := 3
		if programCount < minPrograms {
			minPrograms = programCount
		}

		foundPrograms := 0
		offset := 0

		for offset < len(data)-4 {
			if isValidPacketHeader(data, offset) {
				foundPrograms++
				if foundPrograms >= minPrograms {
					return true
				}

				// 跳过这个包
				if offset+12 < len(data) {
					contentLength := binary.LittleEndian.Uint16(data[offset+10 : offset+12])
					offset += 12 + int(contentLength)
				} else {
					// 数据可能不完整
					break
				}
			} else {
				offset++
			}
		}

		return false
	}

	// 对于媒体列表，检查是否至少接收到一个完整包
	if tag == TagGetAllProgramMedia {
		// 媒体项有固定大小，检查是否至少有一个完整的媒体项
		if !isValidPacketHeader(data, 0) {
			return false
		}

		if len(data) < 16 {
			return false
		}

		// 检查内容长度
		contentLength := binary.LittleEndian.Uint16(data[10:12])
		return len(data) >= 12+int(contentLength)
	}

	// 默认情况
	return false
}

// sendCommandWithTimeout sends command with a specific timeout and returns response
func (c *PlayerClient) sendCommandWithTimeout(tag uint16, data []byte, timeout time.Duration) ([]byte, error) {
	// 创建连接
	conn, err := net.DialTimeout("tcp", c.addr, connectTimeout)
	if err != nil {
		return nil, fmt.Errorf("connect failed: %v", err)
	}
	defer conn.Close()

	// 设置特定的读取超时
	err = conn.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return nil, fmt.Errorf("set read deadline failed: %v", err)
	}

	// 构建数据包
	packet := c.buildPacket(tag, data)
	fmt.Printf("Send packet: %s\n", bytesToHexString(packet))

	// 发送数据
	_, err = conn.Write(packet)
	if err != nil {
		return nil, fmt.Errorf("write failed: %v", err)
	}

	// 对于GetProgramList命令，预分配更大的缓冲区
	buffer := make([]byte, bufferSize*2) // 为程序列表分配更大的缓冲区

	totalRead := 0
	lastDataTime := time.Now()
	receiveStartTime := time.Now()

	// 设置绝对超时，防止无限等待
	absoluteTimeout := time.Now().Add(timeout)

	// 读取响应数据，可能需要多次读取
	for time.Now().Before(absoluteTimeout) {
		// 根据距离上次接收数据的时间来调整读取超时
		timeLeft := absoluteTimeout.Sub(time.Now())
		if timeLeft <= 0 {
			break // 已达到绝对超时时间
		}

		// 计算本次读取的超时时间
		var readTimeout time.Duration
		if totalRead == 0 {
			// 第一次读取给予更长的等待时间
			readTimeout = 5 * time.Second
		} else if time.Since(lastDataTime) > 2*time.Second {
			// 如果距离上次读取数据超过2秒，使用短超时
			readTimeout = 1 * time.Second
		} else {
			// 正常读取中使用中等超时
			readTimeout = 2 * time.Second
		}

		// 超时不应超过剩余的绝对超时时间
		if readTimeout > timeLeft {
			readTimeout = timeLeft
		}

		// 设置本次读取的截止时间
		err = conn.SetReadDeadline(time.Now().Add(readTimeout))
		if err != nil {
			return nil, fmt.Errorf("set read deadline failed: %v", err)
		}

		n, err := conn.Read(buffer[totalRead:])
		currentTime := time.Now()

		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				// 超时错误处理
				if totalRead > 0 {
					// 如果已经有数据且满足以下条件之一，认为传输完成:
					// 1. 距离第一个字节接收时间已超过最大等待时间的80%
					// 2. 距离上次接收数据已超过3秒
					if time.Since(receiveStartTime) > timeout*4/5 ||
						time.Since(lastDataTime) > 3*time.Second {
						fmt.Printf("Read timeout after receiving %d bytes, considering complete\n", totalRead)
						break
					}
				}
				// 否则继续尝试读取
				continue
			}
			// 非超时错误
			return nil, fmt.Errorf("read failed: %v", err)
		}

		if n == 0 {
			// 连接关闭
			break
		}

		// 更新计时器
		if totalRead == 0 {
			receiveStartTime = currentTime // 第一个字节的接收时间
		}
		lastDataTime = currentTime

		totalRead += n
		fmt.Printf("Received %d bytes, total %d bytes\n", n, totalRead)

		// 检查缓冲区容量，需要时扩展
		if totalRead > len(buffer)*3/4 {
			// 当缓冲区使用超过75%时扩展
			newBuffer := make([]byte, len(buffer)*2)
			copy(newBuffer, buffer)
			buffer = newBuffer
			fmt.Printf("Expanded buffer to %d bytes\n", len(buffer))
		}

		// 对于大型响应，我们需要接收尽可能多的数据
		if tag == TagGetProgramList {
			// 持续接收直到超时或者收到足够多的数据
			continue
		}

		// 检查是否收到完整响应
		if totalRead >= 16 {
			if isValidPacketHeader(buffer, 0) {
				contentLength := binary.LittleEndian.Uint16(buffer[10:12])
				if totalRead >= int(contentLength)+12 { // 12是头部长度
					// 尝试再读取一次，看是否有更多数据
					extraTimeout := 500 * time.Millisecond
					if extraTimeout > timeLeft {
						extraTimeout = timeLeft
					}

					conn.SetReadDeadline(time.Now().Add(extraTimeout))
					time.Sleep(100 * time.Millisecond)

					n, err = conn.Read(buffer[totalRead:])
					if err != nil {
						if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
							// 超时且没有更多数据，认为传输完成
							break
						}
						if n == 0 {
							// 连接已关闭，完成读取
							break
						}
						// 只有在非超时且非连接关闭的情况下返回错误
						return nil, fmt.Errorf("final read failed: %v", err)
					}

					if n > 0 {
						// 还有更多数据，继续读取
						totalRead += n
						lastDataTime = time.Now()
						fmt.Printf("Additional data received: %d bytes, total %d bytes\n", n, totalRead)
						continue
					}

					// 没有更多数据，完成读取
					break
				}
			}
		}
	}

	// 如果完全没有读到数据，返回错误
	if totalRead == 0 {
		return nil, fmt.Errorf("no data received after %v", time.Since(receiveStartTime))
	}

	// 验证响应
	if totalRead < 16 {
		return nil, fmt.Errorf("response too short: received only %d bytes", totalRead)
	}

	// 验证帧头
	if !isValidPacketHeader(buffer, 0) {
		return nil, errors.New("invalid packet header")
	}

	// 提取响应信息
	packetHeader := binary.LittleEndian.Uint32(buffer[0:4])
	packetType := binary.LittleEndian.Uint16(buffer[4:6])
	protocolVersion := binary.LittleEndian.Uint16(buffer[6:8])
	sequence := binary.LittleEndian.Uint16(buffer[8:10])
	contentLength := binary.LittleEndian.Uint16(buffer[10:12])
	tlvTag := binary.LittleEndian.Uint16(buffer[12:14])
	tlvLength := binary.LittleEndian.Uint16(buffer[14:16])

	fmt.Printf("Response details:\n")
	fmt.Printf("  Header: 0x%08X\n", packetHeader)
	fmt.Printf("  Packet Type: %d\n", packetType)
	fmt.Printf("  Protocol Version: 0x%04X\n", protocolVersion)
	fmt.Printf("  Sequence: %d\n", sequence)
	fmt.Printf("  Content Length: %d\n", contentLength)
	fmt.Printf("  TLV Tag: %d (0x%04X) (Expected: %d)\n", tlvTag, tlvTag, tag)
	fmt.Printf("  TLV Length: %d\n", tlvLength)
	fmt.Printf("  Total bytes received: %d\n", totalRead)
	fmt.Printf("  Total receive time: %v\n", time.Since(receiveStartTime))

	// 特殊情况处理：
	// 1. 某些服务器实现可能返回TLV标签为0
	// 2. QueryLayerProgress命令(293/0x0125)可能返回标签28(0x001C)
	if tlvTag != tag && tlvTag != 0 {
		// 如果是QueryLayerProgress并且返回标签是28，这是正常的
		if tag == TagQueryLayerProgress && tlvTag == 28 {
			fmt.Printf("  Accepting tag 28(0x001C) as valid response for QueryLayerProgress\n")
		} else {
			// 其他情况下标签不匹配视为错误
			return nil, fmt.Errorf("unexpected TLV tag: got %d (0x%04X), want %d (0x%04X)",
				tlvTag, tlvTag, tag, tag)
		}
	}

	return buffer[:totalRead], nil
}

// GetProgramList gets the program list
func (c *PlayerClient) GetProgramList() (*ProgramListResponse, error) {
	// 为程序列表专门使用更长的超时设置
	longerTimeout := 30 * time.Second

	// 多次尝试获取完整的程序列表
	var resp []byte
	var err error
	maxRetries := 3

	for attempt := 0; attempt < maxRetries; attempt++ {
		// Send empty data for query
		resp, err = c.sendCommandWithTimeout(TagGetProgramList, nil, longerTimeout)
		if err != nil {
			if attempt < maxRetries-1 {
				fmt.Printf("GetProgramList attempt %d failed: %v, retrying...\n", attempt+1, err)
				time.Sleep(1 * time.Second)
				continue
			}
			return nil, fmt.Errorf("GetProgramList failed after %d attempts: %v", maxRetries, err)
		}

		// 如果响应数据长度明显太短，可能是不完整的
		if len(resp) < 100 {
			if attempt < maxRetries-1 {
				fmt.Printf("GetProgramList response too short (%d bytes), retrying...\n", len(resp))
				time.Sleep(1 * time.Second)
				continue
			}
		}

		// 收到看起来有效的响应
		break
	}

	// Each program comes in a separate packet
	// Each packet: CC 55 CC 55 + version(2) + sequence(2) + tag(2) + length(2) + TLV
	const (
		headerSize       = 12 // CC 55 CC 55 + version + sequence + tag + length
		tlvHeaderSize    = 4  // TLV tag(2) + length(2)
		programCountSize = 4  // Program count is 2 bytes
		programIndexSize = 4
		programIdSize    = 4
		programNameSize  = 128
		isEmptySize      = 1
	)

	// First packet contains the program count
	if len(resp) < headerSize+tlvHeaderSize+programCountSize {
		return nil, fmt.Errorf("response too short for header and program count: len(resp)=%d", len(resp))
	}

	// Get program count from first packet
	programCount := int(binary.LittleEndian.Uint16(resp[16:18]))
	fmt.Printf("Program count: %d (response length: %d bytes)\n", programCount, len(resp))

	// Create result structure
	result := &ProgramListResponse{
		TotalCount: programCount,
		Programs:   make([]Program, 0, programCount),
	}

	// 计算期望的最小响应大小
	expectedMinSize := headerSize + tlvHeaderSize + programCountSize + // 第一个包头和程序计数
		// 每个程序数据包的大小（去掉了programCountSize，因为它只在第一个包中出现）
		(headerSize+tlvHeaderSize+programIndexSize+programIdSize+programNameSize+isEmptySize)*programCount

	if len(resp) < expectedMinSize {
		fmt.Printf("Warning: Response size %d is smaller than expected minimum %d for %d programs\n",
			len(resp), expectedMinSize, programCount)
	}

	// 确定在可用数据中可以解析的程序数量
	maxProgramsInResponse := 0
	offset := 0
	validPrograms := 0

	// 第一步：统计有多少个有效的程序数据包
	for offset < len(resp)-headerSize {
		// 寻找下一个有效包头
		found := false
		for i := offset; i < len(resp)-4; i++ {
			if isValidPacketHeader(resp, i) {
				offset = i
				found = true
				break
			}
		}

		if !found {
			break // 找不到下一个有效包头
		}

		// 检查是否还有足够的数据读取此包
		if offset+headerSize > len(resp) {
			break
		}

		// 读取包长度
		contentLength := binary.LittleEndian.Uint16(resp[offset+10 : offset+12])

		// 检查是否有完整的包内容
		if offset+headerSize+int(contentLength) > len(resp) {
			break // 此包内容不完整
		}

		// 如果这是一个有效的程序数据包
		if offset+headerSize+tlvHeaderSize+programIndexSize+programIdSize+programNameSize+isEmptySize <= len(resp) {
			validPrograms++
		}

		// 移动到下一个包
		offset += headerSize + int(contentLength)
	}

	maxProgramsInResponse = validPrograms
	fmt.Printf("Found %d valid program entries in response\n", maxProgramsInResponse)

	// 如果找到的有效程序数为0，但程序数应该大于0，返回错误
	if maxProgramsInResponse == 0 && programCount > 0 {
		return nil, fmt.Errorf("no valid program data found in response")
	}

	// 第二步：解析程序数据
	offset = 0
	parsedPrograms := 0

	for i := 0; i < programCount && parsedPrograms < maxProgramsInResponse; i++ {
		// 寻找下一个有效包头
		found := false
		for j := offset; j < len(resp)-4; j++ {
			if isValidPacketHeader(resp, j) {
				offset = j
				found = true
				break
			}
		}

		if !found {
			break // 找不到下一个有效包头
		}

		// 验证是否有足够的数据
		packetStart := offset
		if packetStart+headerSize+tlvHeaderSize+programIndexSize+programIdSize+programNameSize+isEmptySize > len(resp) {
			fmt.Printf("Incomplete data for program %d at offset %d\n", i, packetStart)
			break
		}

		// 跳过包头和TLV头部
		dataStart := packetStart + headerSize + tlvHeaderSize

		// 读取程序数据
		program := Program{
			ID:      binary.LittleEndian.Uint32(resp[dataStart+programIndexSize : dataStart+programIndexSize+programIdSize]),
			IsEmpty: resp[dataStart+programIndexSize+programIdSize+programNameSize] == 0,
		}

		// 读取以null结尾的程序名
		nameBytes := resp[dataStart+programIndexSize+programIdSize : dataStart+programIndexSize+programIdSize+programNameSize]
		nullPos := bytes.IndexByte(nameBytes, 0)
		if nullPos != -1 {
			program.Name = string(nameBytes[:nullPos])
		} else {
			program.Name = string(bytes.TrimRight(nameBytes, "\x00"))
		}

		result.Programs = append(result.Programs, program)
		parsedPrograms++

		// 移动到下一个包
		offset = packetStart + headerSize + tlvHeaderSize + programIndexSize + programIdSize + programNameSize + isEmptySize
	}

	// 更新实际解析的程序数量
	result.TotalCount = parsedPrograms
	fmt.Printf("Successfully parsed %d/%d programs\n", parsedPrograms, programCount)

	if len(result.Programs) > 0 {
		return result, nil
	}

	return nil, fmt.Errorf("failed to parse any programs from response")
}

// FadeProgram fades to specified program
func (c *PlayerClient) FadeProgram(programID uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, programID)
	_, err := c.sendCommand(TagFadeProgram, data)
	return err
}

// CutProgram cuts to specified program
func (c *PlayerClient) CutProgram(programID uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, programID)
	_, err := c.sendCommand(TagCutProgram, data)
	return err
}

// PauseProgram pauses specified program
func (c *PlayerClient) PauseProgram(programID uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, programID)
	_, err := c.sendCommand(TagPauseProgram, data)
	return err
}

// PlayProgram plays specified program
func (c *PlayerClient) PlayProgram(programID uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, programID)
	_, err := c.sendCommand(TagPlayProgram, data)
	return err
}

// StopProgram stops specified program
func (c *PlayerClient) StopProgram(programID uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, programID)
	_, err := c.sendCommand(TagStopProgram, data)
	return err
}

// OpenGlobalSound opens global sound
func (c *PlayerClient) OpenGlobalSound() error {
	_, err := c.sendCommand(TagOpenGlobalSound, nil)
	return err
}

// CloseGlobalSound closes global sound
func (c *PlayerClient) CloseGlobalSound() error {
	_, err := c.sendCommand(TagCloseGlobalSound, nil)
	return err
}

// SetGlobalVolume sets global volume (0-100)
func (c *PlayerClient) SetGlobalVolume(volume uint8) error {
	if volume > 100 {
		return fmt.Errorf("volume must be between 0 and 100")
	}
	data := []byte{volume}
	_, err := c.sendCommand(TagSetGlobalVolume, data)
	return err
}

// IncreaseGlobalVolume increases global volume by step
func (c *PlayerClient) IncreaseGlobalVolume(step uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, step)
	_, err := c.sendCommand(TagIncreaseVolume, data)
	return err
}

// DecreaseGlobalVolume decreases global volume by step
func (c *PlayerClient) DecreaseGlobalVolume(step uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, step)
	_, err := c.sendCommand(TagDecreaseVolume, data)
	return err
}

// GetAllProgramMedia gets all media in a program
func (c *PlayerClient) GetAllProgramMedia(programID uint32) (*MediaListResponse, error) {
	// Prepare request data
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, programID)

	// Send command
	resp, err := c.sendCommand(TagGetAllProgramMedia, data)
	if err != nil {
		return nil, fmt.Errorf("GetAllProgramMedia failed: %v", err)
	}

	// Parse response
	if len(resp) < 16 {
		return nil, errors.New("response too short")
	}

	// Create result structure
	result := &MediaListResponse{
		ProgramID: programID,
		Media:     make([]Media, 0),
	}

	// Each media item has: MediaId(4 bytes) + MediaName(32 bytes)
	const (
		headerSize    = 12 // CC 55 CC 55 + version + sequence + tag + length
		tlvHeaderSize = 4  // TLV tag(2) + length(2)
		mediaIdSize   = 4
		mediaNameSize = 32
		mediaItemSize = mediaIdSize + mediaNameSize
	)

	// 获取实际的TLV长度
	tlvLength := binary.LittleEndian.Uint16(resp[14:16])
	fmt.Printf("TLV length from response: %d\n", tlvLength)

	// 根据TLV长度计算媒体项数量，而不是使用固定的计算方式
	// 这样可以适应不同的服务器实现
	dataSize := int(tlvLength)
	if dataSize <= 0 {
		// 回退到原来的计算方式
		dataSize = len(resp) - headerSize - tlvHeaderSize
		fmt.Printf("Using fallback dataSize calculation: %d\n", dataSize)
	}

	if dataSize < 0 {
		return nil, errors.New("invalid response size")
	}

	// 检查是否有至少一个媒体项
	if dataSize < mediaItemSize {
		fmt.Printf("Warning: Data size (%d) less than one media item size (%d)\n", dataSize, mediaItemSize)
		// 如果消息很短但不为零，尝试从中解析任何可能的媒体项
		if dataSize > 0 {
			fmt.Printf("Attempting to parse partial data: %x\n", resp[headerSize+tlvHeaderSize:])
		} else {
			// 没有媒体项数据，返回空列表
			return result, nil
		}
	}

	mediaCount := dataSize / mediaItemSize
	dataStart := headerSize + tlvHeaderSize

	// 记录实际解析出的媒体项数量
	parsedCount := 0

	fmt.Printf("Media resp dataSize: %d, expected mediaCount: %d, data: %x\n", dataSize, mediaCount, resp)

	// 如果数据量不是mediaItemSize的整数倍，可能有额外的头部或尾部
	if dataSize%mediaItemSize != 0 {
		fmt.Printf("Warning: data size %d is not a multiple of media item size %d\n", dataSize, mediaItemSize)
		// 尝试调整起始位置，查找有效数据的起始
		for offset := dataStart; offset < len(resp)-mediaIdSize; offset++ {
			// 寻找可能的媒体ID标记
			if resp[offset] == 0 && resp[offset+1] == 0 && resp[offset+2] == 0 && resp[offset+3] == 0 {
				potentialMediaStart := offset
				fmt.Printf("Potential media data starts at offset %d: %x\n", potentialMediaStart, resp[potentialMediaStart:potentialMediaStart+8])
				// 如果找到可能的媒体起始位置，调整dataStart
				if potentialMediaStart != dataStart {
					fmt.Printf("Adjusting dataStart from %d to %d\n", dataStart, potentialMediaStart)
					dataStart = potentialMediaStart
					break
				}
			}
		}
	}

	// Extract each media item - 使用剩余可用空间来计算可提取的媒体项数
	availableBytes := len(resp) - dataStart
	availableMediaItems := availableBytes / mediaItemSize

	// 如果可用的媒体项数少于预期的媒体项数，调整mediaCount
	if availableMediaItems < mediaCount {
		fmt.Printf("Warning: Can only extract %d media items from available data, instead of expected %d\n",
			availableMediaItems, mediaCount)
		mediaCount = availableMediaItems
	}

	for i := 0; i < mediaCount; i++ {
		offset := dataStart + (i * mediaItemSize)
		if offset+mediaItemSize > len(resp) {
			fmt.Printf("Warning: offset %d + mediaItemSize %d exceeds response length %d\n", offset, mediaItemSize, len(resp))
			break
		}

		// 检查当前媒体项位置的数据是否有效
		if offset+mediaIdSize > len(resp) {
			fmt.Printf("Warning: Not enough data for media ID at offset %d\n", offset)
			break
		}

		mediaId := binary.LittleEndian.Uint32(resp[offset : offset+mediaIdSize])
		fmt.Printf("Media %d ID: %d, raw data: %x\n", i, mediaId, resp[offset:min(offset+mediaItemSize, len(resp))])

		media := Media{
			ID: mediaId,
		}

		// 检查是否有足够的空间来读取媒体名称
		if offset+mediaIdSize+mediaNameSize > len(resp) {
			fmt.Printf("Warning: Not enough data for media name at offset %d\n", offset+mediaIdSize)
			// 即使名称不完整，仍添加这个媒体项
			result.Media = append(result.Media, media)
			parsedCount++
			break
		}

		// Read null-terminated media name
		// 协议文档说明：字符串数组最后一位为'\0'
		nameBytes := resp[offset+mediaIdSize : offset+mediaIdSize+mediaNameSize]
		nullPos := bytes.IndexByte(nameBytes, 0)

		var rawNameBytes []byte
		if nullPos != -1 {
			// 找到空字符，取空字符之前的有效数据
			rawNameBytes = nameBytes[:nullPos]
		} else {
			// 没有找到空字符，去掉尾部的空字节
			rawNameBytes = bytes.TrimRight(nameBytes, "\x00")
		}
		fmt.Printf("Media %d rawNameBytes: %x\n", i, rawNameBytes)

		// 先检查是否已经是有效的UTF-8编码
		if utf8Valid(rawNameBytes) {
			media.Name = string(rawNameBytes)
			fmt.Printf("Media %d is utf8: %s\n", i, media.Name)
		} else {
			// 尝试不同的中文编码转换
			converted := false

			// 尝试GBK (中国大陆最常用的中文编码)
			utf8NameBytes, err := io.ReadAll(transform.NewReader(bytes.NewReader(rawNameBytes), simplifiedchinese.GBK.NewDecoder()))
			if err == nil && utf8Valid(utf8NameBytes) {
				media.Name = string(utf8NameBytes)
				fmt.Printf("Media %d is gbk: %s\n", i, media.Name)
				converted = true
			}

			// 如果GBK失败，尝试GB18030 (GBK的超集，包含更多字符)
			if !converted {
				utf8NameBytes, err = io.ReadAll(transform.NewReader(bytes.NewReader(rawNameBytes), simplifiedchinese.GB18030.NewDecoder()))
				if err == nil && utf8Valid(utf8NameBytes) {
					media.Name = string(utf8NameBytes)
					fmt.Printf("Media %d is gb18030: %s\n", i, media.Name)
					converted = true
				}
			}

			// 如果GB18030失败，尝试HZGB2312 (另一种常见中文编码)
			if !converted {
				utf8NameBytes, err = io.ReadAll(transform.NewReader(bytes.NewReader(rawNameBytes), simplifiedchinese.HZGB2312.NewDecoder()))
				if err == nil && utf8Valid(utf8NameBytes) {
					media.Name = string(utf8NameBytes)
					fmt.Printf("Media is hzgb2312: %d: %s\n", i, media.Name)
					converted = true
				}
			}

			// 如果所有尝试都失败，使用原始字节
			if !converted {
				media.Name = string(rawNameBytes)
				fmt.Printf("Media is raw: %d: %s\n", i, media.Name)
			}
		}

		result.Media = append(result.Media, media)
		parsedCount++
	}

	// 打印实际解析的媒体项数量
	fmt.Printf("Successfully parsed %d/%d media items\n", parsedCount, mediaCount)

	return result, nil
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// utf8Valid checks if a byte slice contains valid UTF-8 data
func utf8Valid(data []byte) bool {
	return strings.HasPrefix(strings.ToValidUTF8(string(data), ""), string(data))
}

// PlayLayerMedia plays media in the specified layer of current program
func (c *PlayerClient) PlayLayerMedia(layerIndex uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, layerIndex)
	_, err := c.sendCommand(TagPlayLayerMedia, data)
	return err
}

// PauseLayerMedia pauses media in the specified layer of current program
func (c *PlayerClient) PauseLayerMedia(layerIndex uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, layerIndex)
	_, err := c.sendCommand(TagPauseLayerMedia, data)
	return err
}

// MuteMediaSound mutes sound for media at specified position
func (c *PlayerClient) MuteMediaSound(col, row uint32) error {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint32(data[0:4], col)
	binary.LittleEndian.PutUint32(data[4:8], row)
	_, err := c.sendCommand(TagMuteMediaSound, data)
	return err
}

// ControlLayerProgress controls playback progress of specified layer
func (c *PlayerClient) ControlLayerProgress(triggerID string, remainTime, totalTime uint32, layerIndex uint16) error {
	// TriggerID (36 bytes) + RemainTime (4 bytes) + TotalTime (4 bytes) + layerIndex (2 bytes)
	data := make([]byte, 36+4+4+2)

	// Copy triggerID (padding with zeros if needed)
	triggerIDBytes := []byte(triggerID)
	if len(triggerIDBytes) > 36 {
		triggerIDBytes = triggerIDBytes[:36]
	}
	copy(data[:36], triggerIDBytes)

	// Copy other fields
	binary.LittleEndian.PutUint32(data[36:40], remainTime)
	binary.LittleEndian.PutUint32(data[40:44], totalTime)
	binary.LittleEndian.PutUint16(data[44:46], layerIndex)

	_, err := c.sendCommand(TagControlLayerProgress, data)
	return err
}

// QueryLayerProgress queries playback progress of specified layer
func (c *PlayerClient) QueryLayerProgress(layerIndex uint16) (*LayerProgressResponse, error) {
	// Prepare request data
	data := make([]byte, 2)
	binary.LittleEndian.PutUint16(data, layerIndex)

	// Send command
	resp, err := c.sendCommand(TagQueryLayerProgress, data)
	if err != nil {
		return nil, fmt.Errorf("QueryLayerProgress failed: %v", err)
	}

	// Parse response
	if len(resp) < 16+11 { // header(12) + tlv(4) + success(1) + layerIndex(2) + remainTime(4) + totalTime(4)
		return nil, errors.New("response too short")
	}

	// 解析响应
	result := &LayerProgressResponse{
		Success:    resp[16] == 1,
		LayerIndex: binary.LittleEndian.Uint16(resp[17:19]),
		RemainTime: binary.LittleEndian.Uint32(resp[19:23]),
		TotalTime:  binary.LittleEndian.Uint32(resp[23:27]),
	}

	fmt.Printf("Progress data: success=%v, layerIndex=%d, remainTime=%d, totalTime=%d\n",
		result.Success, result.LayerIndex, result.RemainTime, result.TotalTime)

	return result, nil
}

// Close is a no-op since we don't maintain persistent connections
func (c *PlayerClient) Close() error {
	return nil
}

// GetAddress returns the address of the client connection
func (c *PlayerClient) GetAddress() string {
	return c.addr
}
