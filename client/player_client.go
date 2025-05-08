package client

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
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
	TagGetCurrentProgram    uint16 = 294 // 查询舞台当前播放节目ID

	// Sound control tags
	TagOpenGlobalSound  uint16 = 262 // 0x0106
	TagCloseGlobalSound uint16 = 263 // 0x0107
	TagSetGlobalVolume  uint16 = 264 // 0x0108
	TagIncreaseVolume   uint16 = 328 // 0x0148
	TagDecreaseVolume   uint16 = 329 // 0x0149

	// Connection constants
	connectTimeout = 2 * time.Second
	readTimeout    = 3 * time.Second

	bufferSize = 655360

	DEBUG = false
)

// TimeoutError represents a timeout error
type TimeoutError struct {
	message string
}

func (e *TimeoutError) Error() string {
	return e.message
}

// IsTimeoutError checks if an error is a TimeoutError
func IsTimeoutError(err error) bool {
	_, ok := err.(*TimeoutError)
	return ok
}

// NewTimeoutError creates a new TimeoutError
func NewTimeoutError(message string) error {
	return &TimeoutError{message: message}
}

// Program represents a program in the player
type Program struct {
	Index   uint32
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

// CurrentProgramResponse represents the response of GetCurrentProgram
type CurrentProgramResponse struct {
	Success   bool   // 执行结果，true表示成功，false表示失败
	ProgramID int32  // 节目ID，-1表示没有节目
	ProgState uint32 // 节目状态：0=播放，1=暂停，2=停止
}

// LayerVolumeMuteStatusResponse 表示查询图层音量和静音状态的响应结构
//
// 响应格式:
//
//	CC 55 CC 55 + header(12) + TLV头(4) + Bsucceed(1) + LayerIndex(2) + Volume(1) + MuteFlag(1)
type LayerVolumeMuteStatusResponse struct {
	Success    bool   // 执行结果，true为成功，false为失败
	LayerIndex uint16 // 图层索引
	Volume     uint8  // 音量大小
	MuteFlag   uint8  // 是否静音，1为静音，0为非静音
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

// sendCommandWithTimeout sends command and returns response with timeout
func (c *PlayerClient) sendCommandWithTimeout(tag uint16, data []byte, timeout time.Duration) ([]byte, error) {
	// 创建连接
	conn, err := net.DialTimeout("tcp", c.addr, connectTimeout)
	if err != nil {
		return nil, fmt.Errorf("connect failed: %v", err)
	}
	defer conn.Close()
	if timeout > 0 {
		// 设置读取超时
		err = conn.SetReadDeadline(time.Now().Add(timeout))
		if err != nil {
			return nil, fmt.Errorf("set read deadline failed: %v", err)
		}
	}

	// 构建数据包
	packet := c.buildPacket(tag, data)
	if DEBUG {
		fmt.Printf("Send packet: %s\n", bytesToHexString(packet))
	}

	// 发送数据
	_, err = conn.Write(packet)
	if err != nil {
		return nil, fmt.Errorf("write failed: %v", err)
	}

	// 如果timeout为0，则不等待响应，延时500ms后直接返回
	if timeout == 0 {
		if DEBUG {
			fmt.Printf("Timeout is 0, not waiting for response\n")
		}
		time.Sleep(500 * time.Millisecond)
		return nil, nil
	}

	// 为大型响应命令预分配更大的缓冲区
	var buffer []byte
	if tag == TagGetProgramList {
		buffer = make([]byte, bufferSize*2) // 为大型响应分配更大的缓冲区
	} else {
		buffer = make([]byte, bufferSize)
	}

	totalRead := 0
	lastDataTime := time.Now()
	receiveStartTime := time.Now()

	// 设置绝对超时，防止无限等待
	absoluteTimeout := time.Now().Add(timeout)

	// 读取响应数据，可能需要多次读取
	for {
		// 检查是否达到绝对超时
		if time.Now().After(absoluteTimeout) {
			return nil, NewTimeoutError(fmt.Sprintf("absolute timeout reached after %v", timeout))
		}

		// 设置读取超时
		readDeadline := time.Now().Add(500 * time.Millisecond)
		if readDeadline.After(absoluteTimeout) {
			readDeadline = absoluteTimeout
		}
		conn.SetReadDeadline(readDeadline)

		// 读取数据
		n, err := conn.Read(buffer[totalRead:])
		if err != nil {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				// 读取超时，检查是否已完成数据接收
				if totalRead > 0 && time.Since(lastDataTime) > 500*time.Millisecond {
					// 特殊命令可能要接收大量数据，需要特殊处理
					if tag == TagGetProgramList {
						// 这些命令可能返回多个包，需要等待更长时间
						// 检查是否有足够的完整包
						if isEnoughDataForCommand(buffer[:totalRead], tag) {
							if DEBUG {
								fmt.Printf("Received enough data for the command after %v\n", time.Since(receiveStartTime))
							}
							break
						}

						// 如果还没有接收到足够的数据，但已经超过一定时间没有新数据，也认为接收完成
						if time.Since(lastDataTime) > 2*time.Second {
							if DEBUG {
								fmt.Printf("No more data received for %v, assuming completed\n", time.Since(lastDataTime))
							}
							break
						}

						// 继续等待更多数据
						continue
					}
					// 对于其他命令，如果已经接收到数据且超过500ms没有新数据，认为接收完成
					break
				}
				// 未接收到数据或者距离上次接收数据时间不够长，继续等待
				continue
			}
			// 如果已经读取了一些数据，返回已读取的数据
			if totalRead > 0 {
				break
			}
			return nil, fmt.Errorf("read failed: %v", err)
		}

		totalRead += n

		// 检查缓冲区容量，需要时扩展
		if totalRead >= len(buffer)-4096 {
			// 扩展缓冲区
			newBuffer := make([]byte, len(buffer)*2)
			copy(newBuffer, buffer)
			buffer = newBuffer
		}

		// 更新最后接收数据的时间
		lastDataTime = time.Now()

		// 检查是否已接收完整数据
		// 对于简单命令，如果已读取到足够的数据则认为完成
		if totalRead >= 16 && !isBigDataCommand(tag) {
			// 检查TLV长度是否已接收完整
			if isCompletePacket(buffer[:totalRead]) {
				break
			}
		}
	}

	// 输出调试信息
	if totalRead < 16 {
		return nil, fmt.Errorf("response too short: %d bytes", totalRead)
	}

	// 解析帧头和TLV信息
	packetHeader := binary.LittleEndian.Uint32(buffer[0:4])
	packetType := binary.LittleEndian.Uint16(buffer[4:6])
	protocolVersion := binary.LittleEndian.Uint16(buffer[6:8])
	sequence := binary.LittleEndian.Uint16(buffer[8:10])
	contentLength := binary.LittleEndian.Uint16(buffer[10:12])
	tlvTag := binary.LittleEndian.Uint16(buffer[12:14])
	tlvLength := binary.LittleEndian.Uint16(buffer[14:16])

	if DEBUG {
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
	}

	return buffer[:totalRead], nil
}

// isBigDataCommand 判断命令是否可能返回大量数据
func isBigDataCommand(tag uint16) bool {
	return tag == TagGetProgramList
}

// isCompletePacket 检查数据包是否接收完整
func isCompletePacket(data []byte) bool {
	if len(data) < 16 {
		return false
	}

	contentLength := binary.LittleEndian.Uint16(data[10:12])
	return len(data) >= int(contentLength)+12 // 12是包头长度
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
		if DEBUG {
			fmt.Printf("Program count in response: %d\n", programCount)
		}

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

// GetProgramList gets the list of programs from the player
func (c *PlayerClient) GetProgramList() (*ProgramListResponse, error) {
	// 设置超时时间
	timeout := 30 * time.Second
	var resp []byte
	var err error

	// 尝试最多3次获取程序列表
	for retries := 0; retries < 3; retries++ {
		resp, err = c.sendCommandWithTimeout(TagGetProgramList, nil, timeout)
		if err != nil {
			if DEBUG {
				fmt.Printf("Failed to get program list (attempt %d): %v\n", retries+1, err)
			}
			time.Sleep(time.Second)
			continue
		}

		// 验证响应长度是否足够
		if len(resp) < 20 {
			if DEBUG {
				fmt.Printf("Response too short (attempt %d): %d bytes\n", retries+1, len(resp))
			}
			time.Sleep(time.Second)
			continue
		}

		break
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get program list after retries: %v", err)
	}

	// 定义节目数据结构大小
	const (
		programHeaderSize  = 16
		programCountSize   = 4                                                                                                              // 包头大小
		programIndexSize   = 4                                                                                                              // 索引大小
		programIdSize      = 4                                                                                                              // ID大小
		programNameSize    = 128                                                                                                            // 名称大小
		programIsEmptySize = 1                                                                                                              // isEmpty标志大小
		programSize        = programHeaderSize + programCountSize + programIndexSize + programIdSize + programNameSize + programIsEmptySize // 157字节
	)
	// 创建结果结构
	result := &ProgramListResponse{
		TotalCount: 0,
		Programs:   make([]Program, 0, 0),
	}

	for i := 0; i < len(resp); i += programSize {
		if DEBUG {
			fmt.Printf("Program %d: %s\n", i, bytesToHexString(resp[i:i+programSize]))
		}
		// 节目数据开始位置（跳过包头）
		dataStart := i + programHeaderSize + programCountSize

		// 解析索引
		programIndex := binary.LittleEndian.Uint32(resp[dataStart : dataStart+programIndexSize])

		// 解析ID
		idPos := dataStart + programIndexSize
		programId := binary.LittleEndian.Uint32(resp[idPos : idPos+programIdSize])

		// 解析名称
		namePos := idPos + programIdSize
		nameBytes := resp[namePos : namePos+programNameSize]
		name := string(bytes.TrimRight(nameBytes, "\x00"))

		// 解析isEmpty（0表示空，1表示非空）
		isEmptyPos := namePos + programNameSize
		isEmptyByte := resp[isEmptyPos]
		isEmpty := isEmptyByte == 0

		// 打印调试信息
		if DEBUG {
			fmt.Printf("Program %d raw details: Index=%d, ID=%d, isEmpty byte=0x%02X at position %d, isEmpty=%v\n",
				i, programIndex, programId, isEmptyByte, isEmptyPos, isEmpty)
		}

		// 创建程序对象
		program := Program{
			Index:   programIndex,
			ID:      programId,
			Name:    name,
			IsEmpty: isEmpty,
		}

		// 添加到结果列表
		result.Programs = append(result.Programs, program)

	}

	return result, nil
}

// FadeProgram fades to the specified program
func (c *PlayerClient) FadeProgram(programID uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, programID)
	_, err := c.sendCommandWithTimeout(TagFadeProgram, data, 0)
	return err
}

// CutProgram cuts to the specified program
func (c *PlayerClient) CutProgram(programID uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, programID)
	_, err := c.sendCommandWithTimeout(TagCutProgram, data, 0)
	return err
}

