package client

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
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

	// Connection constants
	reconnectInterval = 5 * time.Second
	connectTimeout    = 5 * time.Second

	bufferSize = 655360
)

// Program represents a program in the player
type Program struct {
	ID      uint32
	Name    string
	IsEmpty bool
}

// ProgramListResponse represents response of GetProgramList
type ProgramListResponse struct {
	TotalCount int
	Programs   []Program
}

// PlayerClient represents a client to control media player
type PlayerClient struct {
	addr      string
	conn      net.Conn
	sequence  uint16
	mu        sync.RWMutex
	closed    bool
	connected bool
	connChan  chan bool // Channel to signal connection status

	// Channels for async communication
	sendChan    chan *commandRequest
	responseBuf *responseBuffer
}

// commandRequest represents a command to be sent
type commandRequest struct {
	tag      uint16
	data     []byte
	respChan chan *commandResponse
}

// commandResponse represents a response from the server
type commandResponse struct {
	data []byte
	err  error
}

// responseBuffer manages received data
type responseBuffer struct {
	mu       sync.Mutex
	data     []byte
	ready    chan struct{}
	lastRead time.Time
}

// newResponseBuffer creates a new response buffer
func newResponseBuffer() *responseBuffer {
	return &responseBuffer{
		ready:    make(chan struct{}, 1),
		lastRead: time.Now(),
	}
}

// append adds data to the buffer
func (rb *responseBuffer) append(data []byte) {
	rb.mu.Lock()
	rb.data = append(rb.data, data...)
	rb.lastRead = time.Now()
	rb.mu.Unlock()
}

// isComplete checks if the response is complete (no new data for 1 second)
func (rb *responseBuffer) isComplete() bool {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	return len(rb.data) > 0 && time.Since(rb.lastRead) >= time.Second
}

// get retrieves and clears the buffer
func (rb *responseBuffer) get() []byte {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	data := rb.data
	rb.data = nil
	rb.lastRead = time.Now()
	return data
}

// NewTCPClient creates a new TCP client
func NewTCPClient(addr string) (*PlayerClient, error) {
	client := &PlayerClient{
		addr:        addr,
		connChan:    make(chan bool, 1),
		sendChan:    make(chan *commandRequest),
		responseBuf: newResponseBuffer(),
	}

	// Try initial connection
	if err := client.connect(); err != nil {
		fmt.Printf("Initial connection failed: %v, will retry in background\n", err)
	}

	// Start connection manager goroutine
	go client.maintainConnection()

	// Start send/receive goroutines
	go client.sendLoop()
	go client.receiveLoop()

	return client, nil
}

// connect attempts to establish a TCP connection
func (c *PlayerClient) connect() error {
	conn, err := net.DialTimeout("tcp", c.addr, connectTimeout)
	if err != nil {
		return err
	}

	c.mu.Lock()
	c.conn = conn
	c.connected = true
	c.mu.Unlock()

	// Signal connection success
	select {
	case c.connChan <- true:
	default:
	}

	return nil
}

// maintainConnection maintains TCP connection and handles reconnection
func (c *PlayerClient) maintainConnection() {
	for {
		if c.closed {
			return
		}

		c.mu.RLock()
		connected := c.connected
		c.mu.RUnlock()

		if !connected {
			if err := c.connect(); err != nil {
				time.Sleep(reconnectInterval)
			}
		} else {
			// If connected, just wait for failure signal
			time.Sleep(reconnectInterval)
		}
	}
}

// handleConnectionFailure handles connection failure
func (c *PlayerClient) handleConnectionFailure() {
	c.mu.Lock()
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
	c.connected = false
	c.mu.Unlock()
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

	// Write header
	binary.LittleEndian.PutUint32(packet[0:], ProtocolHeader)
	binary.LittleEndian.PutUint16(packet[4:], 1)               // Packet type should be 1
	binary.LittleEndian.PutUint16(packet[6:], ProtocolVersion) // Protocol version
	binary.LittleEndian.PutUint16(packet[8:], c.sequence)      // Sequence number
	binary.LittleEndian.PutUint16(packet[10:], uint16(4+contentLen))

	// Write TLV
	binary.LittleEndian.PutUint16(packet[12:], tag)
	binary.LittleEndian.PutUint16(packet[14:], uint16(contentLen))

	// Write data if any
	if contentLen > 0 {
		copy(packet[16:], data)
	}

	c.sequence++
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

// sendLoop handles sending commands
func (c *PlayerClient) sendLoop() {
	for {
		select {
		case req := <-c.sendChan:
			if req == nil {
				return // Channel closed
			}

			c.mu.RLock()
			conn := c.conn
			connected := c.connected
			c.mu.RUnlock()

			if !connected || conn == nil {
				req.respChan <- &commandResponse{err: errors.New("not connected")}
				continue
			}

			packet := c.buildPacket(req.tag, req.data)
			fmt.Printf("Send packet: %s\n", bytesToHexString(packet))

			_, err := conn.Write(packet)
			if err != nil {
				c.handleConnectionFailure()
				req.respChan <- &commandResponse{err: err}
				continue
			}

			// Wait for response
			select {
			case <-c.responseBuf.ready:
				resp := c.responseBuf.get()
				req.respChan <- &commandResponse{data: resp}
			case <-time.After(10 * time.Second): // Increased timeout for large responses
				req.respChan <- &commandResponse{err: errors.New("response timeout")}
			}
		}
	}
}

// receiveLoop handles receiving data
func (c *PlayerClient) receiveLoop() {
	tempBuf := make([]byte, bufferSize)
	checkTicker := time.NewTicker(100 * time.Millisecond) // Check completion every 100ms
	defer checkTicker.Stop()

	for {
		if c.closed {
			return
		}

		c.mu.RLock()
		conn := c.conn
		connected := c.connected
		c.mu.RUnlock()

		if !connected || conn == nil {
			time.Sleep(time.Second)
			continue
		}

		// Set a read deadline to prevent blocking forever
		if tcpConn, ok := conn.(*net.TCPConn); ok {
			tcpConn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
		}

		select {
		case <-checkTicker.C:
			// Check if we have a complete response
			if c.responseBuf.isComplete() {
				select {
				case c.responseBuf.ready <- struct{}{}:
				default:
				}
			}
		default:
			n, err := conn.Read(tempBuf)
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					// Timeout is expected, continue checking for completion
					continue
				}
				c.handleConnectionFailure()
				continue
			}

			if n > 0 {
				c.responseBuf.append(tempBuf[:n])
				fmt.Printf("Received %d bytes\n", n)
			}
		}
	}
}

