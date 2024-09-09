package main

import (
	"os"
	"testing"
	"time"
)

// When running tests starting around the 8:00 timestamp,
// change the following
// t.Errorf("Expected deck length of 16, but got", len(d))
// to this:
// t.Errorf("Expected deck length of 16, but got %v", len(d))

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
		// --- FAIL: TestNewDeck (0.00s)
		// \deck_test.go:15: Expected deck length of 52, but got 52
		// FAIL
		// FAIL	cards	0.218s
		// ---
		// === RUN   TestNewDeck
		// --- PASS: TestNewDeck (0.00s)
		// PASS
		// ok      cards   0.219s
	}
	first_card := card{"Spades", "2", 1}
	if d[0] != first_card {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	last_card := card{"Clubs", "Ace", 52}
	if d[len(d)-1] != last_card {
		t.Errorf("Expected first card of King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	time.Sleep(time.Second * 3)
	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(d))
	}

	hand, remaining := deal(d, 26)

	if len(hand) != 26 {
		t.Errorf("Expected 26 cards in hand, got %v", len(hand))
	}

	if len(remaining) != 26 {
		t.Errorf("Expected 52 cards in remaining deck, got %v", len(remaining))
	}

	first_card := card{"Spades", "2", 1}
	if hand[0] != first_card {
		t.Errorf("Expected first card of Ace of Spades, but got %v", hand[0])
	}

	last_card := card{"Clubs", "Ace", 52}
	if remaining[len(remaining)-1] != last_card {
		t.Errorf("Expected first card of King of Clubs, but got %v", remaining[len(remaining)-1])
	}
}
