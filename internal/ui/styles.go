package ui

import "github.com/charmbracelet/lipgloss"

var (
	sunYellow  = lipgloss.Color("214")
	sunOrange  = lipgloss.Color("208")
	sunGold    = lipgloss.Color("220")
	warmWhite  = lipgloss.Color("252")
	skyBlue    = lipgloss.Color("39")
	grassGreen = lipgloss.Color("70")
	sandBeige  = lipgloss.Color("180")
	warmGray   = lipgloss.Color("245")
)

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(sunYellow).
			Padding(0, 2)

	headerBorderStyle = lipgloss.NewStyle().
				BorderBottom(true).
				BorderForeground(sunOrange)

	footerStyle = lipgloss.NewStyle().
			Padding(0, 1).
			BorderTop(true).
			BorderForeground(sandBeige).
			Foreground(warmGray)

	categoryStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(sunOrange)

	cursorStyle = lipgloss.NewStyle().
			Foreground(sunYellow)

	checkboxStyle = lipgloss.NewStyle().
			Foreground(sunGold)

	checkboxDoneStyle = lipgloss.NewStyle().
				Foreground(grassGreen).
				Bold(true)

	projectTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(sunYellow).
				Padding(0, 1)

	helpTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(sunOrange).
			Padding(0, 1)

	helpBorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(sunGold).
			Padding(1, 2)

	tableHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(sunYellow).
				Padding(0, 1)

	tableCellStyle = lipgloss.NewStyle().
			Padding(0, 1)

	tableSumStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(sunOrange).
			Padding(0, 1)

	linkStyle = lipgloss.NewStyle().
			Foreground(skyBlue)

	totalSumStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(grassGreen).
			Padding(0, 1)

	editTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(sunYellow).
			Padding(0, 1)

	focusedLabelStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(sunYellow)

	shortcutStyle = lipgloss.NewStyle().
			Foreground(warmWhite)

	deleteStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196"))

	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("196")).
			Padding(1, 2).
			Margin(1)

	dialogButtonStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("196")).
				Foreground(lipgloss.Color("255")).
				Bold(true).
				Padding(0, 2)

	dialogButtonInactiveStyle = lipgloss.NewStyle().
					Background(lipgloss.Color("245")).
					Foreground(lipgloss.Color("255")).
					Bold(true).
					Padding(0, 2)

	categorySelectStyle = lipgloss.NewStyle().
				Foreground(skyBlue)

	categorySelectedStyle = lipgloss.NewStyle().
				Foreground(grassGreen).
				Bold(true)
)
