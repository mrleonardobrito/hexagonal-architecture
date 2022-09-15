package main

import (
	gamesrv "sodouku/internal/core/service"
	"sodouku/internal/handlers/gamehdl"
)

func main() {
	gameService := gamesrv.New()
	gameHandler := gamehdl.NewCMDHandler(gameService)

	gameHandler.NewGame("jogo teste", 3, "easy")
}
