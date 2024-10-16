package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const SuitSpades string = "Spades"
const SuitDiamonds string = "Diamonds"
const SuitHearts string = "Hearts"
const SuitClubs string = "Clubs"

type Card struct {
	Suit  string
	Value string
	Rank  int
}

// Create a new type of 'deck'
// which is a slice of card structs
type Deck []Card

func newDeck() Deck {
	cards := Deck{}
	cardSuits := []string{SuitClubs, SuitDiamonds, SuitHearts, SuitSpades}
	cardValues := []string{
		"2", "3", "4", "5",
		"6", "7", "8", "9", "10",
		"Jack", "Queen", "King", "Ace",
	}

	i := 2

	for _, Value := range cardValues {
		for _, Suit := range cardSuits {
			crd := Card{Suit: Suit, Value: Value, Rank: i}
			cards = append(cards, crd)
		}
		i += 1
	}

	return cards

}

func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) toString() string {
	var string_deck []string

	for _, el := range d {
		string_deck = append(string_deck, el.Value+" of "+el.Suit+" "+strconv.Itoa(el.Rank))
	}
	return strings.Join([]string(string_deck), ",")
}

func (d Deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) Deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option 1: Log the error and return a call to newDeck()
		// Option 2: Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	d := Deck{}

	for _, el := range s {
		sub_el := strings.Split(el, " ")
		card_val := sub_el[0]
		card_Suit := sub_el[2]
		card_Rank, err := strconv.Atoi(sub_el[3])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		crd := Card{Suit: card_Suit, Value: card_val, Rank: card_Rank}
		d = append(d, crd)
	}
	return d
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn((len(d) - 1))
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func dealFlop(d Deck) (Deck, Deck) {
	burn, lessBurn := deal(d, 1)
	flop, remaining := deal(lessBurn, 3)

	fmt.Println("Burn at Flop:", burn)
	fmt.Println("Flop Cards:", flop)

	return flop, remaining
}

func dealTurn(d Deck) (Deck, Deck) {
	burn, lessBurn := deal(d, 1)
	turn, remaining := deal(lessBurn, 1)

	fmt.Println("Burn at Turn:", burn)
	fmt.Println("Turn Card:", turn)

	return turn, remaining
}

func dealRiver(d Deck) (Deck, Deck) {
	burn, lessBurn := deal(d, 1)
	river, remaining := deal(lessBurn, 1)

	fmt.Println("Burn at River:", burn)
	fmt.Println("River Card:", river)

	return river, remaining
}

func compareCardRank(d Deck) func(i, j int) bool {
	return func(i, j int) bool {
		return d[i].Rank > d[j].Rank
	}
}

func (d Deck) sortDeck() {
	sort.Slice(d, compareCardRank(d))
}
