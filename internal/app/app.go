package app

import (
	canasta "github.com/Slintox/kanacta-game/internal/service"
)

func Run() {
	kanacta := canasta.GenerateKanasta(42, 250)

	kanacta.FillCards()

	//unifier := time.Now().Unix()
	//rand.Seed(unifier)
	//
	//uniqueKey := rand.Int63n(1)

	//kanacta.SwapCanasta(42, 250)

	kanacta.PrintCanasta()
}
