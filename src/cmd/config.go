package cmd

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print Spanny current configuration",
	Run: func(cmd *cobra.Command, args []string) {

		var baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.DoubleBorder()).
			BorderForeground(lipgloss.Color("229"))

		columns := []table.Column{
			{Title: "Key", Width: 20},
			{Title: "Value", Width: 30},
		}

		rows := []table.Row{
			{"Env", config.Env},
			{"ProjectId", config.ProjectId},
			{"InstanceId", config.InstanceId},
			{"DatabaseId", config.DatabaseId},
			{"MigrationFilesPath", config.MigrationFilesPath},
			{"ServicePath", config.ServicePath},
			{"UsingSpannerEmulator", strconv.FormatBool(config.UsingSpannerEmulator)},
			{"Port", fmt.Sprint(config.Port)},
		}

		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithHeight(len(rows)+1),
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
	},
}