// PauseProgram pauses the current program
func (c *PlayerClient) PauseProgram(programID uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, programID)
	_, err := c.sendCommandWithTimeout(TagPauseProgram, data, 0)
	return err
}

// PlayProgram plays the specified program
func (c *PlayerClient) PlayProgram(programID uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, programID)
	_, err := c.sendCommandWithTimeout(TagPlayProgram, data, 0)
	return err
}

// StopProgram stops the specified program
func (c *PlayerClient) StopProgram(programID uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, programID)
	_, err := c.sendCommandWithTimeout(TagStopProgram, data, 0)
	return err
}

// OpenGlobalSound opens the global sound
func (c *PlayerClient) OpenGlobalSound() error {
	_, err := c.sendCommandWithTimeout(TagOpenGlobalSound, nil, 0)
	return err
}

// CloseGlobalSound closes the global sound
func (c *PlayerClient) CloseGlobalSound() error {
	_, err := c.sendCommandWithTimeout(TagCloseGlobalSound, nil, 0)
	return err
}

// SetGlobalVolume sets the global volume
func (c *PlayerClient) SetGlobalVolume(volume uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, volume)
	_, err := c.sendCommandWithTimeout(TagSetGlobalVolume, data, readTimeout)
	return err
}

// IncreaseGlobalVolume increases the global volume
func (c *PlayerClient) IncreaseGlobalVolume(step uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, step)
	_, err := c.sendCommandWithTimeout(TagIncreaseVolume, data, readTimeout)
	return err
}

