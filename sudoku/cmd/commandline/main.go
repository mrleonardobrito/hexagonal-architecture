package main

import (
	"log"
	gamesrv "sodouku/internal/core/service"
	"sodouku/internal/handlers/gamehdl"
	"time"
)

func main() {
	time := time.Now()
	gameService := gamesrv.New()
	gameHandler := gamehdl.NewCMDHandler(gameService)

	gameHandler.NewGame("jogo teste", 4, "easy")

	timeTrack(time, "main")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start).String()
	log.Printf("%s took %s", name, elapsed)
}
