package main

import (
	"fmt"
	"math/rand"
)

type PlayerData struct {
	PlayerId     int
	StartingHand Deck
	HandData     HandResult
}

// Generates a pseudo random int between 2-8.
func getPlayerCount() int {
	return rand.Intn(7) + 2
}

// Returns a slice of PlayerData structs.
func createPlayers(n int) []PlayerData {
	playerSlice := []PlayerData{}

	for i := 1; i <= n; i++ {
		currentPlayerData := PlayerData{
			PlayerId:     i,
			StartingHand: Deck{},
			HandData:     HandResult{},
		}
		playerSlice = append(playerSlice, currentPlayerData)
	}

	return playerSlice
}

// Hands out 2 cards to each player.
// Returns the slice containing all players' data.
// Returns a deck with the reamaining undealt cards.
func dealToPlayers(s []PlayerData, d Deck) ([]PlayerData, Deck) {
	n := len(s) // player count
	cardsForPlayers, d := deal(d, n*2)

	for idx := range s {
		s[idx].StartingHand = append(
			s[idx].StartingHand, cardsForPlayers[idx], cardsForPlayers[idx+n])
	}

	return s, d
}

// Prints each player's data from the slice in the following format:
// Player: 2 | King of Spades | King of Diamonds
func printPlayersHands(s []PlayerData) {
	for _, p := range s {
		fmt.Println("Player:", p.PlayerId, "|", p.StartingHand[0].Value,
			"of", p.StartingHand[0].Suit, "|", p.StartingHand[1].Value, "of",
			p.StartingHand[1].Suit)
	}
}

// Compares the cards of p1 and p2.
// Returns the pidx of the winner, or -1 in the event of a draw.
func compareHands(s []PlayerData, p1idx int, p2idx int) int {
	hr1, hr2 := s[p1idx].HandData, s[p2idx].HandData

	if hr1.CombinationId > hr2.CombinationId {
		return p1idx
	} else if hr1.CombinationId < hr2.CombinationId {
		return p2idx
	} else {
		for i := 0; i < 5; i++ {
			p1Card, p2Card := hr1.PlayerHand[i], hr2.PlayerHand[i]
			if p1Card.Rank > p2Card.Rank {
				return p1idx
			} else if p1Card.Rank < p2Card.Rank {
				return p2idx
			}
		}
	}
	return -1
}

// Returns a slice of indices associated with the best hand(s).
func determineWinner(s []PlayerData) []int {
	winners := []int{}
	highestScore := 0

	for i, pd := range s {
		if pd.HandData.CombinationId > highestScore {
			winners = []int{i}
			highestScore = pd.HandData.CombinationId
		} else if pd.HandData.CombinationId == highestScore {
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
