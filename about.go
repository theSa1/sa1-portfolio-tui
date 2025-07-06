package main

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type skill struct {
	name    string
	percent int
}

var skills1 = []skill{
	{"HTML", 90},
	{"CSS", 80},
	{"JavaScript", 70},
	{"Python", 60},
	{"Go", 50},
}

func getAboutView(m Model, width, height int) string {
	aboutStyle := lipgloss.NewStyle().
		Width(width).
		Height(height)

	boldStyle := lipgloss.NewStyle().
		Bold(true)

	usableHeight := height - 2
	usableWidth := width - 6

	skillsSectionWidth := (usableWidth - 3) / 2

	skillsSection1 := boldStyle.Render("## Languages & Frameworks") + "\n\n"
	for i, skill := range skills1 {
		emptyProgressChar := lipgloss.NewStyle().Background(lipgloss.Color("240")).Render(" ")
		filledProgressChar := lipgloss.NewStyle().Background(lipgloss.Color("205")).Render(" ")
		progressBar := strings.Repeat(filledProgressChar, skill.percent/10) + strings.Repeat(emptyProgressChar, 10-skill.percent/10)
		skillsSection1 += lipgloss.NewStyle().
			Render("- " + skill.name + " " + progressBar + " " + strconv.Itoa(skill.percent) + "%")
		if i < len(skills1)-1 {
			skillsSection1 += "\n\n"
		}
	}

	skillsSection2 := boldStyle.Render("## Tools & Platforms") + "\n\n"
	for i, skill := range skills1 {
		emptyProgressChar := lipgloss.NewStyle().Background(lipgloss.Color("240")).Render(" ")
		filledProgressChar := lipgloss.NewStyle().Background(lipgloss.Color("205")).Render(" ")
		progressBar := strings.Repeat(filledProgressChar, skill.percent/10) + strings.Repeat(emptyProgressChar, 10-skill.percent/10)
		skillsSection2 += lipgloss.NewStyle().
			Render("- " + skill.name + " " + progressBar + " " + strconv.Itoa(skill.percent) + "%")
		if i < len(skills1)-1 {
			skillsSection2 += "\n\n"
		}
	}

	skillsSectionHeight := max(lipgloss.Height(skillsSection1), lipgloss.Height(skillsSection2))
	separator := lipgloss.NewStyle().
		Width(usableWidth - (skillsSectionWidth * 2)).
		Align(lipgloss.Center).
		Render(strings.Repeat("│\n", skillsSectionHeight))

	skillsSectionsStyle := lipgloss.NewStyle().
		Width(skillsSectionWidth)

	skillsSection := lipgloss.JoinHorizontal(
		lipgloss.Left,
		skillsSectionsStyle.Render(skillsSection1),
		separator,
		skillsSectionsStyle.Render(skillsSection2),
	)

	aboutText := boldStyle.Render("# About Me") + `

I'm a web developer/programmer who likes to build websites and services as a passion.

I live in a small village in Gujarat India, my age is only 15 but I don't let that be a problem for me, I have built several websites, and although they are far from perfect but as a 15 y/o I'm proud of that achievement.

I want to explore deep into web/app development and expand my skill set I am also exploring all the technologies for development, that way I can ship my apps to more sets of devices.

My goal is to convert this passion of mine to a profession and create apps/websites that add value to people's lives.

I'm a web developer/programmer who likes to build websites and services as a passion.

I live in a small village in Gujarat India, my age is only 15 but I don't let that be a problem for me, I have built several websites, and although they are far from perfect but as a 15 y/o I'm proud of that achievement.

I want to explore deep into web/app development and expand my skill set I am also exploring all the technologies for development, that way I can ship my apps to more sets of devices.

My goal is to convert this passion of mine to a profession and create apps/websites that add value to people's lives.`

	aboutText += "\n\n" + skillsSection

	aboutContent := lipgloss.NewStyle().
		Width(usableWidth).
		Height(usableHeight).
		Render(aboutText)

	totalLines := lipgloss.Height(aboutContent)
	scrollViewTotalLines = totalLines
	scrollViewUsableHeight = usableHeight
	if totalLines > usableHeight {
		aboutContentLines := strings.Split(aboutContent, "\n")
		startLine := m.scrollOffset
		endLine := min(int(startLine)+usableHeight, totalLines)
		aboutContent = strings.Join(aboutContentLines[startLine:endLine], "\n")
	}

	topStyle := lipgloss.NewStyle().
		Width(usableWidth).
		Height(1).
		Align(lipgloss.Center)

	bottomStyle := lipgloss.NewStyle().
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
	return aboutStyle.Render(lipgloss.Place(
		width, height,
		lipgloss.Center, lipgloss.Center,
		top+"\n"+aboutContent+"\n"+bottom,
	))
}
