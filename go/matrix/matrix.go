// This package implements matrix parsing utilities, to generate arrays
// from string representations of matrices.
package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const KEY_FORMAT string = "%v,%v"
const ROW_SEPARATOR string = "\n"
const COLUMN_SEPARATOR string = " "

//  struct representation of a matrix, keep track of its dimensions, as
//  well as its values, in a map with string keys formatted with
//  KEY_FORMAT
type matrix struct {
	numCols int
	numRows int
	values  map[string]int
}

//  New creates a new matrix object based on a string representation of
//  said matrix. This representation contains one or more non empty
//  lines, each containing the same number of integer values, separated
//  by spaces.
func New(matrixString string) (*matrix, error) {
	var values = make(map[string]int)
	var numCols int = 0
	var err error = nil

	rows := strings.Split(matrixString, ROW_SEPARATOR)
	var numRows int = len(rows)

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
			for colInd, value := range rowCells {
				key := fmt.Sprintf(KEY_FORMAT, rowInd, colInd)
				intValue, convErr := strconv.Atoi(value)
				if convErr != nil {
					err = convErr
					break
				}
				values[key] = intValue
			}
		}
	}

	if err == nil {
		m := matrix{numCols, numRows, values}
		return &m, err
	}

	return nil, err
}

//  Rows returns the list of the rows of a matrix.
func (m *matrix) Rows() [][]int {
	var rows = make([][]int, m.numRows)
	for row := 0; row < m.numRows; row++ {
		for col := 0; col < m.numCols; col++ {
			key := fmt.Sprintf(KEY_FORMAT, row, col)
			rows[row] = append(rows[row], m.values[key])
		}
	}
	return rows
}

//  Cols returns the list of the columns of a matrix
func (m *matrix) Cols() [][]int {
	var columns = make([][]int, m.numCols)
	for col := 0; col < m.numCols; col++ {
		for row := 0; row < m.numRows; row++ {
			key := fmt.Sprintf(KEY_FORMAT, row, col)
			columns[col] = append(columns[col], m.values[key])
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
	if row >= 0 && row < m.numRows {
		if column >= 0 && column < m.numCols {
			key := fmt.Sprintf(KEY_FORMAT, row, column)
			m.values[key] = value
			ok = true
		}
	}
	return ok
}
