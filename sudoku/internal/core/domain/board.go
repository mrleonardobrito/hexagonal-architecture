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

	for i := range board {
		for _, region := range board[i] {

			boardString += "\n"
		}
		boardString += "----------------------- \n"
	}

	return boardString
}
