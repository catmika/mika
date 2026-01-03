package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	primaryColor   = lipgloss.Color("212") // pink/magenta
	secondaryColor = lipgloss.Color("86")  // cyan
	mutedColor     = lipgloss.Color("241") // gray

	// Styles
	bannerStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true)
)

const banner = `
        _ _
  _ __ (_) | ____ _
 | '_ \| | |/ / _' |
 | | | | |   < (_| |
 |_| |_|_|_|\_\__,_|
`

func PrintBanner() {
	fmt.Println(bannerStyle.Render(banner))
	fmt.Println(subtitleStyle.Render("Clean web project scaffolder"))
	fmt.Println()
}

func PrintSuccess(projectName string) {
	successStyle := lipgloss.NewStyle().
		Foreground(secondaryColor).
		Bold(true)

	fmt.Println()
	fmt.Println(successStyle.Render(fmt.Sprintf("✓ Project '%s' created successfully!", projectName)))
	fmt.Println()
	fmt.Println(mutedStyle().Render(fmt.Sprintf("  cd %s", projectName)))
	fmt.Println(mutedStyle().Render("  npm install"))
	fmt.Println(mutedStyle().Render("  npm run dev"))
	fmt.Println()
}

func mutedStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(mutedColor)
}
