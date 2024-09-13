package main

import (
	"fmt"
	"slices"
)

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	playersCount := getPlayerCount()
	fmt.Println("\nPlayers Count:", playersCount)

	playersMap := createPlayers(playersCount)

	playersMap, cards = dealToPlayers(playersMap, cards)
	printPlayersHands(playersMap)

	flop, cards := dealFlop(cards)
	turn, cards := dealTurn(cards)
	river, cards := dealRiver(cards)
	communityCards := slices.Concat(flop, turn, river)

	fmt.Println("\nCommunity Cards:\n", communityCards)
	fmt.Println("\nRemaining Cards in Deck:", len(cards))

	fmt.Println("Evaluate hand:", communityCards.evaluateHand())

	pairTesting()
}

// temporary function to test functionality
func pairTesting() {
	p1 := deck{
		card{suit: "Diamonds", value: "2", rank: 2},
		card{suit: "Spades", value: "6", rank: 6},
		card{suit: "Spades", value: "Ace", rank: 13},
		card{suit: "Hearts", value: "9", rank: 9},
		card{suit: "Diamonds", value: "7", rank: 7},
		card{suit: "Diamonds", value: "5", rank: 5},
		card{suit: "Clubs", value: "2", rank: 2},
	}
	p1.sortDeck()
	fmt.Println(p1)
	p1Result := p1.hasPair()
	fmt.Println(p1Result)
	// expected:
	// {2 [{Clubs 2 2} {Diamonds 2 2} {Spades Ace 13} {Hearts 9 9} {Diamonds 7 7}]}
}
