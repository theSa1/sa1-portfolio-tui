package main

import "github.com/charmbracelet/lipgloss"

func getContactView(m Model, width, height int) string {
	boldStyle := m.renderer.NewStyle().
		Foreground(m.theme.primary).
		Bold(true)

	// Contact form
	content := boldStyle.Render("# Contact Me") + "\n\n"

	// Add all three input fields with proper spacing
	sectionWidth := (width - 2) / 2
	nameField := lipgloss.NewStyle().Width(sectionWidth).Height(2).Render("Name:\n" + m.nameInput.View())
	emailField := lipgloss.NewStyle().Width(sectionWidth).Height(2).Render("Email:\n" + m.emailInput.View())
	content += lipgloss.JoinHorizontal(lipgloss.Left, nameField, "  ", emailField) + "\n\n"
	content += "Message:\n" + m.messageInput.View() + "\n\n"

	if m.isSubmitFocused {
		content += m.renderer.NewStyle().
			Foreground(m.theme.background).
			Background(m.theme.primary).
			Bold(true).
			Render(" > Submit ")
	} else {
		content += m.renderer.NewStyle().
			Foreground(m.theme.primary).
			Bold(true).
			Render("   Submit ")
	}

	content += " " + m.contactMessage

	content += "\n\nUse Tab to navigate."

	content += "\n\n" + createHyperlink("https://github.com/thesa1", "GitHub") + " | " +
		createHyperlink("https://x.com/sa1wastooshort", "X") + " | " +
		createHyperlink("http://linkedin.com/in/savan-bhanderi", "LinkedIn") + " | " +
		createHyperlink("https://sa1.dev", "sa1.dev")

	// Render the content with the theme and dimensions
	container := m.renderer.NewStyle().
		Width(width).
		Height(height).
		Padding(1, 3).
		Render(content)

	return m.renderer.NewStyle().
		Width(width).
		Height(height).
		Render(container)
}
