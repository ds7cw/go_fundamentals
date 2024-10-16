package main

import (
	"bytes"
	"io"
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
	fc := Card{SuitClubs, "2", 2}
	if d[0].Suit != fc.Suit || d[0].Rank != fc.Rank || d[0].Value != fc.Value {
		t.Errorf("Expected first %v of %v, but got %v of %v",
			fc.Value, fc.Suit, d[0].Value, d[0].Suit)
	}

	lc := Card{SuitSpades, "Ace", 14}
	if d[51].Suit != lc.Suit || d[51].Rank != lc.Rank || d[51].Value != lc.Value {
		t.Errorf("Expected first %v of %v, but got %v of %v",
			lc.Value, lc.Suit, d[51].Value, d[51].Suit)
	}
}

func TestPrintDeck(t *testing.T) {
	d := Deck{Card{SuitSpades, "Jack", 11}}

	// Save the original stdout
	originalStdout := os.Stdout

	// Create a pipe to capture the output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	d.print()

	// Close the writer and restore stdout
	w.Close()
	os.Stdout = originalStdout

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Check the output
	got := buf.String()
	want := "0 {Spades Jack 11}\n"

	if got != want {
		t.Errorf("Deck.print() = %q; want %q", got, want)
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

	fc := Card{SuitClubs, "2", 2}
	if hand[0] != fc {
		t.Errorf("Expected first card %v of %v, but got %v of %v",
			fc.Value, fc.Suit, hand[0].Value, hand[0].Suit)
	}

	lc := Card{SuitSpades, "Ace", 14}
	if remaining[len(remaining)-1] != lc {
		t.Errorf("Expected last card %v of %v, but got %v of %v",
			fc.Value, fc.Suit, remaining[25].Value, remaining[25].Suit)
	}
}

func TestShuffle(t *testing.T) {
	d := Deck{
		Card{SuitSpades, "2", 2},
		Card{SuitSpades, "3", 3},
		Card{SuitSpades, "4", 4},
		Card{SuitSpades, "5", 5},
		Card{SuitSpades, "6", 6},
	}

	d.shuffle()

	if d[0].Rank == 2 && d[2].Rank == 4 && d[4].Rank == 6 {
		t.Errorf("Expected elements to be different from '2', '4' & '6', got %q, %q & %q",
			d[0].Rank, d[2].Rank, d[4].Rank)
	}
}

func TestDealFlop(t *testing.T) {
	d := Deck{
		Card{SuitSpades, "2", 2},
		Card{SuitSpades, "3", 3},
		Card{SuitSpades, "4", 4},
		Card{SuitSpades, "5", 5},
		Card{SuitSpades, "6", 6},
	}

	f, r := dealFlop(d)
	if f[0].Rank != 3 || f[1].Rank != 4 || f[2].Rank != 5 {
		t.Errorf("Expected cards '2', '3' and '4', got %q, %q and %q",
			f[0].Rank, f[1].Rank, f[2].Rank)
	}

	expLen := 3
	if len(f) != expLen {
		t.Errorf("Expected size of Flop deck %d, got %d",
			expLen, len(f))
	}

	expCard := Card{SuitSpades, "6", 6}
	if r[0].Rank != expCard.Rank || r[0].Suit != expCard.Suit || r[0].Value != expCard.Value {
		t.Errorf("Expected card %d of %v, got %d of %v",
			expCard.Rank, expCard.Suit, r[0].Rank, r[0].Suit)
	}
}

func TestDealTurn(t *testing.T) {
	d := Deck{
		Card{SuitSpades, "2", 2},
		Card{SuitSpades, "3", 3},
		Card{SuitSpades, "4", 4},
	}

	turn, r := dealTurn(d)
	expCard := Card{SuitSpades, "3", 3}
	if turn[0].Rank != expCard.Rank || turn[0].Suit != expCard.Suit || turn[0].Value != expCard.Value {
		t.Errorf("Expected card %d of %v, got %d of %v",
			expCard.Rank, expCard.Suit, turn[0].Rank, turn[0].Suit)
	}

	expLen := 1
	if len(turn) != expLen {
		t.Errorf("Expected size of Flop deck %d, got %d",
			expLen, len(turn))
	}

	expCard = Card{SuitSpades, "4", 4}
	if r[0].Rank != expCard.Rank || r[0].Suit != expCard.Suit || r[0].Value != expCard.Value {
		t.Errorf("Expected card %d of %v, got %d of %v",
			expCard.Rank, expCard.Suit, r[0].Rank, r[0].Suit)
	}
}

func TestDealRiver(t *testing.T) {
	d := Deck{
		Card{SuitSpades, "8", 8},
		Card{SuitSpades, "9", 9},
		Card{SuitSpades, "10", 10},
	}

	rvr, rmng := dealRiver(d)
	expCard := Card{SuitSpades, "9", 9}
	if rvr[0].Rank != expCard.Rank || rvr[0].Suit != expCard.Suit || rvr[0].Value != expCard.Value {
		t.Errorf("Expected card %d of %v, got %d of %v",
			expCard.Rank, expCard.Suit, rvr[0].Rank, rvr[0].Suit)
	}

	expLen := 1
	if len(rvr) != expLen {
		t.Errorf("Expected size of Flop deck %d, got %d",
			expLen, len(rvr))
	}

	expCard = Card{SuitSpades, "10", 10}
	if rmng[0].Rank != expCard.Rank || rmng[0].Suit != expCard.Suit || rmng[0].Value != expCard.Value {
		t.Errorf("Expected card %d of %v, got %d of %v",
			expCard.Rank, expCard.Suit, rmng[0].Rank, rmng[0].Suit)
	}
}

func TestCompareCardRank(t *testing.T) {
	d := Deck{
		Card{SuitSpades, "8", 8},
		Card{SuitSpades, "9", 9},
		Card{SuitSpades, "10", 10},
	}

	d.sortDeck()

	expected := Deck{
		Card{SuitSpades, "10", 10},
		Card{SuitSpades, "9", 9},
		Card{SuitSpades, "8", 8},
	}

	for i, card := range d {
		if card.Rank != expected[i].Rank {
			t.Errorf("Expected rank %d at position %d, but got %d",
				expected[i].Rank, i, card.Rank)
		}
	}
}
