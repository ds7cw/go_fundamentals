package main

import (
	"fmt"
	"slices"
)

const RoyalFlushId int = 10   // As Ks Qs Js 10s
const StraightFlushId int = 9 // Jc 10c 9c 8c 7c
const FourKindId int = 8      // 9 9 9 9 [K]
const FullHouseId int = 7     // A A A 3 3
const FlushId int = 6         // Kc 10c 8c 7c 5c
const StraightId int = 5      // 10h 9c 8d 7c 6h
const ThreeKindId int = 4     // 7 7 7 [Q 3]
const TwoPairId int = 3       // J J 5 5 [7]
const SinglePairId int = 2    // A A [K J 7]
const HighCardId int = 1      // K [8 Q 2 7]
const NotMatchId int = 0      // Ah 4s 9d Qc 6h

type HandResult struct {
	CombinationId     int
	PlayerHand        Deck
	CombinationValues []string
	// a slice of unique card values, forming a combination.
	// Full house of AAAKK holds the following values [A K]
}

func (d Deck) evaluateHand() HandResult {

	d.sortDeck()
	fmt.Println("Sorted D:", d)

	if rFlush := d.hasStraightFlush(); rFlush.CombinationId == RoyalFlushId ||
		rFlush.CombinationId == StraightFlushId {
		return rFlush
	}

	if four := d.hasFour(); four.CombinationId == FourKindId {
		// append 1 highest remaining card
		four.PlayerHand.addHighCards(d, four.CombinationValues)
		return four
	}

	if house := d.hasFullHouse(); house.CombinationId == FullHouseId {
		return house
	}

	if flush := d.hasFlush(); flush.CombinationId == FlushId {
		return flush
	}

	if straight := d.hasStraight(); straight.CombinationId == StraightId {
		return straight
	}

	if three := d.hasThree(); three.CombinationId == ThreeKindId {
		// append the 2 highest remaining cards
		three.PlayerHand.addHighCards(d, three.CombinationValues)
		return three
	}

	if twoPair := d.hasTwoPair(); twoPair.CombinationId == TwoPairId {
		twoPair.PlayerHand.addHighCards(d, twoPair.CombinationValues)
		return twoPair
	}

	if pair := d.hasPair(); pair.CombinationId == SinglePairId {
		// append the 3 highest remaining cards
		pair.PlayerHand.addHighCards(d, pair.CombinationValues)
		return pair
	}

	return HandResult{CombinationId: HighCardId, PlayerHand: d[:5]}
}

// Checks for a royal flush card combination.
func (d Deck) hasStraightFlush() HandResult {
	// Group cards by suit
	suitGroups := map[string]Deck{}
	for _, c := range d {
		suitGroups[c.Suit] = append(suitGroups[c.Suit], c)
	}

	// Check each suit group for royal flush or straight flush
	for _, suitHand := range suitGroups {
		if len(suitHand) >= 5 {
			straightResult := suitHand.hasStraight()

			// Royal flush if the lowest card is 10
			if straightResult.CombinationId == StraightId && straightResult.PlayerHand[4].Rank == 10 {
				straightResult.CombinationId = RoyalFlushId

				return straightResult
			} else if straightResult.CombinationId == StraightId {
				// Straigth flush if the 5 cards form a straight but no 10 as the lowest card
				straightResult.CombinationId = StraightFlushId

				return straightResult
			}
		}
	}
	return HandResult{CombinationId: NotMatchId, PlayerHand: Deck{}, CombinationValues: []string{}}
}

// Checks for a four of a kind card combination.
func (d Deck) hasFour() HandResult {
	cardCounter := make(map[string][]int)

	for idx, c := range d {
		if cardSlice, ok := cardCounter[c.Value]; ok {
			// add new instance idx to the relevant slice
			cardCounter[c.Value] = append(cardSlice, idx)
			if len(cardCounter[c.Value]) == 4 { // found four of a kind
				cardCounter[c.Value] = append(cardSlice, idx)

				// add the four of a kind to the players 5-card hand
				bestHand := addCombinationCards(d, cardCounter[c.Value])

				return HandResult{
					CombinationId:     FourKindId,
					PlayerHand:        bestHand,
					CombinationValues: []string{c.Value},
				}
			}
		} else { // found first instance
			cardCounter[c.Value] = []int{idx}
		}
	}

	return HandResult{CombinationId: NotMatchId, PlayerHand: Deck{}, CombinationValues: []string{}}
}

