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
}
