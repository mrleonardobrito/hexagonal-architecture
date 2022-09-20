package gamehdl

import (
	"errors"
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

func (hdl *CMDHandler) NewGame(name string, size uint, dificulty string) error {
	game, err := hdl.gameService.Create(name, size, dificulty)

	if err != nil {
		return errors.New("GameService failed to create a new game!!!")
	}

	ShowBoard(game.Board)

	return nil
}

func ShowBoard(board domain.Board) {
	fmt.Println(board.ToString())
}
