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

//  rowParsing parses a matrix row represented as a string, and converts
//  it to a slice of integer values.
//  parameters :
//    * strRow : the string representation of the line, with cells
//               separated by COLUMN_SEPARATOR
func rowParsing(strRow string) ([]int, error) {
	returnedRow := make([]int, 0)

	for _, value := range strings.Split(strRow, COLUMN_SEPARATOR) {
		intValue, convErr := strconv.Atoi(value)
		if convErr != nil {
			return make([]int, 0), convErr
		}
		returnedRow = append(returnedRow, intValue)
	}
	return returnedRow, nil
}

//  New creates a new matrix object based on a string representation of
//  said matrix. This representation contains one or more non empty
//  lines, each containing the same number of integer values, separated
//  by spaces.
func New(matrixString string) (*matrix, error) {
	rows := strings.Split(matrixString, ROW_SEPARATOR)
	var numRows int = len(rows)
	var values = make([][]int, numRows)

	rowCells, parseError := rowParsing(rows[0])
	if parseError != nil {
		return nil, parseError
	}
	if len(rowCells) == 0 { // Check first line's dimension
		return nil, errors.New("Empty first row")
	}
	values[0] = rowCells
	var numCols = len(rowCells)

	for rowInd, row := range rows[1:] {
		row = strings.Trim(row, " ")
		rowCells, parseError := rowParsing(row)
		if parseError != nil {
			return nil, parseError
		}
		if len(rowCells) != numCols { // Check line size consistency
			return nil, errors.New("Inconsistent line numbers")
		}
		values[rowInd+1] = rowCells
	}

	m := matrix{values}
	return &m, nil
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

//  Set sets a cell in the matrix to a given value
//  parameters:
//    * row : row of the target cell
//    * column : columnof the target cell
//    * value : target value for the cell
func (m *matrix) Set(row int, column int, value int) bool {
	if row < 0 || row >= len(m.values) {
		return false
	}
	if column < 0 || column >= len(m.values[0]) {
		return false
	}
	m.values[row][column] = value
	return true
}
