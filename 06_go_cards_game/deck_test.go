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
	fc := card{SuitClubs, "2", 2}
	if d[0].suit != fc.suit || d[0].rank != fc.rank || d[0].value != fc.value {
		t.Errorf("Expected first %v of %v, but got %v of %v",
			fc.value, fc.suit, d[0].value, d[0].suit)
	}

	lc := card{SuitSpades, "Ace", 14}
	if d[51].suit != lc.suit || d[51].rank != lc.rank || d[51].value != lc.value {
		t.Errorf("Expected first %v of %v, but got %v of %v",
			lc.value, lc.suit, d[51].value, d[51].suit)
	}
}

func TestPrintDeck(t *testing.T) {
	d := deck{card{SuitSpades, "Jack", 11}}

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
		t.Errorf("deck.print() = %q; want %q", got, want)
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

	fc := card{SuitClubs, "2", 2}
	if hand[0] != fc {
		t.Errorf("Expected first card %v of %v, but got %v of %v",
			fc.value, fc.suit, hand[0].value, hand[0].suit)
	}

	lc := card{SuitSpades, "Ace", 14}
	if remaining[len(remaining)-1] != lc {
		t.Errorf("Expected last card %v of %v, but got %v of %v",
			fc.value, fc.suit, remaining[25].value, remaining[25].suit)
	}
}

func TestShuffle(t *testing.T) {
	d := deck{
		card{SuitSpades, "2", 2},
		card{SuitSpades, "3", 3},
		card{SuitSpades, "4", 4},
		card{SuitSpades, "5", 5},
		card{SuitSpades, "6", 6},
	}

	d.shuffle()

	if d[0].rank == 2 && d[2].rank == 4 && d[4].rank == 6 {
		t.Errorf("Expected elements to be different from '2', '4' & '6', got %q, %q & %q",
			d[0].rank, d[2].rank, d[4].rank)
	}
}

func TestDealFlop(t *testing.T) {
	d := deck{
		card{SuitSpades, "2", 2},
		card{SuitSpades, "3", 3},
		card{SuitSpades, "4", 4},
		card{SuitSpades, "5", 5},
		card{SuitSpades, "6", 6},
	}

	f, r := dealFlop(d)
	if f[0].rank != 3 || f[1].rank != 4 || f[2].rank != 5 {
		t.Errorf("Expected cards '2', '3' and '4', got %q, %q and %q",
			f[0].rank, f[1].rank, f[2].rank)
	}

	expLen := 3
	if len(f) != expLen {
		t.Errorf("Expected size of Flop deck %d, got %d",
			expLen, len(f))
	}

	expCard := card{SuitSpades, "6", 6}
	if r[0].rank != expCard.rank || r[0].suit != expCard.suit || r[0].value != expCard.value {
		t.Errorf("Expected card %d of %v, got %d of %v",
			expCard.rank, expCard.suit, r[0].rank, r[0].suit)
	}
}

func TestDealTurn(t *testing.T) {
	d := deck{
		card{SuitSpades, "2", 2},
		card{SuitSpades, "3", 3},
		card{SuitSpades, "4", 4},
	}

	turn, r := dealTurn(d)
	expCard := card{SuitSpades, "3", 3}
	if turn[0].rank != expCard.rank || turn[0].suit != expCard.suit || turn[0].value != expCard.value {
		t.Errorf("Expected card %d of %v, got %d of %v",
			expCard.rank, expCard.suit, turn[0].rank, turn[0].suit)
	}

	expLen := 1
	if len(turn) != expLen {
		t.Errorf("Expected size of Flop deck %d, got %d",
			expLen, len(turn))
	}

	expCard = card{SuitSpades, "4", 4}
	if r[0].rank != expCard.rank || r[0].suit != expCard.suit || r[0].value != expCard.value {
		t.Errorf("Expected card %d of %v, got %d of %v",
			expCard.rank, expCard.suit, r[0].rank, r[0].suit)
	}
}

func TestDealRiver(t *testing.T) {
	d := deck{
		card{SuitSpades, "8", 8},
		card{SuitSpades, "9", 9},
		card{SuitSpades, "10", 10},
	}

	rvr, rmng := dealRiver(d)
	expCard := card{SuitSpades, "9", 9}
	if rvr[0].rank != expCard.rank || rvr[0].suit != expCard.suit || rvr[0].value != expCard.value {
		t.Errorf("Expected card %d of %v, got %d of %v",
			expCard.rank, expCard.suit, rvr[0].rank, rvr[0].suit)
	}

	expLen := 1
	if len(rvr) != expLen {
		t.Errorf("Expected size of Flop deck %d, got %d",
			expLen, len(rvr))
	}

	expCard = card{SuitSpades, "10", 10}
	if rmng[0].rank != expCard.rank || rmng[0].suit != expCard.suit || rmng[0].value != expCard.value {
		t.Errorf("Expected card %d of %v, got %d of %v",
			expCard.rank, expCard.suit, rmng[0].rank, rmng[0].suit)
	}
}

func TestCompareCardRank(t *testing.T) {
	d := deck{
		card{SuitSpades, "8", 8},
		card{SuitSpades, "9", 9},
		card{SuitSpades, "10", 10},
	}

	d.sortDeck()

	expected := deck{
		card{SuitSpades, "10", 10},
		card{SuitSpades, "9", 9},
		card{SuitSpades, "8", 8},
	}

	for i, card := range d {
		if card.rank != expected[i].rank {
			t.Errorf("Expected rank %d at position %d, but got %d",
				expected[i].rank, i, card.rank)
		}
	}
}
