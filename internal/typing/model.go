package typing

import "time"

type Model struct {
	// test config
	Mode     Mode
	Target   []rune
	Duration time.Duration

	// runtime state
	StartedAt time.Time
	EndedAt   time.Time
	Done      bool

	// input tracking
	Keystrokes []Keystroke

	// UI
	FocusMode bool
	Err       error
}

type Keystroke struct {
	Rune      rune
	Expected  rune
	Time      time.Time
	Correct   bool
	Backspace bool
}

type Mode int

const (
	ModeTime Mode = iota
	ModeWords
	ModeZen
)
