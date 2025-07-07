package main

import (
	"github.com/charmbracelet/lipgloss"
)

type Project struct {
	name        string
	description string
	languages   []string
	tools       []string
}

func getProjectsView(m Model, width, height int) string {
	projectsStyle := m.renderer.NewStyle().
		Width(width).
		Height(height)

	projectsContent := "Projects view is under construction."

	return projectsStyle.Render(lipgloss.Place(
		width, height,
		lipgloss.Center, lipgloss.Center,
		m.renderer.NewStyle().Render(projectsContent),
	))
}