// sendCommand sends command and returns response
func (c *PlayerClient) sendCommand(tag uint16, data []byte) ([]byte, error) {
	respChan := make(chan *commandResponse, 1)
	req := &commandRequest{
		tag:      tag,
		data:     data,
		respChan: respChan,
	}

	// Send request
	c.sendChan <- req

	// Wait for response with timeout
	select {
	case resp := <-respChan:
		if resp.err != nil {
			return nil, fmt.Errorf("command failed: %v", resp.err)
		}

		// Verify response
		if len(resp.data) < 16 {
			return nil, errors.New("response too short")
		}

		packetHeader := binary.LittleEndian.Uint32(resp.data[0:4])
		tlvTag := binary.LittleEndian.Uint16(resp.data[12:14])
		fmt.Printf("Packet verification: header=0x%08X, TLV tag=%d\n", packetHeader, tlvTag)

		return resp.data, nil
	case <-time.After(10 * time.Second): // Overall timeout
		return nil, errors.New("command timeout")
	}
}

// GetProgramList gets the program list
func (c *PlayerClient) GetProgramList() (*ProgramListResponse, error) {
	// Send empty data for query
	resp, err := c.sendCommand(TagGetProgramList, nil)
	if err != nil {
		return nil, fmt.Errorf("GetProgramList failed: %v", err)
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
		return nil, fmt.Errorf("response too short for header and program count")
	}

	// Get program count from first packet
	programCount := int(binary.LittleEndian.Uint16(resp[16:18]))
	fmt.Printf("Program count: %d\n", programCount)

	// Create result structure
	result := &ProgramListResponse{
		TotalCount: programCount,
		Programs:   make([]Program, 0, programCount),
	}

	// Process each packet
	offset := 0
	for i := 0; i < programCount; i++ {
		// Find next packet start
		for offset < len(resp)-4 {
			if resp[offset] == 0xCC && resp[offset+1] == 0x55 &&
				resp[offset+2] == 0xCC && resp[offset+3] == 0x55 {
				break
			}
			offset++
		}

		// Verify we have enough data for a complete packet
		packetStart := offset
		if packetStart+headerSize+tlvHeaderSize+programCountSize+
			programIndexSize+programIdSize+programNameSize+isEmptySize > len(resp) {
			return nil, fmt.Errorf("incomplete packet data at program %d", i)
		}

		// Skip header and TLV header
		dataStart := packetStart + headerSize + tlvHeaderSize + programCountSize

		// Read program data
		program := Program{
			ID:      binary.LittleEndian.Uint32(resp[dataStart+programIndexSize : dataStart+programIndexSize+programIdSize]),
			IsEmpty: resp[dataStart+programIndexSize+programIdSize+programNameSize] == 0,
		}

		// Read null-terminated program name
		nameBytes := resp[dataStart+programIndexSize+programIdSize : dataStart+programIndexSize+programIdSize+programNameSize]
		nullPos := bytes.IndexByte(nameBytes, 0)
		if nullPos != -1 {
			program.Name = string(nameBytes[:nullPos])
		} else {
			program.Name = string(bytes.TrimRight(nameBytes, "\x00"))
		}

		result.Programs = append(result.Programs, program)

		// Move to next packet
		offset = packetStart + headerSize + tlvHeaderSize + programCountSize +
			programIndexSize + programIdSize + programNameSize + isEmptySize
	}

	return result, nil
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

// Close closes the connection and stops all goroutines
func (c *PlayerClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.closed = true
	close(c.sendChan) // Signal sendLoop to stop

	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
