package main

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
