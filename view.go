package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	width := m.width
	height := m.height

	if width > 80 {
		width = 80
	}
	if height > 24 {
		height = 24
	}

	if m.currentView == splashView {
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, getSplashScreen(m, width, height))
	}

	header := getHeader(m, width)

	footer := lipgloss.NewStyle().
		Width(width - 2).
		Height(1).
		Align(lipgloss.Center).
		Render("Press 'q' to quit | Use arrow keys to navigate")

	width -= 2
	height = height - lipgloss.Height(header) - lipgloss.Height(footer) - 2

	contentStyle := lipgloss.NewStyle().
		Width(width).
		Height(height).
		Align(lipgloss.Center).
		BorderStyle(m.theme.border)

	var content string
	switch m.currentView {
	case homeView:
		content = contentStyle.Render(getHomeView(m, width, height))
	case aboutView:
		content = contentStyle.Render(getAboutView(m, width, height))
	case projectsView:
		content = contentStyle.Render(getProjectsView(m, width, height))
	case contactView:
		content = contentStyle.Render("Contact view is under construction.")
	default:
		content = contentStyle.Render(fmt.Sprintf("Current View: %d", m.currentView))
	}
	body := lipgloss.JoinVertical(lipgloss.Left, header, content, footer)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, body)
}
