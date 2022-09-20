package domain

type Board [][]Region

func NewEmptyBoard(size uint) Board {
	board := make([][]Region, size)

	for i := range board {
		board[i] = make([]Region, size)
	}

	for i := range board {
		for j := range board[0] {
			board[i][j] = NewEmptyRegion(GAME_DIFFICULT_EASY)
		}
	}

	return board
}

func (board Board) ToString() string {
	boardString := ""

	for _, row := range board {
		boardString += renderBoardRow(row)
		for i := range row {
			if i < len(row)-1 {
				boardString += "--------"
			} else {
				boardString += "-------"
			}
		}
		boardString += "\n"
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
