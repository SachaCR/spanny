package dbops

import (
	"strconv"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/types/known/structpb"
)

func rowIterToString(iter *spanner.RowIterator) (columns []string, rows [][]string, err error) {
	for {
		row, err := iter.Next()

		if err != nil {
			if err == iterator.Done {
				break
			}

			return nil, nil, err
		}

		if len(columns) == 0 {
			columns = row.ColumnNames()
		}

		var values []string

		for index := range columns {
			value := row.ColumnValue(index)
			_, isBool := value.Kind.(*structpb.Value_BoolValue)
			_, isNumber := value.Kind.(*structpb.Value_NumberValue)

			if isBool {
				values = append(values, strconv.FormatBool(value.GetBoolValue()))
				continue
			}

			if isNumber {
				values = append(values, strconv.FormatFloat(value.GetNumberValue(), 'f', -2, 64))
				continue
			}

			values = append(values, value.GetStringValue())
		}

		rows = append(rows, values)
	}

	return columns, rows, nil
}
