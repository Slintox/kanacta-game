package service

import (
	"fmt"
	"math/rand"
)

type CardType string

const (
	CardTypeDefault CardType = "default"
	CardTypeTrump   CardType = "trump"
)

// Suits
const (
	CardSuitSpades   = "Пики"
	CardSuitHearts   = "Черви"
	CardSuitClubs    = "Трефы"
	CardSuitDiamonds = "Бубны"

	CardSuitSymbolSpades   = "♠"
	CardSuitSymbolHearts   = "♥"
	CardSuitSymbolClubs    = "♣"
	CardSuitSymbolDiamonds = "♦"
)

type CardSuit struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func GenerateCardSuits() []CardSuit {
	var cs []CardSuit

	cs = []CardSuit{
		{
			Name:   CardSuitSpades,
			Symbol: CardSuitSymbolSpades,
		},
		{
			Name:   CardSuitHearts,
			Symbol: CardSuitSymbolHearts,
		},
		{
			Name:   CardSuitClubs,
			Symbol: CardSuitSymbolClubs,
		},
		{
			Name:   CardSuitDiamonds,
			Symbol: CardSuitSymbolDiamonds,
		},
	}

	return cs
}

// Values
const (
	CardValue2 = iota + 2
	CardValue3
	CardValue4
	CardValue5
	CardValue6
	CardValue7
	CardValue8
	CardValue9
	CardValueJ
	CardValueQ
	CardValueK
	CardValueA
	CardValueJoker
)

// Names
const (
	CardName2     = "2"
	CardName3     = "3"
	CardName4     = "4"
	CardName5     = "5"
	CardName6     = "6"
	CardName7     = "7"
	CardName8     = "8"
	CardName9     = "9"
	CardNameJ     = "J"
	CardNameQ     = "Q"
	CardNameK     = "K"
	CardNameA     = "A"
	CardNameJoker = "*"
)

// Costs
const (
	CardCostSmall  = 5
	CardCostMiddle = 10
	CardCostBig    = 20
	CardCostJoker  = 50
)

type Card struct {
	CardSuit *CardSuit `json:"cardSuit"`
	Value    int       `json:"value"`
	Name     string    `json:"name"`
	Cost     int       `json:"cost"`
	Type     CardType  `json:"type"`
}

var UniqueCards = []Card{
	{
		Value: CardValue2,
		Name:  CardName2,
		Cost:  CardCostSmall,
		Type:  CardTypeTrump,
	},
	{
		Value: CardValue3,
		Name:  CardName3,
		Cost:  CardCostSmall,
		Type:  CardTypeDefault,
	},
	{
		Value: CardValue4,
		Name:  CardName4,
		Cost:  CardCostSmall,
		Type:  CardTypeDefault,
	},
	{
		Value: CardValue5,
		Name:  CardName5,
		Cost:  CardCostSmall,
		Type:  CardTypeDefault,
	},
	{
		Value: CardValue6,
		Name:  CardName6,
		Cost:  CardCostSmall,
		Type:  CardTypeDefault,
	},
	{
		Value: CardValue7,
		Name:  CardName7,
		Cost:  CardCostSmall,
		Type:  CardTypeDefault,
	},
	{
		Value: CardValue8,
		Name:  CardName8,
		Cost:  CardCostMiddle,
		Type:  CardTypeDefault,
	},
	{
		Value: CardValue9,
		Name:  CardName9,
		Cost:  CardCostMiddle,
		Type:  CardTypeDefault,
	},
	{
		Value: CardValueJ,
		Name:  CardNameJ,
		Cost:  CardCostMiddle,
		Type:  CardTypeDefault,
	},
	{
		Value: CardValueQ,
		Name:  CardNameQ,
		Cost:  CardCostMiddle,
		Type:  CardTypeDefault,
	},
	{
		Value: CardValueK,
		Name:  CardNameK,
		Cost:  CardCostMiddle,
		Type:  CardTypeDefault,
	},
	{
		Value: CardValueA,
		Name:  CardNameA,
		Cost:  CardCostBig,
		Type:  CardTypeDefault,
	},
	{
		Value: CardValueJoker,
		Name:  CardNameJoker,
		Cost:  CardCostJoker,
		Type:  CardTypeTrump,
	},
}

type Canasta struct {
	cards []Card
}

func GenerateKanasta(uniqueKey int64, swapsNum int) *Canasta {
	kanasta := &Canasta{}

	kanasta.SwapCanasta(uniqueKey, swapsNum)

	return kanasta
}

// FillCards заполняет канасту картами
func (k *Canasta) FillCards() {
	cSuits := GenerateCardSuits()

	for i := range cSuits {
		for _, uniqueCard := range UniqueCards {

			uniqueCard.CardSuit = &cSuits[i]

			k.cards = append(k.cards, uniqueCard)
		}
	}
}

// SwapCanasta перемешивает канасту
func (k *Canasta) SwapCanasta(uniqueKey int64, swapsNum int) {
	// Дефолтное значение итераций перемешивания
	if swapsNum == -1 {
		swapsNum = 250
	}

	// Перемешиваем канасту
	rand.Seed(uniqueKey)
	for i := 0; i < swapsNum; i++ {
		f := rand.Intn(len(k.cards))
		s := rand.Intn(len(k.cards))

		// Меняем карты местами
		k.swapCards(f, s)
	}
}

// Меняет первую и вторую карты местами
func (k *Canasta) swapCards(f, s int) {
	k.cards[f], k.cards[s] = k.cards[s], k.cards[f]
}

// PrintCanasta вывод всех карт в канасте
func (k *Canasta) PrintCanasta() {
	// Вывод всех карт в канасте
	for i := range k.cards {
		fmt.Printf("%s%s\n", k.cards[i].Name, k.cards[i].CardSuit.Symbol)
	}
}
