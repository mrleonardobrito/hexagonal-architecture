package domain

import "math"

type Board []Region

func NewEmptyBoard(size uint) Board {
	board := make([]Region, size*size)

	for i := range board {
		board[i] = NewEmptyRegion("easy")
	}

	return board
}

func (board Board) ToString() string {
	boardString := ""
	size := int(math.Sqrt(float64(len(board))))

	for i := 0; i < len(board); i += size {
		row := board[i : i+int(size)]
		renderedRow := renderBoardRow(row)
		boardString += renderedRow
		boardString += "-------------\n"
	}
	return boardString
}

func renderBoardRow(row []Region) string {
	rowString := ""

	for i := 0; i < 3; i++ {
		for j := range row {
			rowString += row[j].RenderLine(i)
		}
		rowString += "\n"
	}

	return rowString
}
