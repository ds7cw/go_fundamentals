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

type handResult struct {
	combinationId     int
	playerHand        deck
	combinationValues []string
	// a slice of unique card values, forming a combination.
	// Full house of AAAKK holds the following values [A K]
}

func (d deck) evaluateHand() handResult {

	d.sortDeck()
	fmt.Println("Sorted D:", d)

	if rFlush := d.hasStraightFlush(); rFlush.combinationId == RoyalFlushId ||
		rFlush.combinationId == StraightFlushId {
		return rFlush
	}

	if four := d.hasFour(); four.combinationId == FourKindId {
		// append 1 highest remaining card
		four.playerHand.addHighCards(d, four.combinationValues)
		return four
	}

	if house := d.hasFullHouse(); house.combinationId == FullHouseId {
		return house
	}

	if flush := d.hasFlush(); flush.combinationId == FlushId {
		return flush
	}

	if straight := d.hasStraight(); straight.combinationId == StraightId {
		return straight
	}

	if three := d.hasThree(); three.combinationId == ThreeKindId {
		// append the 2 highest remaining cards
		three.playerHand.addHighCards(d, three.combinationValues)
		return three
	}

	if twoPair := d.hasTwoPair(); twoPair.combinationId == TwoPairId {
		twoPair.playerHand.addHighCards(d, twoPair.combinationValues)
		return twoPair
	}

	if pair := d.hasPair(); pair.combinationId == SinglePairId {
		// append the 3 highest remaining cards
		pair.playerHand.addHighCards(d, pair.combinationValues)
		return pair
	}

	return handResult{combinationId: HighCardId, playerHand: d[:5]}
}

// Checks for a royal flush card combination.
func (d deck) hasStraightFlush() handResult {
	// Group cards by suit
	suitGroups := map[string]deck{}
	for _, c := range d {
		suitGroups[c.suit] = append(suitGroups[c.suit], c)
	}

	// Check each suit group for royal flush or straight flush
	for _, suitHand := range suitGroups {
		if len(suitHand) >= 5 {
			straightResult := suitHand.hasStraight()

			// Royal flush if the lowest card is 10
			if straightResult.combinationId == StraightId && straightResult.playerHand[4].rank == 10 {
				straightResult.combinationId = RoyalFlushId

				return straightResult
			} else if straightResult.combinationId == StraightId {
				// Straigth flush if the 5 cards form a straight but no 10 as the lowest card
				straightResult.combinationId = StraightFlushId

				return straightResult
			}
		}
	}
	return handResult{combinationId: NotMatchId, playerHand: deck{}, combinationValues: []string{}}
}

// Checks for a four of a kind card combination.
func (d deck) hasFour() handResult {
	cardCounter := make(map[string][]int)

	for idx, c := range d {
		if cardSlice, ok := cardCounter[c.value]; ok {
			// add new instance idx to the relevant slice
			cardCounter[c.value] = append(cardSlice, idx)
			if len(cardCounter[c.value]) == 4 { // found four of a kind
				cardCounter[c.value] = append(cardSlice, idx)

				// add the four of a kind to the players 5-card hand
				bestHand := addCombinationCards(d, cardCounter[c.value])

				return handResult{
					combinationId:     FourKindId,
					playerHand:        bestHand,
					combinationValues: []string{c.value},
				}
			}
		} else { // found first instance
			cardCounter[c.value] = []int{idx}
		}
	}

	return handResult{combinationId: NotMatchId, playerHand: deck{}, combinationValues: []string{}}
}

// Checks for a full house card combination.
func (d deck) hasFullHouse() handResult {
	threeOfAKind := d.hasThree()

	if threeOfAKind.combinationId != ThreeKindId {
		return handResult{combinationId: NotMatchId, playerHand: deck{}, combinationValues: []string{}}
	}

	// remove the three of a kind cards from deck
	deckWithoutThreeKind := d.removeFromDeck(threeOfAKind.combinationValues)
	twoOfAKind := deckWithoutThreeKind.hasPair()

	// check remaining 4 cards for a two of a kind
	if twoOfAKind.combinationId != SinglePairId {
		return handResult{combinationId: NotMatchId, playerHand: deck{}, combinationValues: []string{}}
	}

	fullHouseDeck := slices.Concat(threeOfAKind.playerHand, twoOfAKind.playerHand)

	return handResult{
		combinationId:     FullHouseId,
		playerHand:        fullHouseDeck,
		combinationValues: []string{},
	}
}

// Checks for a flush card combination.
func (d deck) hasFlush() handResult {
	cardCounter := make(map[string][]int)

	for idx, c := range d {
		if cardSlice, ok := cardCounter[c.suit]; ok {
			// add new instance idx to the relevant slice
			cardCounter[c.suit] = append(cardSlice, idx)

			if len(cardCounter[c.suit]) == 5 {
				// 5 cards of the same suit found
				bestHand := addCombinationCards(d, cardCounter[c.suit])
				bestHand.sortDeck()

				return handResult{
					combinationId:     FlushId,
					playerHand:        bestHand,
					combinationValues: []string{},
				}
			}
		} else { // first encounter of current suit
			cardCounter[c.suit] = []int{idx}
		}
	}

	return handResult{combinationId: NotMatchId, playerHand: deck{}, combinationValues: []string{}}
}

