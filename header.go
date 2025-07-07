package main

import (
	"github.com/charmbracelet/lipgloss"
)

func getHeader(m Model, width int) string {
	logo := m.renderer.NewStyle().
		Padding(0, 3, 0, 1).
		Render("Sa1 | Portfolio" + blinkingCursor(m.frame, true, m))

	navBarStyle := m.renderer.NewStyle().
		Width(width - lipgloss.Width(logo)).
		Align(lipgloss.Right)

	// Nav bar items
	navItems := []string{"Home", "About", "Projects", "Contact"}
	for i, item := range navItems {
		switch m.currentView {
		case uint(i) + 1:
			navItems[i] = m.renderer.NewStyle().
				Foreground(m.theme.foreground).
				Background(m.theme.primary).
				Bold(true).
				Padding(0, 2).
				MarginRight(1).
				Render(item)
		default:
			navItems[i] = m.renderer.NewStyle().
				Foreground(m.theme.foreground).
				Padding(0, 2).
				MarginRight(1).
				Render(item)
		}
	}
	navBar := navBarStyle.Render(
		lipgloss.JoinHorizontal(lipgloss.Center, navItems...),
	)

	header := lipgloss.JoinHorizontal(lipgloss.Left, logo, navBar)

	return header
}