// Checks for a full house card combination.
func (d Deck) hasFullHouse() HandResult {
	threeOfAKind := d.hasThree()

	if threeOfAKind.CombinationId != ThreeKindId {
		return HandResult{CombinationId: NotMatchId, PlayerHand: Deck{}, CombinationValues: []string{}}
	}

	// remove the three of a kind cards from deck
	deckWithoutThreeKind := d.removeFromDeck(threeOfAKind.CombinationValues)
	twoOfAKind := deckWithoutThreeKind.hasPair()

	// check remaining 4 cards for a two of a kind
	if twoOfAKind.CombinationId != SinglePairId {
		return HandResult{CombinationId: NotMatchId, PlayerHand: Deck{}, CombinationValues: []string{}}
	}

	fullHouseDeck := slices.Concat(threeOfAKind.PlayerHand, twoOfAKind.PlayerHand)

	return HandResult{
		CombinationId:     FullHouseId,
		PlayerHand:        fullHouseDeck,
		CombinationValues: []string{},
	}
}

// Checks for a flush card combination.
func (d Deck) hasFlush() HandResult {
	cardCounter := make(map[string][]int)

	for idx, c := range d {
		if cardSlice, ok := cardCounter[c.Suit]; ok {
			// add new instance idx to the relevant slice
			cardCounter[c.Suit] = append(cardSlice, idx)

			if len(cardCounter[c.Suit]) == 5 {
				// 5 cards of the same suit found
				bestHand := addCombinationCards(d, cardCounter[c.Suit])
				bestHand.sortDeck()

				return HandResult{
					CombinationId:     FlushId,
					PlayerHand:        bestHand,
					CombinationValues: []string{},
				}
			}
		} else { // first encounter of current suit
			cardCounter[c.Suit] = []int{idx}
		}
	}

	return HandResult{CombinationId: NotMatchId, PlayerHand: Deck{}, CombinationValues: []string{}}
}

func (d Deck) hasStraight() HandResult {
	// User a map to filter out duplicate values
	cardSet := map[int]Card{}
	for _, c := range d {
		cardSet[c.Rank] = c
	}

	// Create a deck of cards with unique values
	uniqueRanks := make(Deck, 0, len(cardSet))
	for rank := range cardSet {
		uniqueRanks = append(uniqueRanks, cardSet[rank])
	}

	// Sort the cards in descending order
	uniqueRanks.sortDeck()

	// Check for a sequence of 5 consecutive ranks
	for i := 0; i <= len(uniqueRanks)-5; i++ {
		if uniqueRanks[i].Rank-uniqueRanks[i+4].Rank == 4 {
			return HandResult{
				CombinationId:     StraightId,
				PlayerHand:        uniqueRanks[i : i+5],
				CombinationValues: []string{},
			}
		}
	}

	// Check for a low straight: A 2 3 4 5
	lowStraightResult := lookForLowStraight(cardSet)
	if len(lowStraightResult) == 5 {
		return HandResult{
			CombinationId:     StraightId,
			PlayerHand:        lowStraightResult,
			CombinationValues: []string{},
		}
	}

	return HandResult{CombinationId: NotMatchId, PlayerHand: Deck{}, CombinationValues: []string{}}
}

// Checks for a three of a kind card combination.
func (d Deck) hasThree() HandResult {
	cardCounter := make(map[string][]int)

	for idx, c := range d {
		if cardSlice, ok := cardCounter[c.Value]; ok {
			// add new instance idx to the relevant slice
			cardCounter[c.Value] = append(cardSlice, idx)
			if len(cardCounter[c.Value]) == 3 { // found three of a kind
				// add the three of a kind to the players 5-card hand
				bestHand := addCombinationCards(d, cardCounter[c.Value])

				return HandResult{
					CombinationId:     ThreeKindId,
					PlayerHand:        bestHand,
					CombinationValues: []string{c.Value},
				}
			}

		} else { // found first instance
			cardCounter[c.Value] = []int{idx}
		}
	}

	return HandResult{CombinationId: NotMatchId, PlayerHand: Deck{}, CombinationValues: []string{}}
}

