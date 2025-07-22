# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

GoModo is a simple Pomodoro timer CLI tool written in Go, designed for developers using lightweight terminal environments. It implements the standard Pomodoro technique with 25-minute work sessions followed by 5-minute breaks.

## Common Commands

### Build
```bash
go build -o gomodo
```

### Run
```bash
./gomodo start
# or if installed via go install:
gomodo start
```

### Install
```bash
go install github.com/atoyr/gomodo@latest
```

## Architecture

The application consists of two main files:

- `main.go`: Core timer logic and command-line interface
  - Implements the Pomodoro cycle loop with work/break intervals
  - Handles interrupt signals (Ctrl+C) for graceful shutdown
  - Entry point with simple command parsing

- `notify.go`: Cross-platform notification system
  - Supports macOS (osascript), Linux (notify-send), Windows (PowerShell), and WSL environments
  - Automatically detects WSL by checking `/proc/version` for "microsoft"
  - WSL notifications use external `wsl-notify-send.exe` tool at `/mnt/c/tools/wsl-notify-send.exe`

## Key Constants

- Work duration: 25 minutes
- Break duration: 5 minutes
- App ID for notifications: "gomodo"

## Platform-Specific Notes

- WSL detection is handled automatically via `/proc/version` file parsing
- WSL notifications require `wsl-notify-send.exe` to be available at the hardcoded path
- Notification failures are non-fatal and won't interrupt the timer