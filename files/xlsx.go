package files

import (
	"strings"

	xlsx "github.com/tealeg/xlsx/v3"
)

func ReadSheetSlice(sheet *xlsx.Sheet) [][]string {
	arr := make([][]string, 0, 100)
	sheet.ForEachRow(func(r *xlsx.Row) error {
		row := make([]string, 0, 20)
		r.ForEachCell(func(c *xlsx.Cell) error {
			row = append(row, strings.TrimSpace(c.Value))
			return nil
		})
		arr = append(arr, row)
		return nil
	})
	return arr
}
