package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const toWrite = "Sa1 | Portfolio"

func typeWriterAnimation(frame int, text string, totalFrames int) string {
	textLength := len(text)
	framesPerChar := totalFrames / textLength
	toShowTextLength := frame / framesPerChar
	if toShowTextLength < textLength {
		return text[:toShowTextLength+1] + blinkingCursor(frame, false) + strings.Repeat(" ", textLength-toShowTextLength-1)
	}
	return text + blinkingCursor(frame, true)
}

func getSplashScreen(m Model, width, height int) string {
	if width < 2 || height < 2 {
		return "Invalid dimensions for splash screen"
	}
	splashStyle := lipgloss.NewStyle().
		Width(width).
		Height(height)

	splashContent := typeWriterAnimation(m.frame, toWrite, 30)

	// if (m.frame/10)%2 == 0 {
	// 	splashContent += lipgloss.NewStyle().
	// 		Background(lipgloss.Color("240")).
	// 		Bold(true).
	// 		Render(" ")
	// } else {
	// 	splashContent += lipgloss.NewStyle().
	// 		Bold(true).
	// 		Render(" ")
	// }

	return splashStyle.Render(lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, splashContent))
}