// Checks for 2x two of a kind card combination.
func (d Deck) hasTwoPair() HandResult {
	// check the 7 cards for a two of a kind
	firstPair := d.hasPair()

	if firstPair.CombinationId != SinglePairId {
		return HandResult{CombinationId: NotMatchId, PlayerHand: Deck{}, CombinationValues: []string{}}
	}

	// remove two of a kind cards from deck
	deckWithoutFirstPair := d.removeFromDeck(firstPair.CombinationValues)

	// check remaining 5 cards for a two of a kind
	secondPair := deckWithoutFirstPair.hasPair()

	if secondPair.CombinationId != SinglePairId {
		return HandResult{CombinationId: NotMatchId, PlayerHand: Deck{}, CombinationValues: []string{}}
	}

	twoPairsDeck := slices.Concat(firstPair.PlayerHand, secondPair.PlayerHand)
	valuesToAvoid := slices.Concat(firstPair.CombinationValues, secondPair.CombinationValues)

	if secondPair.PlayerHand[0].Rank > firstPair.PlayerHand[0].Rank {
		twoPairsDeck[0], twoPairsDeck[2] = twoPairsDeck[2], twoPairsDeck[0]
		twoPairsDeck[1], twoPairsDeck[3] = twoPairsDeck[3], twoPairsDeck[1]
		valuesToAvoid[0], valuesToAvoid[1] = valuesToAvoid[1], valuesToAvoid[0]
	}

	return HandResult{
		CombinationId:     TwoPairId,
		PlayerHand:        twoPairsDeck,
		CombinationValues: valuesToAvoid,
	}
}

// Checks for a single pair card combination.
func (d Deck) hasPair() HandResult {
	cardCounter := make(map[string][]int)

	for idx, c := range d {
		if cardSlice, ok := cardCounter[c.Value]; ok {
			// add second instance idx to the relevant slice
			cardCounter[c.Value] = append(cardSlice, idx)
			bestHand := addCombinationCards(d, cardCounter[c.Value])

			return HandResult{
				CombinationId:     SinglePairId,
				PlayerHand:        bestHand,
				CombinationValues: []string{c.Value},
			}
		}
		// found first instance
		cardCounter[c.Value] = []int{idx}
	}

	return HandResult{CombinationId: NotMatchId, PlayerHand: Deck{}, CombinationValues: []string{}}
}

func (fiveDeck *Deck) addHighCards(sevenDeck Deck, valsToAvoid []string) {
	for i := 0; len(*fiveDeck) < 5; i++ {
		// slices.Contains() method introduced in go 1.21
		if !slices.Contains(valsToAvoid, sevenDeck[i].Value) {
			*fiveDeck = append(*fiveDeck, sevenDeck[i])
		}
	}
}

// Returns a deck with cards from a matching combination.
// Three of a kind example: 3 of clubs + 3 of hearts + 3 of spades
func addCombinationCards(d Deck, idxSlice []int) Deck {
	result := Deck{}
	for _, idx := range idxSlice {
		result = append(result, d[idx])
	}

	return result
}

func (d Deck) removeFromDeck(valuesToRemove []string) Deck {
	subDeck := Deck{}

	for _, el := range d {
		if !slices.Contains(valuesToRemove, el.Value) {
			// if current card value not among the values to remove
			subDeck = append(subDeck, el)
		}
	}

	return subDeck
}

// Checks whether a str:card map contains a given str key.
func contains(cardSet map[int]Card, rank int) bool {
	_, exists := cardSet[rank]
	return exists
}

// Returns an Ace-low straight (A, 2, 3, 4, 5) if those cards are present.
// Otherwise the function returns an empty deck.
func lookForLowStraight(cardSet map[int]Card) Deck {
	if contains(cardSet, 14) && contains(cardSet, 5) &&
		contains(cardSet, 4) && contains(cardSet, 3) &&
		contains(cardSet, 2) {
		return Deck{
			cardSet[5],
			cardSet[4],
			cardSet[3],
			cardSet[2],
			cardSet[14],
		}
	}

	return Deck{}
}
