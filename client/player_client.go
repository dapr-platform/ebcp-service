package client

import (
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
	ProtocolVersion uint16 = 0x0100

	// Command tags
	TagPlayCurrent   uint16 = 24000
	TagPlayPrev      uint16 = 24002
	TagPlayNext      uint16 = 24003
	TagPlayByIndex   uint16 = 24004
	TagPause         uint16 = 24005
	TagResume        uint16 = 24006
	TagStop          uint16 = 24007
	TagSetVolume     uint16 = 24014
	TagSetWindow     uint16 = 24016
	TagSetVisibility uint16 = 24018

	// Connection constants
	reconnectInterval = 5 * time.Second
	connectTimeout    = 5 * time.Second
)

// PlayerClient represents a client to control media player
type PlayerClient struct {
	addr      string
	conn      net.Conn
	sequence  uint16
	mu        sync.RWMutex
	closed    bool
	connected bool
	connChan  chan bool // Channel to signal connection status
}

// NewTCPClient creates a new TCP client
func NewTCPClient(addr string) (*PlayerClient, error) {
	client := &PlayerClient{
		addr:     addr,
		connChan: make(chan bool, 1),
	}

	// Try initial connection
	if err := client.connect(); err != nil {
		fmt.Printf("Initial connection failed: %v, will retry in background\n", err)
	}

	// Start connection manager goroutine
	go client.maintainConnection()

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

// waitForConnection waits for connection to be established
func (c *PlayerClient) waitForConnection(timeout time.Duration) error {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case <-c.connChan:
		return nil
	case <-timer.C:
		return errors.New("connection timeout")
	}
}

// Close closes the connection
func (c *PlayerClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.closed = true
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
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

// sendCommand sends command and returns response
func (c *PlayerClient) sendCommand(tag uint16, data []byte) (bool, uint64, error) {
	// Set timeout for the entire operation
	deadline := time.Now().Add(connectTimeout)

	for {
		if time.Now().After(deadline) {
			return false, 0, errors.New("operation timeout")
		}

		c.mu.RLock()
		conn := c.conn
		connected := c.connected
		c.mu.RUnlock()

		// If not connected, wait for connection
		if !connected || conn == nil {
			time.Sleep(time.Second)
			continue
		}

		// Set deadline for this attempt
		if tcpConn, ok := conn.(*net.TCPConn); ok {
			tcpConn.SetDeadline(time.Now().Add(time.Second * 3))
		}

		packet := c.buildPacket(tag, data)
		fmt.Printf("Send packet: %s\n", bytesToHexString(packet))

		_, err := conn.Write(packet)
		if err != nil {
			c.handleConnectionFailure()
			continue
		}

		// Read response
		resp := make([]byte, 1024)
		n, err := conn.Read(resp)
		if err != nil {
			c.handleConnectionFailure()
			continue
		}

		fmt.Printf("Received packet: %s\n", bytesToHexString(resp[:n]))

		if n < 17 { // Minimum response size
			c.handleConnectionFailure()
			continue
		}

		success := resp[16] == 1
		var resourceID uint64
		if n >= 25 {
			resourceID = binary.LittleEndian.Uint64(resp[17:25])
		}

		return success, resourceID, nil
	}
}

// PlayCurrent plays current resource
func (c *PlayerClient) PlayCurrent() (bool, uint64, error) {
	return c.sendCommand(TagPlayCurrent, nil)
}

// PlayPrev plays previous resource
func (c *PlayerClient) PlayPrev() (bool, uint64, error) {
	return c.sendCommand(TagPlayPrev, nil)
}

// PlayNext plays next resource
func (c *PlayerClient) PlayNext() (bool, uint64, error) {
	return c.sendCommand(TagPlayNext, nil)
}

// PlayByIndex plays resource by index
func (c *PlayerClient) PlayByIndex(index uint64) (bool, uint64, error) {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, index)
	return c.sendCommand(TagPlayByIndex, data)
}

// Pause pauses current playback
func (c *PlayerClient) Pause() (bool, uint64, error) {
	return c.sendCommand(TagPause, nil)
}

// Resume resumes playback
func (c *PlayerClient) Resume() (bool, uint64, error) {
	return c.sendCommand(TagResume, nil)
}

// Stop stops playback
func (c *PlayerClient) Stop() (bool, uint64, error) {
	return c.sendCommand(TagStop, nil)
}

// SetVolume sets volume level (0-100)
func (c *PlayerClient) SetVolume(volume uint16) (bool, error) {
	if volume > 100 {
		return false, errors.New("volume must be between 0 and 100")
	}

	data := make([]byte, 2)
	binary.LittleEndian.PutUint16(data, volume)

	success, _, err := c.sendCommand(TagSetVolume, data)
	return success, err
}

// SetWindow sets window position and size
func (c *PlayerClient) SetWindow(x, y, width, height uint16, fullscreen bool) (bool, error) {
	data := make([]byte, 9)
	binary.LittleEndian.PutUint16(data[0:], x)
	binary.LittleEndian.PutUint16(data[2:], y)
	binary.LittleEndian.PutUint16(data[4:], width)
	binary.LittleEndian.PutUint16(data[6:], height)
	if fullscreen {
		data[8] = 1
	}

	success, _, err := c.sendCommand(TagSetWindow, data)
	return success, err
}

// SetVisibility sets window visibility
func (c *PlayerClient) SetVisibility(visible bool) (bool, error) {
	data := make([]byte, 2)
	if visible {
		binary.LittleEndian.PutUint16(data, 1)
	}

	success, _, err := c.sendCommand(TagSetVisibility, data)
	return success, err
}
