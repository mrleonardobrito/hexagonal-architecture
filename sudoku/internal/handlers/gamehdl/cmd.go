package gamehdl

import (
	"fmt"
	"sodouku/internal/core/domain"
	"sodouku/internal/core/ports"
)

type CMDHandler struct {
	gameService ports.GameService
}

func NewCMDHandler(gameService ports.GameService) *CMDHandler {
	return &CMDHandler{
		gameService: gameService,
	}
}

func (hdl *CMDHandler) NewGame(name string, size uint, dificulty string) {
	game, err := hdl.gameService.Create(name, size, dificulty)

	if err != nil {
		return
	}

	ShowBoard(game.Board)
}

func ShowBoard(board domain.Board) {
	region0 := &board[0][0]
	*region0 = domain.Region{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	fmt.Println(board.ToString())
}