func (d deck) hasStraight() handResult {
	// User a map to filter out duplicate values
	cardSet := map[int]card{}
	for _, c := range d {
		cardSet[c.rank] = c
	}

	// Create a deck of cards with unique values
	uniqueRanks := make(deck, 0, len(cardSet))
	for rank := range cardSet {
		uniqueRanks = append(uniqueRanks, cardSet[rank])
	}

	// Sort the cards in descending order
	uniqueRanks.sortDeck()

	// Check for a sequence of 5 consecutive ranks
	for i := 0; i <= len(uniqueRanks)-5; i++ {
		if uniqueRanks[i].rank-uniqueRanks[i+4].rank == 4 {
			return handResult{
				combinationId:     StraightId,
				playerHand:        uniqueRanks[i : i+5],
				combinationValues: []string{},
			}
		}
	}

	// Check for a low straight: A 2 3 4 5
	lowStraightResult := lookForLowStraight(cardSet)
	if len(lowStraightResult) == 5 {
		return handResult{
			combinationId:     StraightId,
			playerHand:        lowStraightResult,
			combinationValues: []string{},
		}
	}

	return handResult{combinationId: NotMatchId, playerHand: deck{}, combinationValues: []string{}}
}

// Checks for a three of a kind card combination.
func (d deck) hasThree() handResult {
	cardCounter := make(map[string][]int)

	for idx, c := range d {
		if cardSlice, ok := cardCounter[c.value]; ok {
			// add new instance idx to the relevant slice
			cardCounter[c.value] = append(cardSlice, idx)
			if len(cardCounter[c.value]) == 3 { // found three of a kind
				// add the three of a kind to the players 5-card hand
				bestHand := addCombinationCards(d, cardCounter[c.value])

				return handResult{
					combinationId:     ThreeKindId,
					playerHand:        bestHand,
					combinationValues: []string{c.value},
				}
			}

		} else { // found first instance
			cardCounter[c.value] = []int{idx}
		}
	}

	return handResult{combinationId: NotMatchId, playerHand: deck{}, combinationValues: []string{}}
}

// Checks for 2x two of a kind card combination.
func (d deck) hasTwoPair() handResult {
	// check the 7 cards for a two of a kind
	firstPair := d.hasPair()

	if firstPair.combinationId != SinglePairId {
		return handResult{combinationId: NotMatchId, playerHand: deck{}, combinationValues: []string{}}
	}

	// remove two of a kind cards from deck
	deckWithoutFirstPair := d.removeFromDeck(firstPair.combinationValues)

	// check remaining 5 cards for a two of a kind
	secondPair := deckWithoutFirstPair.hasPair()

	if secondPair.combinationId != SinglePairId {
		return handResult{combinationId: NotMatchId, playerHand: deck{}, combinationValues: []string{}}
	}

	twoPairsDeck := slices.Concat(firstPair.playerHand, secondPair.playerHand)
	valuesToAvoid := slices.Concat(firstPair.combinationValues, secondPair.combinationValues)

	if secondPair.playerHand[0].rank > firstPair.playerHand[0].rank {
		twoPairsDeck[0], twoPairsDeck[2] = twoPairsDeck[2], twoPairsDeck[0]
		twoPairsDeck[1], twoPairsDeck[3] = twoPairsDeck[3], twoPairsDeck[1]
		valuesToAvoid[0], valuesToAvoid[1] = valuesToAvoid[1], valuesToAvoid[0]
	}

	return handResult{
		combinationId:     TwoPairId,
		playerHand:        twoPairsDeck,
		combinationValues: valuesToAvoid,
	}
}

// Checks for a single pair card combination.
func (d deck) hasPair() handResult {
	cardCounter := make(map[string][]int)

	for idx, c := range d {
		if cardSlice, ok := cardCounter[c.value]; ok {
			// add second instance idx to the relevant slice
			cardCounter[c.value] = append(cardSlice, idx)
			bestHand := addCombinationCards(d, cardCounter[c.value])

			return handResult{
				combinationId:     SinglePairId,
				playerHand:        bestHand,
				combinationValues: []string{c.value},
			}
		}
		// found first instance
		cardCounter[c.value] = []int{idx}
	}

	return handResult{combinationId: NotMatchId, playerHand: deck{}, combinationValues: []string{}}
}

func (fiveDeck *deck) addHighCards(sevenDeck deck, valsToAvoid []string) {
	for i := 0; len(*fiveDeck) < 5; i++ {
		// slices.Contains() method introduced in go 1.21
		if !slices.Contains(valsToAvoid, sevenDeck[i].value) {
			*fiveDeck = append(*fiveDeck, sevenDeck[i])
		}
	}
}

// Returns a deck with cards from a matching combination.
// Three of a kind example: 3 of clubs + 3 of hearts + 3 of spades
func addCombinationCards(d deck, idxSlice []int) deck {
	result := deck{}
	for _, idx := range idxSlice {
		result = append(result, d[idx])
	}

	return result
}

func (d deck) removeFromDeck(valuesToRemove []string) deck {
	subDeck := deck{}

	for _, el := range d {
		if !slices.Contains(valuesToRemove, el.value) {
			// if current card value not among the values to remove
			subDeck = append(subDeck, el)
		}
	}

	return subDeck
}

// Checks whether a str:card map contains a given str key.
func contains(cardSet map[int]card, rank int) bool {
	_, exists := cardSet[rank]
	return exists
}

// Returns an Ace-low straight (A, 2, 3, 4, 5) if those cards are present.
// Otherwise the function returns an empty deck.
func lookForLowStraight(cardSet map[int]card) deck {
	if contains(cardSet, 14) && contains(cardSet, 5) &&
		contains(cardSet, 4) && contains(cardSet, 3) &&
		contains(cardSet, 2) {
		return deck{
			cardSet[5],
			cardSet[4],
			cardSet[3],
			cardSet[2],
			cardSet[14],
		}
	}

	return deck{}
}
