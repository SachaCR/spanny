package ui

import (
	"strconv"

	"cloud.google.com/go/spanner"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/types/known/structpb"
)

func RenderTableFromRowIterator(iter *spanner.RowIterator) {
	var baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("229"))

	var columns []string
	var tableRows []table.Row

	for {
		row, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			println(err.Error())
			return
		}

		if len(columns) == 0 {
			columns = row.ColumnNames()
		}

		var values []string

		for index := range columns {
			value := row.ColumnValue(index)

			_, isBool := value.GetKind().(*structpb.Value_BoolValue)
			if isBool {
				values = append(values, strconv.FormatBool(value.GetBoolValue()))
				continue
			}
			values = append(values, value.GetStringValue())

		}

		tableRows = append(tableRows, values)
	}

	if len(tableRows) == 0 {
		println("No rows to display")
		return
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
