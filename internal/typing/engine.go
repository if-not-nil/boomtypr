package typing

import (
	"time"
)

type CharState int

const (
	CharPending CharState = iota
	CharCorrect
	CharIncorrect
)

type Engine struct {
	Text        string
	CurrentChar int
	CurrentLine int
	Cursor      int
	LineBreaks  []int
	Track       []CharState
	StartTime   time.Time
	Started     bool
	Finished    bool
}

func NewEngine(text string, lineBreaks []int) *Engine {
	track := make([]CharState, len(text))
	return &Engine{
		Text:        text,
		LineBreaks:  lineBreaks,
		Track:       track,
		Cursor:      0,
		CurrentChar: 0,
		CurrentLine: 0,
	}
}

func (e *Engine) Start() {
	// TODO: implement
}

func (e *Engine) TypeChar(char rune) {
	if e.Finished {
		return
	}
	if e.LineBreaks[e.CurrentLine] == e.CurrentChar && string(char) != " " {
		return
	}
	if e.CurrentChar < len(e.Text) && e.CurrentChar == e.LineBreaks[e.CurrentLine] {
		e.CurrentLine++
		e.Cursor = 0
	} else {
		e.Cursor++
	}
	if string(char) == string(e.Text[e.CurrentChar]) {
		e.Track[e.CurrentChar] = CharCorrect
	} else {
		e.Track[e.CurrentChar] = CharIncorrect
	}
	e.CurrentChar++
	if e.CurrentChar >= len(e.Text) {
		e.Finished = true
	}
}

func (e *Engine) Backspace() {
	if e.Finished || e.CurrentChar == 0 {
		return
	}
	if e.CurrentLine > 0 && e.CurrentChar == e.LineBreaks[e.CurrentLine-1]+1 {
		e.CurrentLine--
		e.Cursor = 0
	}
	e.CurrentChar--
	e.Cursor--
	e.Track[e.CurrentChar] = CharPending
}

func (e *Engine) NextWord() {
	// TODO: implement
}

func (e *Engine) Reset() {
	// TODO: implement
}
