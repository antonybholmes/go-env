package sys

import (
	"bytes"
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Table struct {
	index   []string   `json:"index"`
	columns []string   `json:"columns"`
	data    [][]string `json:"data"`
}

func xlsx_to_text(input []byte, index int, header int) (*Table, error) {
	r := bytes.NewReader(input)

	f, err := excelize.OpenReader(r) // .OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	colStart := max(index+1, 0)
	dataStart := max(header+1, 0)
	cols := len(rows) - colStart - 1

	indexNames := make([][]string, 0, 100)
	columns := make([][]string, cols)
	data := make([][]string, 0, 100)

	for i := 0; i < cols; i++ {
		columns[i] = make([]string, header)
	}

	for ri, row := range rows {
		if ri <= header {
			for i := colStart; i < cols; i++ {
				columns[i-colStart][ri] = row[i]
			}
		} else {
			indexNames = append(indexNames, make([]string, dataStart))

			for i := 0; i < index; i++ {
				indexNames[len(indexNames)-1][i] = row[i]
			}

			for i := index; i < cols; i++ {
				rowData := make([]string, cols)

				// for ci, colCell := range row {
				// 	fmt.Print(colCell, "\t")
				// }

				data = append(data, rowData)
			}

		}
	}

	return nil, nil
}
