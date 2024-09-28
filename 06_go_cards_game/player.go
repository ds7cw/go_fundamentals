package main

import (
	"fmt"
	"math/rand"
)

type playerData struct {
	playerId     int
	startingHand deck
	handData     handResult
}

// Generates a pseudo random int between 2-8.
func getPlayerCount() int {
	return rand.Intn(7) + 2
}

// Returns a slice of playerData structs.
func createPlayers(n int) []playerData {
	playerSlice := []playerData{}

	for i := 1; i <= n; i++ {
		currentPlayerData := playerData{
			playerId:     i,
			startingHand: deck{},
			handData:     handResult{},
		}
		playerSlice = append(playerSlice, currentPlayerData)
	}

	return playerSlice
}

// Hands out 2 cards to each player.
// Returns the slice containing all players' data.
// Returns a deck with the reamaining undealt cards.
func dealToPlayers(s []playerData, d deck) ([]playerData, deck) {
	n := len(s) // player count
	cardsForPlayers, d := deal(d, n*2)

	for idx := range s {
		s[idx].startingHand = append(
			s[idx].startingHand, cardsForPlayers[idx], cardsForPlayers[idx+n])
	}

	return s, d
}

// Prints each player's data from the slice in the following format:
// Player: 2 | King of Spades | King of Diamonds
func printPlayersHands(s []playerData) {
	for _, p := range s {
		fmt.Println("Player:", p.playerId, "|", p.startingHand[0].value,
			"of", p.startingHand[0].suit, "|", p.startingHand[1].value, "of",
			p.startingHand[1].suit)
	}
}
