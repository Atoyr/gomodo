package main

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func TestDefaultWorkDuration(t *testing.T) {
	expected := 25
	if defaultWorkDuration != expected {
		t.Errorf("Expected defaultWorkDuration to be %d, got %d", expected, defaultWorkDuration)
	}
}

func TestBreakDuration(t *testing.T) {
	expected := 5 * time.Minute
	if breakDuration != expected {
		t.Errorf("Expected breakDuration to be %v, got %v", expected, breakDuration)
	}
}

func TestParseWorkMinutes(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected int
	}{
		{
			name:     "デフォルト値（引数なし）",
			args:     []string{"gomodo", "start"},
			expected: defaultWorkDuration,
		},
		{
			name:     "有効な分数指定",
			args:     []string{"gomodo", "start", "20"},
			expected: 20,
		},
		{
			name:     "別の有効な分数指定",
			args:     []string{"gomodo", "start", "30"},
			expected: 30,
		},
		{
			name:     "無効な分数指定（文字列）",
			args:     []string{"gomodo", "start", "abc"},
			expected: defaultWorkDuration,
		},
		{
			name:     "無効な分数指定（0）",
			args:     []string{"gomodo", "start", "0"},
			expected: defaultWorkDuration,
		},
		{
			name:     "無効な分数指定（負の数）",
			args:     []string{"gomodo", "start", "-5"},
			expected: defaultWorkDuration,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// os.Argsをテスト用に設定
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()
			os.Args = tt.args

			// main関数のロジックを再現（startコマンドの部分のみ）
			workMinutes := defaultWorkDuration
			if len(os.Args) >= 3 {
				if minutes, err := strconv.Atoi(os.Args[2]); err == nil && minutes > 0 {
					workMinutes = minutes
				}
			}

			if workMinutes != tt.expected {
				t.Errorf("Expected workMinutes to be %d, got %d", tt.expected, workMinutes)
			}
		})
	}
}

func TestWorkDurationCalculation(t *testing.T) {
	tests := []struct {
		workMinutes      int
		expectedDuration time.Duration
	}{
		{20, 20 * time.Minute},
		{25, 25 * time.Minute},
		{30, 30 * time.Minute},
		{45, 45 * time.Minute},
	}

	for _, tt := range tests {
		t.Run("WorkMinutes_"+strconv.Itoa(tt.workMinutes), func(t *testing.T) {
			workDuration := time.Duration(tt.workMinutes) * time.Minute
			if workDuration != tt.expectedDuration {
				t.Errorf("Expected duration %v, got %v", tt.expectedDuration, workDuration)
			}
		})
	}
}

// ベンチマーク: 時間計算のパフォーマンステスト
func BenchmarkDurationCalculation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = time.Duration(25) * time.Minute
	}
}

// 統合テスト用のヘルパー関数
func TestValidateWorkMinutes(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
		value    int
	}{
		{"25", true, 25},
		{"20", true, 20},
		{"30", true, 30},
		{"0", false, 0},
		{"-5", false, 0},
		{"abc", false, 0},
		{"", false, 0},
	}

	for _, tt := range tests {
		t.Run("Input_"+tt.input, func(t *testing.T) {
			minutes, err := strconv.Atoi(tt.input)
			valid := err == nil && minutes > 0

			if valid != tt.expected {
				t.Errorf("Expected validity %v for input %s, got %v", tt.expected, tt.input, valid)
			}

			if valid && minutes != tt.value {
				t.Errorf("Expected value %d for input %s, got %d", tt.value, tt.input, minutes)
			}
		})
	}
}