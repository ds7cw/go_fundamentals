package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// Tests the getPlayerCount func. Must return an int between 2 and 8.
func TestGetPlayerCount(t *testing.T) {
	for i := 0; i < 5; i++ {
		result := getPlayerCount()
		if result < 2 || result > 8 {
			t.Errorf("getPlayerCount() = %d; expected an int between 2 and 8", result)
		}
	}
}

// Tests the CreatePlayers func. Must return a slice with size equal to the input int.
func TestCreatePlayers(t *testing.T) {
	count := 3
	resultSlice := createPlayers(count)
	if len(resultSlice) != count {
		t.Errorf("createPlayers(3) returns a slice of length = %d, expected %d ",
			len(resultSlice), count)
	}

	if p3Id := resultSlice[2].playerId; p3Id != count {
		t.Errorf("Player 3 ID = %d, expected %d",
			resultSlice[2].playerId, count)
	}
}

// Tests dealToPlayers func. Each player must receive 2 cards.
// All remaining cards must be dealt to the deck parameter.
func TestDealToPlayers(t *testing.T) {
	playerSlice := createPlayers(2)
	testDeck := deck{
		card{suit: SuitSpades, value: "10", rank: 10},
		card{suit: SuitSpades, value: "Jack", rank: 10},
		card{suit: SuitHearts, value: "Queen", rank: 12},
		card{suit: SuitHearts, value: "King", rank: 13},
		card{suit: SuitDiamonds, value: "Ace", rank: 14},
	}
	playerSlice, remaining := dealToPlayers(playerSlice, testDeck)
	expected := []string{"Ace", "Diamonds"}
	if remaining[0].value != expected[0] || remaining[0].suit != expected[1] {
		t.Errorf("Expected remaining card to be %v of %v, got %v of %v instead",
			expected[0], expected[1], remaining[0].value, remaining[0].suit)
	}

	p1Cards := playerSlice[0].startingHand
	expected = []string{"10", "Spades"}
	if p1Cards[0].value != expected[0] || p1Cards[0].suit != expected[1] {
		t.Errorf("Expected P1 1st Card to be %v of %v, got %v of %v instead",
			expected[0], expected[1], p1Cards[0].value, p1Cards[0].suit)
	}

	expected = []string{"Queen", "Hearts"}
	if p1Cards[1].value != expected[0] || p1Cards[1].suit != expected[1] {
		t.Errorf("Expected P1 2nd Card to be %v of %v, got %v of %v instead",
			expected[0], expected[1], p1Cards[1].value, p1Cards[1].suit)
	}
}

// Tests printPlayersHands. Must correctly print the player ID & starting hand.
func TestPrintPlayersHands(t *testing.T) {
	playerSlice := createPlayers(1)
	playerSlice[0].startingHand = deck{
		card{SuitDiamonds, "Jack", 11},
		card{SuitHearts, "Jack", 11},
	}

	// Save the original stdout
	originalStdout := os.Stdout

	// Create a pipe to capture the output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	printPlayersHands(playerSlice)

	// Close the writer and restore stdout
	w.Close()
	os.Stdout = originalStdout

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Check the output
	got := buf.String()
	want := "Player: 1 | Jack of Diamonds | Jack of Hearts\n"

	if got != want {
		t.Errorf("PrintHelloWorld() = %q; want %q", got, want)
	}
}

// Tests the compareHands function.
// Must return the idx of the winnin player, if there is one, -1 otherwise.
func TestCompareHands(t *testing.T) {
	playerSlice := createPlayers(4)
	playerSlice[0].handData = handResult{
		combinationId: FullHouseId,
		playerHand: deck{
			card{SuitClubs, "Ace", 14},
			card{SuitSpades, "Ace", 14},
			card{SuitClubs, "King", 13},
			card{SuitSpades, "King", 13},
			card{SuitHearts, "King", 13},
		},
	}

	playerSlice[1].handData = handResult{
		combinationId: FlushId,
		playerHand: deck{
			card{SuitClubs, "Ace", 14},
			card{SuitClubs, "King", 13},
			card{SuitClubs, "Jack", 11},
			card{SuitClubs, "10", 10},
			card{SuitClubs, "9", 9},
		},
	}

	playerSlice[2].handData = handResult{
		combinationId: FullHouseId,
		playerHand: deck{
			card{SuitClubs, "Ace", 14},
			card{SuitSpades, "Ace", 14},
			card{SuitClubs, "King", 13},
			card{SuitSpades, "King", 13},
			card{SuitHearts, "King", 13},
		},
	}

	playerSlice[3].handData = handResult{
		combinationId: FlushId,
		playerHand: deck{
			card{SuitClubs, "Ace", 14},
			card{SuitClubs, "King", 13},
			card{SuitClubs, "Queen", 12},
			card{SuitClubs, "10", 10},
			card{SuitClubs, "9", 9},
		},
	}

	// Compare equal hands.
	expected := -1
	got := compareHands(playerSlice, 0, 2)
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}

	// Compare different hands when p1 > p2.
	expected = 0
	got = compareHands(playerSlice, 0, 1)
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}

	// Compare different hands when p2 > p1.
	expected = 2
	got = compareHands(playerSlice, 1, 2)
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}

	// Compare equal combination when when p2 has a high card.
	expected = 3
	got = compareHands(playerSlice, 1, 3)
	if expected != got {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestDetermineWinner(t *testing.T) {
	playerSlice := createPlayers(3)
	playerSlice[0].handData = handResult{
		combinationId: FullHouseId,
		playerHand: deck{
			card{SuitClubs, "King", 13},
			card{SuitSpades, "King", 13},
			card{SuitHearts, "King", 13},
			card{SuitSpades, "Ace", 14},
			card{SuitHearts, "Ace", 14},
		},
	}

	playerSlice[1].handData = handResult{
		combinationId: FullHouseId,
		playerHand: deck{
			card{SuitClubs, "Ace", 14},
			card{SuitSpades, "Ace", 14},
			card{SuitDiamonds, "Ace", 14},
			card{SuitSpades, "King", 13},
			card{SuitHearts, "King", 13},
		},
	}

	playerSlice[2].handData = handResult{
		combinationId: FullHouseId,
		playerHand: deck{
			card{SuitClubs, "Ace", 14},
			card{SuitSpades, "Ace", 14},
			card{SuitDiamonds, "Ace", 14},
			card{SuitSpades, "King", 13},
			card{SuitHearts, "King", 13},
		},
	}

	winners := determineWinner(playerSlice)
	expected := []int{1, 2}

	if len(winners) != len(expected) {
		t.Errorf("%d winners expected, got %d",
			len(expected), len(winners))
	}

	if winners[0] != expected[0] || winners[1] != expected[1] {
		t.Errorf("Expected slice of winners %v, got %v", expected, winners)
	}
}
