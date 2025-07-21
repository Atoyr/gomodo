package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

const (
	workDuration  = 25 * time.Minute
	breakDuration = 5 * time.Minute
)

func startPomodoro() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		fmt.Printf("start %d min", workDuration/time.Minute)
		notify("gomode", "start 25 min")

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
		fmt.Println("使い方: pomodoro start")
		return
	}

	switch os.Args[1] {
	case "start":
		startPomodoro()
	default:
		fmt.Println("未知のコマンド:", os.Args[1])
	}
}
