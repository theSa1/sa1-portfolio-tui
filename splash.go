package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const toWrite = "Sa1 | Portfolio"

func typeWriterAnimation(m Model, text string, totalFrames int) string {
	textLength := len(text)
	framesPerChar := totalFrames / textLength
	toShowTextLength := m.frame / framesPerChar
	if toShowTextLength < textLength {
		return text[:toShowTextLength+1] + blinkingCursor(m.frame, false, m) + strings.Repeat(" ", textLength-toShowTextLength-1)
	}
	return text + blinkingCursor(m.frame, true, m)
}

func getSplashScreen(m Model, width, height int) string {
	if width < 2 || height < 2 {
		return "Invalid dimensions for splash screen"
	}
	splashStyle := m.renderer.NewStyle().
		Width(width).
		Height(height)

	splashContent := typeWriterAnimation(m, toWrite, 30)

	// if (m.frame/10)%2 == 0 {
	// 	splashContent += m.renderer.NewStyle().
	// 		Background(lipgloss.Color("240")).
	// 		Bold(true).
	// 		Render(" ")
	// } else {
	// 	splashContent += m.renderer.NewStyle().
	// 		Bold(true).
	// 		Render(" ")
	// }

	return splashStyle.Render(lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, splashContent))
}
