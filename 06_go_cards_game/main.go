package main

import (
	"fmt"
	"html/template"
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
		currentPlayerCards := slices.Concat(communityCards, pc.StartingHand)
		playersSlice[i].HandData = currentPlayerCards.evaluateHand()
		fmt.Println("Player", playersSlice[i].PlayerId,
			"ComboID:", playersSlice[i].HandData.CombinationId,
			"Hand:", playersSlice[i].HandData.PlayerHand)
	}

	// Determine winners
	winningPlayersIdx := determineWinner(playersSlice)
	// Print winning hand(s)
	for _, n := range winningPlayersIdx {
		cpd := playersSlice[n]
		fmt.Println(
			"Winner:", cpd.PlayerId,
			"ComboID:", cpd.HandData.CombinationId,
			"Cards:", cpd.HandData.PlayerHand,
		)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/new-game", gameHandleFunc)
	http.ListenAndServe(":8080", nil)

}

type TemplateContext struct {
	Players     []PlayerData
	FaceUpCards Deck
	Winners     []int
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
		currentPlayerCards := slices.Concat(communityCards, pc.StartingHand)
		playersSlice[i].HandData = currentPlayerCards.evaluateHand()
		fmt.Println("Player", playersSlice[i].PlayerId,
			"ComboID:", playersSlice[i].HandData.CombinationId,
			"Hand:", playersSlice[i].HandData.PlayerHand)
	}

	ctx := TemplateContext{Players: playersSlice, FaceUpCards: communityCards}
	tmpl := template.Must(template.ParseFiles("templates/new-game.html"))
	tmpl.Execute(w, ctx)

	// fmt.Fprintf(w, "Community Cardz:\n%q\nPlayers:\n%q",
	// 	communityCards, playersSlice)
}
