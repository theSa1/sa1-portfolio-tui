package main

import (
	"net/smtp"
	"os"
	"regexp"

	"github.com/charmbracelet/lipgloss"
)

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

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func isValidEmail(email string) bool {
	// Validate email using regex
	return emailRegex.MatchString(email)
}

func submitForm(email string, name string, message string) bool {
	if !isValidEmail(email) {
		return false
	}

	from := os.Getenv("SMTP_USERNAME")
	to := os.Getenv("SMTP_EMAIL_TO")
	body := "A message from the terminal portfolio\n\nName: " + name + "\nEmail: " + email + "\nMessage:\n---------------------------\n" + message
	auth := smtp.PlainAuth("", from, os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte("Subject: Contact Form Submission\n\n"+body))

	return err == nil
}
