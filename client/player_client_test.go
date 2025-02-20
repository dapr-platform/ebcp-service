package client

import (
	"testing"
	"time"
)

func TestPlayerClient(t *testing.T) {
	// Create TCP client
	client, err := NewTCPClient("182.92.117.41:40306")
	if err != nil {
		t.Fatalf("Failed to create TCP client: %v", err)
	}
	defer client.Close()
	time.Sleep(time.Second)
	// Test PlayCurrent
	success, resourceID, err := client.PlayCurrent()
	if err != nil {
		t.Errorf("PlayCurrent failed: %v", err)
	}
	t.Logf("PlayCurrent result: success=%v, resourceID=%d", success, resourceID)

	// Test PlayNext
	success, resourceID, err = client.PlayNext()
	if err != nil {
		t.Errorf("PlayNext failed: %v", err)
	}
	t.Logf("PlayNext result: success=%v, resourceID=%d", success, resourceID)

	// Test PlayPrev
	success, resourceID, err = client.PlayPrev()
	if err != nil {
		t.Errorf("PlayPrev failed: %v", err)
	}
	t.Logf("PlayPrev result: success=%v, resourceID=%d", success, resourceID)

	// Test PlayByIndex
	success, resourceID, err = client.PlayByIndex(1)
	if err != nil {
		t.Errorf("PlayByIndex failed: %v", err)
	}
	t.Logf("PlayByIndex result: success=%v, resourceID=%d", success, resourceID)

	// Test Pause
	success, resourceID, err = client.Pause()
	if err != nil {
		t.Errorf("Pause failed: %v", err)
	}
	t.Logf("Pause result: success=%v, resourceID=%d", success, resourceID)

	// Test Resume
	success, resourceID, err = client.Resume()
	if err != nil {
		t.Errorf("Resume failed: %v", err)
	}
	t.Logf("Resume result: success=%v, resourceID=%d", success, resourceID)

	// Test SetVolume
	success, err = client.SetVolume(50)
	if err != nil {
		t.Errorf("SetVolume failed: %v", err)
	}
	t.Logf("SetVolume result: success=%v", success)

	// Test SetWindow
	success, err = client.SetWindow(0, 0, 1280, 720, false)
	if err != nil {
		t.Errorf("SetWindow failed: %v", err)
	}
	t.Logf("SetWindow result: success=%v", success)

	// Test SetVisibility
	success, err = client.SetVisibility(true)
	if err != nil {
		t.Errorf("SetVisibility failed: %v", err)
	}
	t.Logf("SetVisibility result: success=%v", success)

	// Test Stop
	success, resourceID, err = client.Stop()
	if err != nil {
		t.Errorf("Stop failed: %v", err)
	}
	t.Logf("Stop result: success=%v, resourceID=%d", success, resourceID)
}
