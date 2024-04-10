package ui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func RenderTable(columns []string, rows [][]string) {
	var baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("229"))

	var tableRows []table.Row

	for _, row := range rows {
		tableRows = append(tableRows, row)
	}

	var tableHeaders []table.Column

	for _, columnName := range columns {
		tableHeaders = append(tableHeaders, table.Column{Title: columnName, Width: 40})
	}

	t := table.New(
		table.WithColumns(tableHeaders),
		table.WithRows(tableRows),
		table.WithHeight(len(tableRows)+1),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("229")).
		BorderBottom(true).
		Bold(true)

	s.Selected = lipgloss.NewStyle()

	t.SetStyles(s)

	println(baseStyle.Render(t.View()))

}
