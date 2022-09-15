package domain

const (
	GAME_STATE_NEW  = "new"
	GAME_STATE_WON  = "won"
	GAME_STATE_LOST = "lost"
)

const (
	GAME_DIFFICULT_EASY   = "easy"
	GAME_DIFFICULT_MEDIUM = "medium"
	GAME_DIFFICULT_HARD   = "hard"
)

type Game struct {
	id        string
	name      string
	size      uint
	state     string
	dificulty string
	Board     Board
}

func NewGame(id string, name string, size uint, dificulty string) Game {
	return Game{
		id:        id,
		name:      name,
		size:      size,
		state:     GAME_STATE_NEW,
		dificulty: dificulty,
		Board:     NewEmptyBoard(size),
	}
}
