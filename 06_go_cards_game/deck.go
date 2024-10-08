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

type card struct {
	suit  string
	value string
	rank  int
}

// Create a new type of 'deck'
// which is a slice of card structs
type deck []card

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{SuitClubs, SuitDiamonds, SuitHearts, SuitSpades}
	cardValues := []string{
		"2", "3", "4", "5",
		"6", "7", "8", "9", "10",
		"Jack", "Queen", "King", "Ace",
	}

	i := 2

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			crd := card{suit: suit, value: value, rank: i}
			cards = append(cards, crd)
		}
		i += 1
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	var string_deck []string

	for _, el := range d {
		string_deck = append(string_deck, el.value+" of "+el.suit+" "+strconv.Itoa(el.rank))
	}
	return strings.Join([]string(string_deck), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option 1: Log the error and return a call to newDeck()
		// Option 2: Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	d := deck{}

	for _, el := range s {
		sub_el := strings.Split(el, " ")
		card_val := sub_el[0]
		card_suit := sub_el[2]
		card_rank, err := strconv.Atoi(sub_el[3])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		crd := card{suit: card_suit, value: card_val, rank: card_rank}
		d = append(d, crd)
	}
	return d
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn((len(d) - 1))
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func dealFlop(d deck) (deck, deck) {
	burn, lessBurn := deal(d, 1)
	flop, remaining := deal(lessBurn, 3)

	fmt.Println("Burn at Flop:", burn)
	fmt.Println("Flop Cards:", flop)

	return flop, remaining
}

func dealTurn(d deck) (deck, deck) {
	burn, lessBurn := deal(d, 1)
	turn, remaining := deal(lessBurn, 1)

	fmt.Println("Burn at Turn:", burn)
	fmt.Println("Turn Card:", turn)

	return turn, remaining
}

func dealRiver(d deck) (deck, deck) {
	burn, lessBurn := deal(d, 1)
	river, remaining := deal(lessBurn, 1)

	fmt.Println("Burn at River:", burn)
	fmt.Println("River Card:", river)

	return river, remaining
}

func compareCardRank(d deck) func(i, j int) bool {
	return func(i, j int) bool {
		return d[i].rank > d[j].rank
	}
}

func (d deck) sortDeck() {
	sort.Slice(d, compareCardRank(d))
}
