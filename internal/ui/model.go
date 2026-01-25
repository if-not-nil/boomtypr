package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/yagnikpt/boomtypr/internal/typing"
	"github.com/yagnikpt/boomtypr/internal/utils"
)

var (
	frameStyles        = lipgloss.NewStyle().Padding(2, 10)
	pendingCharStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#4C4C4C"))
	incorrectCharStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Underline(true)
	cursorStyle        = lipgloss.NewStyle().Background(lipgloss.Color("7")).Foreground(lipgloss.Color("0"))
)

type State int

const (
	StateMenu State = iota
	StateTyping
	StateResults
)

var LinesWindowSize = 3

type Line struct {
	Text  []rune
	Start int
}

type Model struct {
	State  State
	Engine *typing.Engine
	Lines  []Line
	Width  int
	Height int
}

func NewModel(words []string) Model {
	joinedWords := strings.Join(words, " ")
	termWidth, _, _ := GetTermDimensions()
	frameX, _ := frameStyles.GetFrameSize()
	wrappedPara := text.WrapSoft(joinedWords, termWidth-frameX)
	// fmt.Println(wrappedPara)
	lineBreaks := utils.LineBreakIndexes(wrappedPara)
	lines := make([]Line, utils.CountLines(wrappedPara))
	linesFromPara := utils.SplitIntoLines(wrappedPara)
	for i, line := range linesFromPara {
		startIdx := 0
		if i > 0 {
			startIdx = lineBreaks[i-1] + 1
		}

		lines[i] = Line{
			Text:  []rune(line),
			Start: startIdx,
		}
	}

	return Model{
		State:  StateMenu,
		Lines:  lines,
		Engine: typing.NewEngine(joinedWords, utils.LineBreakIndexes(wrappedPara)),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "backspace":
			m.Engine.Backspace()
		default:
			if len(msg.Runes) > 0 {
				m.Engine.TypeChar(msg.Runes[0])
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	var b strings.Builder

	padding := GetPaddingToCenterVertically(m.Height, LinesWindowSize, 2)
	b.WriteString(padding)
	for i := m.Engine.CurrentLine; i < m.Engine.CurrentLine+LinesWindowSize && i < len(m.Lines); i++ {
		line := m.Lines[i]
		for j, char := range line.Text {
			charIndex := line.Start + j

			rendered := string(char)
			switch m.Engine.Track[charIndex] {
			case typing.CharPending:
				rendered = pendingCharStyle.Render(rendered)
			case typing.CharIncorrect:
				rendered = incorrectCharStyle.Render(rendered)
			}

			if charIndex == m.Engine.CurrentChar {
				rendered = cursorStyle.Render(rendered)
			}
			b.WriteString(rendered)
		}
		b.WriteString("\n")
	}

	return frameStyles.Render(b.String())
}
