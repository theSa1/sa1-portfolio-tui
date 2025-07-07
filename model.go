package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const SPLASH_SCREEN_FRAME_COUNT = 80

var borders = []lipgloss.Border{
	lipgloss.RoundedBorder(),
	lipgloss.DoubleBorder(),
	lipgloss.NormalBorder(),
	lipgloss.ASCIIBorder(),
	lipgloss.ThickBorder(),
	lipgloss.MarkdownBorder(),
	lipgloss.HiddenBorder(),
}
var currentBorder = 0

var scrollViewTotalLines = 0
var scrollViewUsableHeight = 0

const (
	splashView uint = iota
	homeView
	aboutView
	projectsView
	contactView
)

type Theme struct {
	background      lipgloss.Color
	foreground      lipgloss.Color
	foregroundMuted lipgloss.Color
	primary         lipgloss.Color
	secondary       lipgloss.Color

	border lipgloss.Border
}

type Model struct {
	renderer *lipgloss.Renderer

	scrollOffset int

	currentView uint
	height      int
	width       int

	// for animations
	frame int

	// theme
	theme Theme
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func NewModel(renderer *lipgloss.Renderer) Model {
	return Model{
		renderer:    renderer,
		currentView: splashView,
		theme: Theme{
			foreground:      lipgloss.Color("255"),
			foregroundMuted: lipgloss.Color("244"),
			primary:         lipgloss.Color("#FF5C00"),
			secondary:       lipgloss.Color("45"),
			border:          borders[currentBorder],
		},
	}
}

func (m Model) Init() tea.Cmd {
	return tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		m.frame++
		if m.frame > 100 {
			m.frame = 0
		}
		if m.currentView == splashView && m.frame > SPLASH_SCREEN_FRAME_COUNT {
			m.currentView = homeView
		}
		return m, tick()
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "right":
			if m.currentView < contactView {
				m.currentView++
				m.scrollOffset = 0
			}
		case "left":
			if m.currentView > homeView {
				m.currentView--
				m.scrollOffset = 0
			}
		case "up":
			if m.currentView == aboutView && m.scrollOffset > 0 {
				m.scrollOffset--
			}
		case "down":
			if m.currentView == aboutView && m.scrollOffset < scrollViewTotalLines-scrollViewUsableHeight {
				m.scrollOffset++
			}
		case "b":
			currentBorder = (currentBorder + 1) % len(borders)
			m.theme.border = borders[currentBorder]
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}
	return m, nil
}