// DecreaseGlobalVolume decreases the global volume
func (c *PlayerClient) DecreaseGlobalVolume(step uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, step)
	_, err := c.sendCommandWithTimeout(TagDecreaseVolume, data, readTimeout)
	return err
}

// GetAllProgramMedia gets the media list of all programs
func (c *PlayerClient) GetAllProgramMedia(programID uint32) (*MediaListResponse, error) {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, programID)

	resp, err := c.sendCommandWithTimeout(TagGetAllProgramMedia, data, readTimeout)
	if err != nil {
		return nil, err
	}

	// 检查响应的有效性
	if len(resp) < 16+4 {
		return nil, fmt.Errorf("get all program media response too short: %d bytes", len(resp))
	}

	result := &MediaListResponse{
		ProgramID: programID,
		Media:     make([]Media, 0),
	}

	// 获取实际的TLV长度
	// Each media item has: MediaId(4 bytes) + MediaName(32 bytes)
	const (
		headerSize    = 12 // CC 55 CC 55 + version + sequence + tag + length
		tlvHeaderSize = 4  // TLV tag(2) + length(2)
		mediaIdSize   = 4
		mediaNameSize = 300
		mediaItemSize = mediaIdSize + mediaNameSize
	)

	if len(resp) < headerSize+tlvHeaderSize {
		return nil, fmt.Errorf("response too short, len=%d", len(resp))
	}

	// Get TLV length
	tlvLength := binary.LittleEndian.Uint16(resp[14:16])
	if DEBUG {
		fmt.Printf("TLV length: %d bytes\n", tlvLength)
	}

	// 这样可以适应不同的服务器实现
	dataSize := int(tlvLength)

	if dataSize < 0 {
		return nil, fmt.Errorf("invalid dataSize: %d", dataSize)
	}

	// 通过TLV长度计算媒体项数量，取最小值以防止越界
	mediaCount := dataSize / mediaItemSize
	if mediaCount*mediaItemSize > dataSize {
		mediaCount-- // 如果有不完整的项，则减少计数
	}

	if DEBUG {
		fmt.Printf("Media count based on dataSize (%d) and mediaItemSize (%d): %d\n",
			dataSize, mediaItemSize, mediaCount)
	}

	// 解析每个媒体项
	offset := headerSize + tlvHeaderSize
	for i := 0; i < mediaCount && offset+mediaItemSize <= len(resp); i++ {
		// 读取媒体ID
		mediaId := binary.LittleEndian.Uint32(resp[offset : offset+mediaIdSize])

		// 读取媒体名称
		nameOffset := offset + mediaIdSize
		nameBytes := resp[nameOffset : nameOffset+mediaNameSize]

		// 去除尾部的NULL字符
		var name string
		nullPos := bytes.IndexByte(nameBytes, 0)
		if nullPos != -1 {
			name = string(nameBytes[:nullPos])
		} else {
			name = string(bytes.TrimRight(nameBytes, "\x00"))
		}

		// 创建媒体项
		media := Media{
			ID:   mediaId,
			Name: name,
		}
		if DEBUG {
			fmt.Printf("Parsed media: ID=%d, Name=%s\n", media.ID, media.Name)
		}

		// 添加到结果
		result.Media = append(result.Media, media)

		// 移动到下一个媒体项
		offset += mediaItemSize
	}

	return result, nil
}

