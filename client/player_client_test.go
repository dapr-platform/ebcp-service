package client

import (
	"fmt"
	"testing"
	"time"
)

func TestNewPlayerClient(t *testing.T) {
	client, err := NewTCPClient("182.92.117.41:40306")
	if err != nil {
		t.Fatalf("Failed to create TCP client: %v", err)
	}
	defer client.Close()

	// Wait for connection
	time.Sleep(time.Second)
}

func TestGetProgramList(t *testing.T) {
	client, err := NewTCPClient("182.92.117.41:40306")
	if err != nil {
		t.Fatalf("Failed to create TCP client: %v", err)
	}
	defer client.Close()

	// Wait for connection
	time.Sleep(time.Second)

	resp, err := client.GetProgramList()
	if err != nil {
		t.Errorf("GetProgramList failed: %v", err)
	}
	if resp == nil || len(resp.Programs) == 0 {
		t.Error("GetProgramList returned empty response")
	}
	fmt.Printf("GetProgramList response: total=%d, programs=%+v\n", resp.TotalCount, resp.Programs)
}

func TestFadeProgram(t *testing.T) {
	client, err := NewTCPClient("182.92.117.41:40306")
	if err != nil {
		t.Fatalf("Failed to create TCP client: %v", err)
	}
	defer client.Close()

	// Wait for connection
	time.Sleep(time.Second)

	err = client.FadeProgram(1)
	if err != nil {
		t.Errorf("FadeProgram failed: %v", err)
	}
	fmt.Println("FadeProgram executed successfully")
}

func TestCutProgram(t *testing.T) {
	client, err := NewTCPClient("182.92.117.41:40306")
	if err != nil {
		t.Fatalf("Failed to create TCP client: %v", err)
	}
	defer client.Close()

	// Wait for connection
	time.Sleep(time.Second)

	err = client.CutProgram(1)
	if err != nil {
		t.Errorf("CutProgram failed: %v", err)
	}
	fmt.Println("CutProgram executed successfully")
}

func TestPauseProgram(t *testing.T) {
	client, err := NewTCPClient("182.92.117.41:40306")
	if err != nil {
		t.Fatalf("Failed to create TCP client: %v", err)
	}
	defer client.Close()

	// Wait for connection
	time.Sleep(time.Second)

	err = client.PauseProgram(1)
	if err != nil {
		t.Errorf("PauseProgram failed: %v", err)
	}
	fmt.Println("PauseProgram executed successfully")
}

func TestPlayProgram(t *testing.T) {
	client, err := NewTCPClient("182.92.117.41:40306")
	if err != nil {
		t.Fatalf("Failed to create TCP client: %v", err)
	}
	defer client.Close()

	// Wait for connection
	time.Sleep(time.Second)

	err = client.PlayProgram(0)
	if err != nil {
		t.Errorf("PlayProgram failed: %v", err)
	}
	fmt.Println("PlayProgram executed successfully")
}

func TestStopProgram(t *testing.T) {
	client, err := NewTCPClient("182.92.117.41:40306")
	if err != nil {
		t.Fatalf("Failed to create TCP client: %v", err)
	}
	defer client.Close()

	// Wait for connection
	time.Sleep(time.Second)

	err = client.StopProgram(1)
	if err != nil {
		t.Errorf("StopProgram failed: %v", err)
	}
	fmt.Println("StopProgram executed successfully")
}
