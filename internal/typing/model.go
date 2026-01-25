package typing

import "time"

type Model struct {
	// test config
	Mode     Mode
	Target   []rune        // full text to type
	Duration time.Duration // for time mode

	// runtime state
	StartedAt time.Time
	EndedAt   time.Time
	Index     int // cursor position in Target
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
