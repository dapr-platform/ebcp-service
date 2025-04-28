package client

import (
	"fmt"
	"testing"
	"time"
)

func TestNewPlayerClient(t *testing.T) {
	client, err := NewTCPClient("182.92.117.41:40013")
	if err != nil {
		t.Fatalf("Failed to create TCP client: %v", err)
	}
	defer client.Close()

	// Wait for connection
	time.Sleep(time.Second)
}

func TestGetProgramList(t *testing.T) {
	client, err := NewTCPClient("182.92.117.41:40013")
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

func TestGetAllProgramMedia(t *testing.T) {
	client, err := NewTCPClient("192.168.3.234:17613")
	if err != nil {
		t.Fatalf("Failed to create TCP client: %v", err)
	}
	defer client.Close()

	// Wait for connection
	time.Sleep(time.Second)
	programs, err := client.GetProgramList()
	if err != nil {
		t.Errorf("GetProgramList failed: %v", err)
	}
	if programs == nil {
		t.Error("GetProgramList returned empty response")
	}
	fmt.Printf("GetProgramList response: total=%d, programs=%+v\n", programs.TotalCount, programs.Programs)
	time.Sleep(time.Second)
	// 只查询前3个程序或更少（如果程序总数小于3）
	count := 1
	if len(programs.Programs) < count {
		count = len(programs.Programs)
	}

	for i := 0; i < count; i++ {
		program := programs.Programs[i]
		if program.IsEmpty {
			fmt.Printf("Program %d is empty, skipping\n", program.ID)
			continue
		}
		media, err := client.GetAllProgramMedia(program.ID)
		if err != nil {
			t.Errorf("GetAllProgramMedia failed for program %d: %v", program.ID, err)
		}
		fmt.Printf("GetAllProgramMedia response: programID=%d, media=%+v\n", program.ID, media)
		time.Sleep(time.Second)
	}
	for i := 0; i < 3; i++ {
		progress, err := client.QueryLayerProgress(uint16(i))
		if err != nil {
			t.Errorf("QueryLayerProgress failed for layer %d: %v", i, err)
		}
		time.Sleep(time.Second)
		fmt.Printf("QueryLayerProgress response: layerIndex=%d, progress=%+v\n", i, progress)
	}
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
	client, err := NewTCPClient("192.168.3.234:17613")
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
	client, err := NewTCPClient("192.168.3.234:17613")
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

func TestGetCurrentProgram(t *testing.T) {
	client, err := NewTCPClient("192.168.3.234:17613")
	if err != nil {
		t.Fatalf("Failed to create TCP client: %v", err)
	}
	defer client.Close()

	// Wait for connection
	time.Sleep(time.Second)

	currentProgram, err := client.GetCurrentProgram()
	if err != nil {
		t.Errorf("GetCurrentProgram failed: %v", err)
	}
	fmt.Printf("GetCurrentProgram response: %+v\n", currentProgram)
}
