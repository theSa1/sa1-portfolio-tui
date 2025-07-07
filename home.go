package main

import (
	"github.com/charmbracelet/lipgloss"
)

func getHomeView(m Model, width, height int) string {
	homeStyle := m.renderer.NewStyle().
		Width(width).
		Height(height)

	heroText := `███████╗  █████╗   ██╗
██╔════╝ ██╔══██╗ ███║
███████╗ ███████║ ╚██║
╚════██║ ██╔══██║  ██║
███████║ ██║  ██║  ██║
╚══════╝ ╚═╝  ╚═╝  ╚═╝

Inspired by ` + createHyperlink("https://terminal.shop", "terminal.shop") + `
Go to ` + createHyperlink("https://sa1.dev", "sa1.dev") + ` for web version`

	homeContent := m.renderer.NewStyle().
		Width(width - 6).
		Height(lipgloss.Height(heroText)).
		Align(lipgloss.Center).
		Render(heroText)

	return homeStyle.Render(lipgloss.Place(
		width, height,
		lipgloss.Center, lipgloss.Center,
		homeContent,
	))
}
