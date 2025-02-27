package client

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

type CentralControlClient struct {
	addr      string
	conn      *net.UDPConn
	sendChan  chan []byte
	recvChan  chan []byte
	respChans map[string]chan []byte
	mu        sync.RWMutex
}

func NewCentralControlClient(host string, port int) (*CentralControlClient, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return nil, err
	}

	client := &CentralControlClient{
		addr:      addr,
		conn:      conn,
		sendChan:  make(chan []byte, 100),
		recvChan:  make(chan []byte, 100),
		respChans: make(map[string]chan []byte),
	}

	go client.sendLoop()
	go client.receiveLoop()

	return client, nil
}

func (c *CentralControlClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *CentralControlClient) sendLoop() {
	for data := range c.sendChan {
		_, err := c.conn.Write(data)
		if err != nil {
			fmt.Printf("Failed to send data: %v\n", err)
		}
	}
}

func (c *CentralControlClient) receiveLoop() {
	buf := make([]byte, 1024)
	for {
		n, err := c.conn.Read(buf)
		if err != nil {
			fmt.Printf("Failed to receive data: %v\n", err)
			continue
		}

		data := make([]byte, n)
		copy(data, buf[:n])
		c.recvChan <- data

		// Handle response
		c.mu.RLock()
		for cmd, ch := range c.respChans {
			select {
			case ch <- data:
			default:
			}
			delete(c.respChans, cmd)
		}
		c.mu.RUnlock()
	}
}

func (c *CentralControlClient) SendCommand(cmd string) ([]byte, error) {
	respChan := make(chan []byte, 1)
	
	c.mu.Lock()
	c.respChans[cmd] = respChan
	c.mu.Unlock()

	// Send command
	c.sendChan <- []byte(cmd)

	// Wait for response with timeout
	select {
	case resp := <-respChan:
		fmt.Printf("Received response: %s\n", string(resp))
		return resp, nil
	case <-time.After(5 * time.Second):
		c.mu.Lock()
		delete(c.respChans, cmd)
		c.mu.Unlock()
		return nil, errors.New("command timeout")
	}
}

// Predefined commands
func (c *CentralControlClient) AllPowerOn() ([]byte, error) {
	return c.SendCommand("A1")
}

func (c *CentralControlClient) AllPowerOff() ([]byte, error) {
	return c.SendCommand("A2") 
}

func (c *CentralControlClient) SequencePowerOn() ([]byte, error) {
	return c.SendCommand("A3")
}

func (c *CentralControlClient) SequencePowerOff() ([]byte, error) {
	return c.SendCommand("A4")
}

func (c *CentralControlClient) LEDScreenOn() ([]byte, error) {
	return c.SendCommand("A9")
}

func (c *CentralControlClient) LEDScreenOff() ([]byte, error) {
	return c.SendCommand("A10")
}
