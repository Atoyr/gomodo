package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const appId string = "gomodo"

func notify(title, message string) {
	var err error = nil
	switch runtime.GOOS {
	case "darwin": // macOS
		err = exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "%s"`, message, title)).Run()

	case "linux":
		// WSLでなければLinux（X11）とみなす
		if isWSL() {
			err = notifyWsl(title, message)
		} else {
			err = notifyLinux(title, message)
		}

	case "windows": // native WindowsでのGoビルド（通常対象外）
		err = exec.Command("powershell.exe", "-Command", fmt.Sprintf(`New-BurntToastNotification -Text "%s","%s"`, title, message)).Run()

	default:
		err = fmt.Errorf("通知未対応のOSです:%s", runtime.GOOS)
	}

	// 通知がエラーになっても処理自体は停止しない
	if err != nil {
		fmt.Println("通知に失敗:", err)
	}
}

// isWSL
func isWSL() bool {
	// /proc/version に "microsoft" が含まれていればWSL
	if data, err := os.ReadFile("/proc/version"); err == nil {
		return strings.Contains(strings.ToLower(string(data)), "microsoft")
	}
	return false
}

func notifyLinux(title, message string) error {
	return exec.Command("notify-send", title, message).Run()
}

func notifyWsl(title, message string) error {
	// FIXME: パスを動的に設定するかパスを通してあることを前提としたい
	exePath := "/mnt/c/tools/wsl-notify-send.exe"
	args := []string{
		"--appId", appId,
		"-c", title,
		message,
	}

	return exec.Command(exePath, args...).Run()
}
