package main

import (
	"log"
	gamesrv "sodouku/internal/core/service"
	"sodouku/internal/handlers/gamehdl"
	"time"
)

func main() {
	start := time.Now()
	gameService := gamesrv.New()
	gameHandler := gamehdl.NewCMDHandler(gameService)

	gameHandler.NewGame("jogo teste", 2, "easy")
	timeTrack(start, "main")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start).String()
	log.Printf("%s took %s", name, elapsed)
}
