package main

import (
	"fmt"
	"net/http"
	"slices"
)

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	playersCount := getPlayerCount()
	fmt.Println("\nPlayers Count:", playersCount)

	playersSlice := createPlayers(playersCount)

	playersSlice, cards = dealToPlayers(playersSlice, cards)
	printPlayersHands(playersSlice)

	flop, cards := dealFlop(cards)
	turn, cards := dealTurn(cards)
	river, cards := dealRiver(cards)
	communityCards := slices.Concat(flop, turn, river)

	fmt.Println("\nCommunity Cards:\n", communityCards)
	fmt.Println("\nRemaining Cards in Deck:", len(cards))

	// Evaluate all players' hands
	for i, pc := range playersSlice {
		currentPlayerCards := slices.Concat(communityCards, pc.startingHand)
		playersSlice[i].handData = currentPlayerCards.evaluateHand()
		fmt.Println("Player", playersSlice[i].playerId,
			"ComboID:", playersSlice[i].handData.combinationId,
			"Hand:", playersSlice[i].handData.playerHand)
	}

	// Determine winners
	winningPlayersIdx := determineWinner(playersSlice)
	// Print winning hand(s)
	for _, n := range winningPlayersIdx {
		cpd := playersSlice[n]
		fmt.Println(
			"Winner:", cpd.playerId,
			"ComboID:", cpd.handData.combinationId,
			"Cards:", cpd.handData.playerHand,
		)
	}

	http.HandleFunc("/new-game", gameHandleFunc)
	http.ListenAndServe(":8080", nil)

}

func gameHandleFunc(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to the Poker Simulator!")
	cards := newDeck()
	cards.shuffle()

	playersCount := getPlayerCount()
	playersSlice := createPlayers(playersCount)

	playersSlice, cards = dealToPlayers(playersSlice, cards)
	flop, cards := dealFlop(cards)
	turn, cards := dealTurn(cards)
	river, _ := dealRiver(cards)
	communityCards := slices.Concat(flop, turn, river)

	for i, pc := range playersSlice {
		currentPlayerCards := slices.Concat(communityCards, pc.startingHand)
		playersSlice[i].handData = currentPlayerCards.evaluateHand()
		fmt.Println("Player", playersSlice[i].playerId,
			"ComboID:", playersSlice[i].handData.combinationId,
			"Hand:", playersSlice[i].handData.playerHand)
	}

	fmt.Fprintf(w, "Community Cardz:\n%q\nPlayers:\n%q",
		communityCards, playersSlice)
}
