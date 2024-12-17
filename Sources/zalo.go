package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// Chờ 2 giây
	time.Sleep(2 * time.Second)

	// AppleScript để kiểm tra và xóa login item "Zalo"
	appleScript := `tell application "System Events"
		if exists login item "Zalo" then
			delete login item "Zalo"
		end if
	end tell`

	// Chạy lệnh osascript
	cmd := exec.Command("osascript", "-e", appleScript)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing AppleScript: %v\n", err)
		return
	}
}
