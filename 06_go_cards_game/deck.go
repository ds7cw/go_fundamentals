package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{
		"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
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
	return strings.Join([]string(d), ",")
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
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn((len(d) - 1))
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func getPlayerCount() int {
	return rand.Intn(7) + 2
}

func createPlayers(n int) map[string]deck {
	playerMap := map[string]deck{}

	for i := 1; i <= n; i++ {
		currentPlayer := "Player" + strconv.Itoa(i)
		playerMap[currentPlayer] = deck{}
	}

	return playerMap
}

func dealToPlayers(m map[string]deck, d deck) (map[string]deck, deck) {
	for key, _ := range m {
		m[key], d = deal(d, 2)
	}

	return m, d
}

func printPlayersHands(m map[string]deck) {
	for key, val := range m {
		fmt.Println(key, val)
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
