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
	combinationId int
	playerHand    deck
}

func (d deck) evaluateHand() handResult {

	d.sortDeck()
	fmt.Println("Sorted D:", d)

	if rFlush := d.hasRoyalFlush(); rFlush.combinationId == RoyalFlushId {
		return rFlush
	}

	if sFlush := d.hasStraightFlush(); sFlush.combinationId == StraightFlushId {
		return sFlush
	}

	if four := d.hasFour(); four.combinationId == FourKindId {
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
		return three
	}

	if twoPair := d.hasTwoPair(); twoPair.combinationId == TwoPairId {
		return twoPair
	}

	if pair := d.hasPair(); pair.combinationId == SinglePairId {
		return pair
	}

	return handResult{combinationId: HighCardId, playerHand: d[:5]}
}

func (d deck) hasRoyalFlush() handResult {
	hr := handResult{combinationId: RoyalFlushId, playerHand: d[:5]}
	return hr
}

func (d deck) hasStraightFlush() handResult {
	hr := handResult{combinationId: StraightFlushId, playerHand: d[:5]}
	return hr
}

func (d deck) hasFour() handResult {
	cardCounter := make(map[string][]int)

	for idx, c := range d {
		if cardSlice, ok := cardCounter[c.value]; ok {
			if len(cardSlice) == 3 { // found four of a kind
				cardCounter[c.value] = append(cardSlice, idx)
				// // add the four cards to the players 5-card hand
				bestHand := addCombinationCards(d, cardCounter[c.value])
				bestHand = bestHand.addHighCards(d, []string{c.value})
				return handResult{combinationId: FourKindId, playerHand: bestHand}

			} else { // found second or third instance
				cardCounter[c.value] = append(cardSlice, idx)
			}
		} else { // found first instance
			cardCounter[c.value] = []int{idx}
		}
	}

	return handResult{combinationId: NotMatchId, playerHand: d[:5]}
}

func (d deck) hasFullHouse() handResult {
	hr := handResult{combinationId: FullHouseId, playerHand: d[:5]}
	return hr
}

func (d deck) hasFlush() handResult {
	hr := handResult{combinationId: FlushId, playerHand: d[:5]}
	return hr
}

func (d deck) hasStraight() handResult {
	hr := handResult{combinationId: StraightId, playerHand: d[:5]}
	return hr
}

func (d deck) hasThree() handResult {
	cardCounter := make(map[string][]int)
	bestHand := deck{} // the player's five best cards

	for idx, c := range d {
		if cardSlice, ok := cardCounter[c.value]; ok {
			if len(cardSlice) == 2 { // found three of a kind
				firstInstIdx := cardCounter[c.value][0]  // get 1st instance index
				secondInstIdx := cardCounter[c.value][1] // get 2nd instance index
				// add the three cards to the players 5-card hand
				bestHand = append(bestHand, d[firstInstIdx], d[secondInstIdx], d[idx])
				bestHand = bestHand.addHighCards(d, []string{c.value})
				return handResult{combinationId: ThreeKindId, playerHand: bestHand}

			} else { // found second instance
				cardCounter[c.value] = append(cardSlice, idx)
			}
		} else { // found first instance
			cardCounter[c.value] = []int{idx}
		}
	}

	return handResult{combinationId: NotMatchId, playerHand: d[:5]}
}

func (d deck) hasTwoPair() handResult {
	hr := handResult{combinationId: TwoPairId, playerHand: d[:5]}
	return hr
}

func (d deck) hasPair() handResult {
	uniques := make(map[string][]int)
	bestHand := deck{} // the player's five best cards

	for idx, c := range d {
		if _, ok := uniques[c.value]; ok {
			firstInstIdx := uniques[c.value][0]                    // get the 1st instance index
			bestHand = append(bestHand, c, d[firstInstIdx])        // append the 1st and 2nd instances
			bestHand = bestHand.addHighCards(d, []string{c.value}) // append the 3 highest remaining cards
			return handResult{combinationId: SinglePairId, playerHand: bestHand}
		}

		uniques[c.value] = []int{idx}
	}

	return handResult{combinationId: NotMatchId, playerHand: d[:5]}
}

func (d deck) getHighCard() card {
	return d[0]
}

func (fiveDeck deck) addHighCards(sevenDeck deck, valsToAvoid []string) deck {
	for i := 0; len(fiveDeck) < 5; i++ {
		// slices.Contains() method introduced in go 1.21
		if !slices.Contains(valsToAvoid, sevenDeck[i].value) {
			fiveDeck = append(fiveDeck, sevenDeck[i])
		}
	}
	return fiveDeck
}

// returns a deck with cards from a matching combination
// three of a kind example: 3 of clubs + 3 of hearts + 3 of spades
func addCombinationCards(d deck, idxSlice []int) deck {
	result := deck{}
	for _, idx := range idxSlice {
		result = append(result, d[idx])
	}

	return result
}
