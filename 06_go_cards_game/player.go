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

// Compares the cards of p1 and p2.
// Returns the pidx of the winner, or -1 in the event of a draw.
func compareHands(s []playerData, p1idx int, p2idx int) int {
	hr1, hr2 := s[p1idx].handData, s[p2idx].handData

	if hr1.combinationId > hr2.combinationId {
		return p1idx
	} else if hr1.combinationId < hr2.combinationId {
		return p2idx
	} else {
		for i := 0; i < 5; i++ {
			p1Card, p2Card := hr1.playerHand[i], hr2.playerHand[i]
			if p1Card.rank > p2Card.rank {
				return p1idx
			} else if p1Card.rank < p2Card.rank {
				return p2idx
			}
		}
	}
	return -1
}

// Returns a slice of indices associated with the best hand(s).
func determineWinner(s []playerData) []int {
	winners := []int{}
	highestScore := 0

	for i, pd := range s {
		if pd.handData.combinationId > highestScore {
			winners = []int{i}
			highestScore = pd.handData.combinationId
		} else if pd.handData.combinationId == highestScore {
			winningIdx := compareHands(s, i, winners[0])
			if winningIdx < 0 {
				winners = append(winners, i)
			} else if winningIdx == i {
				winners = []int{i}
			}
		}
	}
	return winners
}
