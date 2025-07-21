package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

// const SPLASH_SCREEN_FRAME_COUNT = 80

const SPLASH_SCREEN_FRAME_COUNT = 0

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

	// text input
	nameInput       textinput.Model
	emailInput      textinput.Model
	messageInput    textarea.Model
	isSubmitFocused bool
	isSubmitting    bool
	isSubmitted     bool
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func NewModel(renderer *lipgloss.Renderer) Model {
	model := Model{
		renderer:    renderer,
		currentView: splashView,
		theme: Theme{
			foreground:      lipgloss.Color("255"),
			foregroundMuted: lipgloss.Color("244"),
			primary:         lipgloss.Color("#FF5C00"),
			secondary:       lipgloss.Color("45"),
			border:          borders[currentBorder],
		},
		nameInput:    textinput.New(),
		emailInput:   textinput.New(),
		messageInput: textarea.New(),
	}

	model.nameInput.Placeholder = "Your Name"
	model.nameInput.Width = 40
	model.emailInput.Placeholder = "Your Email"
	model.emailInput.Width = 40
	model.messageInput.Placeholder = "Your Message"
	model.messageInput.SetWidth(60)
	model.messageInput.SetHeight(5)
	model.messageInput.ShowLineNumbers = false

	model.messageInput.FocusedStyle.Base = lipgloss.NewStyle().
		Foreground(model.theme.foreground).
		Background(lipgloss.Color("none"))

	model.messageInput.FocusedStyle.CursorLine = lipgloss.NewStyle().
		Background(lipgloss.Color("none")).
		Foreground(model.theme.foreground)

	model.messageInput.BlurredStyle.Base = lipgloss.NewStyle().
		Foreground(model.theme.foreground).
		Background(lipgloss.Color("none"))

	return model
}

func (m Model) Init() tea.Cmd {
	return tick()
}

func (m *Model) OnContactActive() {
	m.nameInput.SetValue("")
	m.nameInput.Focus()
	m.emailInput.SetValue("")
	m.emailInput.Blur()
	m.messageInput.SetValue("")
	m.messageInput.Blur()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.nameInput, cmd = m.nameInput.Update(msg)
	cmds = append(cmds, cmd)

	m.emailInput, cmd = m.emailInput.Update(msg)
	cmds = append(cmds, cmd)

	m.messageInput, cmd = m.messageInput.Update(msg)
	cmds = append(cmds, cmd)

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
				if m.currentView == contactView {
					m.OnContactActive()
				}
				m.scrollOffset = 0
			}
		case "left":
			if m.currentView > homeView {
				m.currentView--
				if m.currentView == contactView {
					m.OnContactActive()
				}
				m.scrollOffset = 0
			}
		case "tab":
			if m.currentView == contactView {
				if m.nameInput.Focused() {
					m.nameInput.Blur()
					m.emailInput.Focus()
				} else if m.emailInput.Focused() {
					m.emailInput.Blur()
					m.messageInput.Focus()
				} else if m.messageInput.Focused() {
					m.messageInput.Blur()
					m.isSubmitFocused = true
				} else if m.isSubmitFocused {
					m.isSubmitFocused = false
					m.nameInput.Focus()
				}
			}
		case "enter":
			if m.currentView == contactView {
				if (!m.isSubmitting || m.isSubmitted) && m.isSubmitFocused {
					m.isSubmitting = true
					m.isSubmitted = false
					m.nameInput.SetValue("")
					m.emailInput.SetValue("")
					m.messageInput.SetValue("")
				}
			}
		case "up":
			if (m.currentView == aboutView || m.currentView == projectsView) && m.scrollOffset > 0 {
				m.scrollOffset--
			}
		case "down":
			if (m.currentView == aboutView || m.currentView == projectsView) && m.scrollOffset < scrollViewTotalLines-scrollViewUsableHeight {
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
	return m, tea.Batch(cmds...)
}
