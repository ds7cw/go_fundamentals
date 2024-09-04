package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	playersCount := getPlayerCount()
	fmt.Println(playersCount)

	playersMap := createPlayers(playersCount)
	fmt.Println(playersMap)

	playersMap = dealToPlayers(playersMap, cards)
	printPlayersHands(playersMap)

}
