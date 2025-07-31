package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Project struct {
	name        string
	description string
	tools       []string
}

var projects = []Project{{
	name:        "University Timetable Manager",
	description: "A web application to manage university timetables, that allows faculties to create and manage acadaemic timetables, and help them use rooms and resources efficiently.",
	tools:       []string{"Next.js", "Tailwind CSS", "Docker"},
}, {
	name:        "AdhyAI: AI based study companion",
	description: "An AI based study companion that helps students with their studies by providing answers to their questions based on their syllabus and textbooks, and also help prepare for exams by providing quizzes and tests.",
	tools:       []string{"Langchain", "Google Gemini", "Next.js", "Tailwind CSS", "Docker"},
}}

func getProjectsView(m Model, width, height int) string {
	boldStyle := m.renderer.NewStyle().
		Foreground(m.theme.primary).
		Bold(true)

	boldNonePrimary := m.renderer.NewStyle().
		Foreground(m.theme.foreground).
		Bold(true)

	projectsStyle := m.renderer.NewStyle().
		Width(width).
		Height(height)

	usableHeight := height - 2
	usableWidth := width - 6

	projectsText := boldStyle.Render("# Projects:") + "\n\n"

	var toolStyle = m.renderer.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).MarginRight(1)
	for i, p := range projects {
		var toolsText []string
		if len(p.tools) > 0 {
			for _, tool := range p.tools {
				toolsText = append(toolsText, toolStyle.Render(tool))
			}
		}

		projectText := m.renderer.NewStyle().
			Width(usableWidth - 3).
			MarginLeft(3).
			Render("" + p.description + "\n" + wrapAndJoin(toolsText, usableWidth-3))
		projectContainer := m.renderer.NewStyle().
			Width(usableWidth).
			Render(boldNonePrimary.Render("## "+p.name) + "\n\n" + projectText)

		projectsText += projectContainer
		if i < len(projects)-1 {
			projectsText += "\n\n\n"
		}
	}

	projectsContent := m.renderer.NewStyle().
		Width(usableWidth).
		Height(usableHeight).
		Render(projectsText)

	totalLines := lipgloss.Height(projectsContent)
	scrollViewTotalLines = totalLines
	scrollViewUsableHeight = usableHeight
	if totalLines > usableHeight {
		projectsContentLines := strings.Split(projectsContent, "\n")
		startLine := m.scrollOffset
		endLine := min(int(startLine)+usableHeight, totalLines)
		projectsContent = strings.Join(projectsContentLines[startLine:endLine], "\n")
	}

	topStyle := m.renderer.NewStyle().
		Width(usableWidth).
		Height(1).
		Align(lipgloss.Center)

	bottomStyle := m.renderer.NewStyle().
		Width(usableWidth).
		Height(1).
		Align(lipgloss.Center)

	top := ""
	bottom := ""

	if m.scrollOffset != 0 {
		top = topStyle.Render("⌃")
	}
	if m.scrollOffset+usableHeight < totalLines {
		bottom = bottomStyle.Render("⌄")
	}

	return projectsStyle.Render(
		lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, top+"\n"+projectsContent+"\n"+bottom),
	)
}
