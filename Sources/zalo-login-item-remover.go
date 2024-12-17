package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	// Khai báo tham số
	setupFlag := flag.Bool("setup", false, "Run setup tasks")
	uninstallFlag := flag.Bool("uninstall", false, "Run uninstall tasks")

	// Phân tích các tham số dòng lệnh
	flag.Parse()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	// Đường dẫn đầy đủ tới file plist
	plistPath := filepath.Join(homeDir, "/Library/LaunchAgents/com.zalo.login-item-remover.plist")

	// Dọn dẹp
	if *uninstallFlag {
		cmd := exec.Command("rm", "-rf", plistPath)
		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error loading launchctl: %v\n", err)
			return
		}
		return
	}

	// Tạo file plist
	if *setupFlag {
		plistContent := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>com.zalo.login-item-remover</string>

		<key>ProgramArguments</key>
		<array>
			<string>/opt/homebrew/bin/zalo-login-item-remover</string>
		</array>

		<key>WatchPaths</key> 
		<array>
			<string>%s/Library/Application Support/ZaloData/startup.log</string>
		</array>

		<key>RunAtLoad</key>
		<true/>
	</dict>
</plist>`, homeDir)

		// Mở hoặc tạo file plist trong thư mục home
		file, err := os.Create(plistPath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		// Ghi nội dung vào file
		_, err = file.WriteString(plistContent)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("File created successfully at:", plistPath)

		cmd := exec.Command("launchctl", "load", plistPath)
		_, err2 := cmd.CombinedOutput()
		if err2 != nil {
			fmt.Printf("Error loading launchctl: %v\n", err2)
			return
		}
		return
	} else {
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
}
