package gamesrv

import (
	"sodouku/internal/core/domain"
)

type service struct{}

func New() *service {
	return &service{}
}

func (srv *service) Create(name string, size uint, dificulty string) (domain.Game, error) {
	newGame := domain.NewGame("10001", name, size, domain.GAME_DIFFICULT_EASY)
	return newGame, nil
}
