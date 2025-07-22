package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"
)

const (
	defaultWorkDuration = 25
	breakDuration       = 5 * time.Minute
)

func startPomodoro(workMinutes int) {
	workDuration := time.Duration(workMinutes) * time.Minute
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		fmt.Printf("作業開始（%d分）\n", workMinutes)
		notify("GoModo", fmt.Sprintf("作業開始！%d分集中しよう！", workMinutes))

		select {
		case <-time.After(workDuration):
			// 作業終了後に通知して休憩へ
		case <-interrupt:
			fmt.Println("\nポモドーロを中断しました。")
			return
		}

		fmt.Println("休憩開始（5分）")
		notify("GoModo", "休憩開始！5分休もう！")

		select {
		case <-time.After(breakDuration):
			// 次の作業へループ
		case <-interrupt:
			fmt.Println("\nポモドーロを中断しました。")
			return
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使い方: gomodo start [分数]")
		fmt.Println("例: gomodo start 20")
		return
	}

	switch os.Args[1] {
	case "start":
		workMinutes := defaultWorkDuration
		if len(os.Args) >= 3 {
			if minutes, err := strconv.Atoi(os.Args[2]); err == nil && minutes > 0 {
				workMinutes = minutes
			} else {
				fmt.Printf("無効な分数です。デフォルトの%d分を使用します。\n", defaultWorkDuration)
			}
		}
		startPomodoro(workMinutes)
	default:
		fmt.Println("未知のコマンド:", os.Args[1])
		fmt.Println("使い方: gomodo start [分数]")
	}
}