// PlayLayerMedia plays media in the specified layer of current program
func (c *PlayerClient) PlayLayerMedia(layerIndex uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, layerIndex)
	_, err := c.sendCommandWithTimeout(TagPlayLayerMedia, data, readTimeout)
	return err
}

// PauseLayerMedia pauses media in the specified layer of current program
func (c *PlayerClient) PauseLayerMedia(layerIndex uint32) error {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, layerIndex)
	_, err := c.sendCommandWithTimeout(TagPauseLayerMedia, data, readTimeout)
	return err
}

// MuteMediaSound mutes sound for media at specified position
func (c *PlayerClient) MuteMediaSound(col, row uint32) error {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint32(data[0:4], col)
	binary.LittleEndian.PutUint32(data[4:8], row)
	_, err := c.sendCommandWithTimeout(TagMuteMediaSound, data, readTimeout)
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

	_, err := c.sendCommandWithTimeout(TagControlLayerProgress, data, readTimeout)
	return err
}

// QueryLayerProgress queries playback progress of specified layer
func (c *PlayerClient) QueryLayerProgress(layerIndex uint16) (*LayerProgressResponse, error) {
	data := make([]byte, 2)
	binary.LittleEndian.PutUint16(data, layerIndex)

	resp, err := c.sendCommandWithTimeout(TagQueryLayerProgress, data, readTimeout)
	if err != nil {
		return nil, fmt.Errorf("QueryLayerProgress failed: %v", err)
	}

	if len(resp) < 16+1+2+4+4 {
		return nil, fmt.Errorf("response too short: %d bytes", len(resp))
	}

	result := &LayerProgressResponse{
		LayerIndex: layerIndex,
	}

	// 响应格式: CC 55 CC 55 + header(12) + TLV头(4) + Success(1) + LayerIndex(2) + RemainTime(4) + TotalTime(4)
	// Success字段位于响应的第17个字节
	result.Success = resp[16] != 0

	// LayerIndex位于Success之后
	result.LayerIndex = binary.LittleEndian.Uint16(resp[17:19])

	// RemainTime位于LayerIndex之后
	result.RemainTime = binary.LittleEndian.Uint32(resp[19:23])

	// TotalTime位于RemainTime之后
	result.TotalTime = binary.LittleEndian.Uint32(resp[23:27])

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

// GetCurrentProgram queries the current playing program ID and state
func (c *PlayerClient) GetCurrentProgram() (*CurrentProgramResponse, error) {
	// 发送获取当前节目的命令，不需要参数
	resp, err := c.sendCommandWithTimeout(TagGetCurrentProgram, nil, readTimeout)
	if err != nil {
		return nil, fmt.Errorf("get current program failed: %v", err)
	}

	// 检查响应的有效性
	// 响应格式: CC 55 CC 55 + header(12) + TLV头(4) + Success(1) + ProgramID(4) + ProgState(4)
	if len(resp) < 16+1+4+4 {
		return nil, fmt.Errorf("response too short: %d bytes", len(resp))
	}

	// 创建响应结构体
	result := &CurrentProgramResponse{}

	// 解析Success字段（第17个字节，索引16）
	result.Success = resp[16] != 0

	// 解析ProgramID（4字节，位于Success之后）
	result.ProgramID = int32(binary.LittleEndian.Uint32(resp[17:21]))

	// 解析ProgState（4字节，位于ProgramID之后）
	result.ProgState = binary.LittleEndian.Uint32(resp[21:25])

	if DEBUG {
		fmt.Printf("Current Program: Success=%v, ProgramID=%d, State=%d\n",
			result.Success, result.ProgramID, result.ProgState)
	}

	return result, nil
}

// QueryLayerVolumeMuteStatus 查询指定图层的音量和静音状态
// 参数: layerIndex 图层索引（从0开始）
// 返回: LayerVolumeMuteStatusResponse, error
func (c *PlayerClient) QueryLayerVolumeMuteStatus(layerIndex uint16) (*LayerVolumeMuteStatusResponse, error) {
	const TagQueryLayerVolumeMuteStatus uint16 = 322 // 0x0142
	data := make([]byte, 2)
	binary.LittleEndian.PutUint16(data, layerIndex)

	resp, err := c.sendCommandWithTimeout(TagQueryLayerVolumeMuteStatus, data, readTimeout)
	if err != nil {
		return nil, err
	}

	// 响应格式: CC 55 CC 55 + header(12) + TLV头(4) + Bsucceed(1) + LayerIndex(2) + Volume(1) + MuteFlag(1)
	if len(resp) < 16+1+2+1+1 {
		return nil, fmt.Errorf("response too short: %d bytes", len(resp))
	}

	result := &LayerVolumeMuteStatusResponse{}
	result.Success = resp[16] != 0
	result.LayerIndex = binary.LittleEndian.Uint16(resp[17:19])
	result.Volume = resp[19]
	result.MuteFlag = resp[20]

	return result, nil
}
