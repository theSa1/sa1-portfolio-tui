package main

import "github.com/charmbracelet/lipgloss"

func blinkingCursor(frame int, blinking bool, m Model) string {
	style := m.renderer.NewStyle().Bold(true)
	if !blinking || (blinking && (frame/10)%2 == 0) {
		style = style.Background(m.theme.primary)
	}
	return style.Render(" ")
}

func createHyperlink(url, text string) string {
	return "\033]8;;" + url + "\033\\" + text + "\033]8;;\033\\"
}

func wrapAndJoin(lines []string, width int) string {
	var wrappedLines []string
	var lineWidth int = 0
	var currLine string

	for _, line := range lines {
		currWidth := lipgloss.Width(line)
		if lineWidth+currWidth < width {
			lineWidth += currWidth
			currLine = lipgloss.JoinHorizontal(lipgloss.Left, currLine, line)
		} else {
			if currLine != "" {
				wrappedLines = append(wrappedLines, currLine)
			}
			currLine = line
			lineWidth = currWidth
		}
	}

	if currLine != "" {
		wrappedLines = append(wrappedLines, currLine)
	}

	return lipgloss.JoinVertical(lipgloss.Left, wrappedLines...)
}
