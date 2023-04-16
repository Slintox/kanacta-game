package app

import (
	"github.com/Slintox/kanacta-game/internal/service"
)

func Run() {
	kanacta := service.GenerateKanasta(42, 250)

	//unifier := time.Now().Unix()
	//rand.Seed(unifier)
	//
	//uniqueKey := rand.Int63n(1)

	//kanacta.SwapCanasta(42, 250)

	kanacta.PrintCanasta()
}
