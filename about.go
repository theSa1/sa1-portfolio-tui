package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Tool struct {
	Name string
	URL  string
}

type Section struct {
	Title string
	Tools []Tool
}

var toolsList = []Section{
	{
		Title: "Languages",
		Tools: []Tool{
			{Name: "Go", URL: "https://go.dev"},
			{Name: "JavaScript", URL: "https://developer.mozilla.org/en-US/docs/Web/JavaScript"},
			{Name: "TypeScript", URL: "https://www.typescriptlang.org"},
			{Name: "Python", URL: "https://www.python.org"},
		},
	},
	{
		Title: "Frameworks & Libraries",
		Tools: []Tool{
			{Name: "Node.js", URL: "https://nodejs.org"},
			{Name: "React", URL: "https://reactjs.org"},
			{Name: "Next.js", URL: "https://nextjs.org"},
			{Name: "Tailwind CSS", URL: "https://tailwindcss.com"},
		},
	},
	{
		Title: "Tools & Platforms",
		Tools: []Tool{
			{Name: "Git", URL: "https://git-scm.com"},
			{Name: "GitHub", URL: "https://github.com"},
			{Name: "Docker", URL: "https://www.docker.com"},
			{Name: "Kubernetes", URL: "https://kubernetes.io"},
			{Name: "AWS", URL: "https://aws.amazon.com"},
		},
	},
}

func getAboutView(m Model, width, height int) string {
	aboutStyle := lipgloss.NewStyle().
		Width(width).
		Height(height)

	boldStyle := lipgloss.NewStyle().
		Bold(true)

	usableHeight := height - 2
	usableWidth := width - 6

	aboutText := boldStyle.Render("# About Me") + `

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec at lectus a tellus sollicitudin vulputate. Aenean non quam porta, consectetur lacus quis, porttitor nisi. Quisque eget ipsum elit. Cras et pharetra tortor. Etiam iaculis enim cursus tempor fringilla. Vestibulum vitae cursus est. Vivamus ut mauris accumsan, laoreet mi porta, vulputate ante.

Ut tincidunt sapien dictum, finibus sem vitae, pretium ipsum. Sed massa mi, lobortis ut ligula eget, tempor viverra neque. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Donec a lectus eleifend, pulvinar tellus id, mattis nunc. Integer fermentum in metus id fringilla. Integer sed diam leo. Donec eu magna id purus commodo congue eget ac dolor. Sed vitae risus enim. Nam a ex vitae nisl accumsan gravida. Proin scelerisque suscipit turpis, quis facilisis erat varius nec. Morbi rhoncus nisl sit amet consequat mollis. Proin non egestas leo.`

	var toolsSections []string

	var toolStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).MarginRight(1)

	for _, section := range toolsList {
		var tools []string
		for _, tool := range section.Tools {
			tools = append(tools, toolStyle.Render(createHyperlink(tool.URL, tool.Name)))
		}
		toolsSections = append(toolsSections, "## "+section.Title+"\n"+lipgloss.JoinHorizontal(lipgloss.Left, tools...)+"\n")
	}

	aboutText += "\n\n\n" + boldStyle.Render("# Tools & Technologies") + "\n\n" +
		lipgloss.JoinVertical(lipgloss.Left, toolsSections...)

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
