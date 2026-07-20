package ui

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4"))

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#24A0ED"))

	ProjectStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			Italic(true)

	ListItemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	SelectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("#7D56F4")).
				Bold(true)

	CompletedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575"))

	CategoryStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF79C6")).
			MarginTop(1)

	StatusBarStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#24A0ED")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 1)

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			Padding(1, 2)
)
