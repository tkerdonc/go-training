// This package implements matrix parsing utilities, to generate arrays
// from string representations of matrices.
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

const ROW_SEPARATOR string = "\n"
const COLUMN_SEPARATOR string = " "

//  struct representation of a matrix
type matrix struct {
	values [][]int
}

//  New creates a new matrix object based on a string representation of
//  said matrix. This representation contains one or more non empty
//  lines, each containing the same number of integer values, separated
//  by spaces.
func New(matrixString string) (*matrix, error) {
	var numCols int = 0
	var err error = nil

	rows := strings.Split(matrixString, ROW_SEPARATOR)
	var numRows int = len(rows)
	var values = make([][]int, numRows)

	for rowInd, row := range rows {
		row = strings.Trim(row, " ")
		rowCells := strings.Split(row, COLUMN_SEPARATOR)
		if rowInd == 0 {
			numCols = len(rowCells)
			if numCols == 0 { // Check first line's dimension
				err = errors.New("Empty first row")
			}
		} else if len(rowCells) != numCols { // Check line size consistency
			err = errors.New("Inconsistent line numbers")
		}

		if err == nil {
			for _, value := range rowCells {
				intValue, convErr := strconv.Atoi(value)
				if convErr != nil {
					err = convErr
					break
				}
				values[rowInd] = append(values[rowInd], intValue)
			}
		}
	}

	if err == nil {
		m := matrix{values}
		return &m, err
	}

	return nil, err
}

//  Rows returns the list of the rows of a matrix.
func (m *matrix) Rows() [][]int {
	var rows = make([][]int, len(m.values))
	for row := 0; row < len(m.values); row++ {
		for col := 0; col < len(m.values[0]); col++ {
			rows[row] = append(rows[row], m.values[row][col])
		}
	}
	return rows
}

//  Cols returns the list of the columns of a matrix
func (m *matrix) Cols() [][]int {
	var columns = make([][]int, len(m.values[0]))
	for col := 0; col < len(m.values[0]); col++ {
		for row := 0; row < len(m.values); row++ {
			columns[col] = append(columns[col], m.values[row][col])
		}
	}
	return columns
}

//  Set target a cell of a matrix by and sets its value to the one
//  passed as para
//  parameters:
//    * row : row of the target cell
//    * column : columnof the target cell
//    * value : target value for the cell
func (m *matrix) Set(row int, column int, value int) bool {
	var ok = false
	if row >= 0 && row < len(m.values) {
		if column >= 0 && column < len(m.values[0]) {
			m.values[row][column] = value
			ok = true
		}
	}
	return ok
}
