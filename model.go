package main

import (
	"errors"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/huh"
)

const SPLASH_SCREEN_FRAME_COUNT = 80

// const SPLASH_SCREEN_FRAME_COUNT = 0

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
	foreground      lipgloss.Color
	foregroundMuted lipgloss.Color
	primary         lipgloss.Color
	secondary       lipgloss.Color
}

type Model struct {
	renderer *lipgloss.Renderer

	scrollOffset int

	currentView uint
	height      int
	width       int

	frame int

	theme Theme

	contactForm     *huh.Form
	isFormSubmitted bool
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(50*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Stolen from terminal.shop
func copyTextStyles(t huh.TextInputStyles) huh.TextInputStyles {
	return huh.TextInputStyles{
		Cursor:      t.Cursor.Copy(),
		Placeholder: t.Placeholder.Copy(),
		Prompt:      t.Prompt.Copy(),
		Text:        t.Text.Copy(),
	}
}

func copyFieldStyles(f huh.FieldStyles) huh.FieldStyles {
	return huh.FieldStyles{
		Base:           f.Base.Copy(),
		Title:          f.Title.Copy(),
		Description:    f.Description.Copy(),
		ErrorIndicator: f.ErrorIndicator.Copy(),
		ErrorMessage:   f.ErrorMessage.Copy(),
		SelectSelector: f.SelectSelector.Copy(),
		// NextIndicator:       f.NextIndicator.Copy(),
		// PrevIndicator:       f.PrevIndicator.Copy(),
		Option: f.Option.Copy(),
		// Directory:           f.Directory.Copy(),
		// File:                f.File.Copy(),
		MultiSelectSelector: f.MultiSelectSelector.Copy(),
		SelectedOption:      f.SelectedOption.Copy(),
		SelectedPrefix:      f.SelectedPrefix.Copy(),
		UnselectedOption:    f.UnselectedOption.Copy(),
		UnselectedPrefix:    f.UnselectedPrefix.Copy(),
		FocusedButton:       f.FocusedButton.Copy(),
		BlurredButton:       f.BlurredButton.Copy(),
		TextInput:           copyTextStyles(f.TextInput),
		Card:                f.Card.Copy(),
		NoteTitle:           f.NoteTitle.Copy(),
		Next:                f.Next.Copy(),
	}
}

func getHuhTheme(m Model) *huh.Theme {
	var t huh.Theme

	t.FieldSeparator = m.renderer.NewStyle().SetString("\n\n")

	t.FieldSeparator = m.renderer.NewStyle().SetString("\n\n")

	f := &t.Focused
	f.Base = m.renderer.NewStyle().
		PaddingLeft(1).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeft(true).
		BorderForeground(m.theme.primary)
	f.Title = m.renderer.NewStyle().Foreground(m.theme.primary)
	f.TextInput.Cursor = m.renderer.NewStyle().Foreground(m.theme.primary)
	f.TextInput.Placeholder = m.renderer.NewStyle().Foreground(m.theme.foregroundMuted)
	f.TextInput.Prompt = m.renderer.NewStyle().Foreground(m.theme.foreground)
	f.TextInput.Text = m.renderer.NewStyle().Foreground(m.theme.foreground)
	f.ErrorIndicator = m.renderer.NewStyle().Foreground(lipgloss.Color("FF0000"))
	f.ErrorMessage = m.renderer.NewStyle().Foreground(lipgloss.Color("FF0000"))
	// t.Help = help.New().Styles

	t.Blurred = copyFieldStyles(*f)
	t.Blurred.Base = t.Blurred.Base.BorderForeground(m.theme.foregroundMuted)
	t.Blurred.Title.Foreground(m.theme.foregroundMuted)

	return &t
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
		},
		isFormSubmitted: false,
	}

	model.contactForm = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("name").
				Title("Name").
				Validate(func(s string) error {
					if s == "" {
						return errors.New("Name cannot be empty")
					}
					if len(s) < 3 {
						return errors.New("Name must be at least 3 characters long")
					}
					return nil
				}),
			huh.NewInput().
				Key("email").
				Title("Email").
				Placeholder("Your Email").
				Validate(func(s string) error {
					if s == "" {
						return errors.New("Email cannot be empty")
					}
					if !isValidEmail(s) {
						return errors.New("Invalid email format")
					}
					return nil
				}),
			huh.NewText().
				Key("message").
				Title("Message").
				Placeholder("Your Message").
				Lines(4).
				Validate(func(s string) error {
					if s == "" {
						return errors.New("Message cannot be empty")
					}
					if len(s) < 3 {
						return errors.New("Message must be at least 10 characters long")
					}
					return nil
				}),
		),
	).WithTheme(getHuhTheme(model)).WithShowHelp(false)

	return model
}

func (m Model) Init() tea.Cmd {
	return tick()
}

func (m *Model) OnContactActive() {

}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

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
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			if m.currentView != contactView {
				return m, tea.Quit
			}
		case "right":
			if m.currentView < contactView {
				m.currentView++
				if m.currentView == contactView {
					// m.OnContactActive()
					cmds = append(cmds, m.contactForm.Init())
				}
				m.scrollOffset = 0
			}
		case "left":
			if m.currentView > homeView {
				m.currentView--
				if m.currentView == contactView {
					// m.OnContactActive()
					cmds = append(cmds, m.contactForm.Init())
				}
				m.scrollOffset = 0
			}
		case "up":
			if (m.currentView == aboutView || m.currentView == projectsView) && m.scrollOffset > 0 {
				m.scrollOffset--
			}
		case "down":
			if (m.currentView == aboutView || m.currentView == projectsView) && m.scrollOffset < scrollViewTotalLines-scrollViewUsableHeight {
				m.scrollOffset++
			}
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}

	form, cmd := m.contactForm.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.contactForm = f
	}
	cmds = append(cmds, cmd)

	if m.contactForm.State == huh.StateCompleted && !m.isFormSubmitted {
		name := m.contactForm.Get("name").(string)
		email := m.contactForm.Get("email").(string)
		message := m.contactForm.Get("message").(string)

		go submitForm(email, name, message)
		m.isFormSubmitted = true
	}

	return m, tea.Batch(cmds...)
}
