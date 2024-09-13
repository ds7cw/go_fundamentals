package main

import (
	"fmt"
	"slices"
)

// 10 royal flush: 		As Ks Qs Js 10s
// 9 straight flush: 	Jc 10c 9c 8c 7c
// 8 four of a kind:	9 9 9 9 [K]
// 7 full house:		A A A 3 3
// 6 flush:				Kc 10c 8c 7c 5c
// 5 straight:			10h 9c 8d 7c 6h
// 4 three of a kind:	7 7 7 [Q 3]
// 3 two pair:			J J 5 5 [7]
// 2 pair:				A A [K J 7]
// 1 high card:			K [8 Q 2 7]

type handResult struct {
	combinationId int
	playerHand    deck
}

func (d deck) evaluateHand() handResult {

	d.sortDeck()
	fmt.Println("Sorted D:", d)

	if rFlush := d.hasRoyalFlush(); rFlush.combinationId == 10 {
		return rFlush
	}

	if sFlush := d.hasStraightFlush(); sFlush.combinationId == 9 {
		return sFlush
	}

	if four := d.hasFour(); four.combinationId == 8 {
		return four
	}

	if house := d.hasFullHouse(); house.combinationId == 7 {
		return house
	}

	if flush := d.hasFlush(); flush.combinationId == 6 {
		return flush
	}

	if straight := d.hasStraight(); straight.combinationId == 5 {
		return straight
	}

	if three := d.hasThree(); three.combinationId == 4 {
		return three
	}

	if twoPair := d.hasTwoPair(); twoPair.combinationId == 3 {
		return twoPair
	}

	if pair := d.hasPair(); pair.combinationId == 2 {
		return pair
	}

	return handResult{combinationId: 1, playerHand: d[:5]}
}

func (d deck) hasRoyalFlush() handResult {
	hr := handResult{combinationId: 10, playerHand: d[:5]}
	return hr
}

func (d deck) hasStraightFlush() handResult {
	hr := handResult{combinationId: 9, playerHand: d[:5]}
	return hr
}

func (d deck) hasFour() handResult {
	hr := handResult{combinationId: 8, playerHand: d[:5]}
	return hr
}

func (d deck) hasFullHouse() handResult {
	hr := handResult{combinationId: 7, playerHand: d[:5]}
	return hr
}

func (d deck) hasFlush() handResult {
	hr := handResult{combinationId: 6, playerHand: d[:5]}
	return hr
}

func (d deck) hasStraight() handResult {
	hr := handResult{combinationId: 5, playerHand: d[:5]}
	return hr
}

func (d deck) hasThree() handResult {
	hr := handResult{combinationId: 4, playerHand: d[:5]}
	return hr
}

func (d deck) hasTwoPair() handResult {
	hr := handResult{combinationId: 3, playerHand: d[:5]}
	return hr
}

func (d deck) hasPair() handResult {
	uniques := make(map[string][]int)
	bestHand := deck{} // the player's five best cards

	for idx, c := range d {
		if _, ok := uniques[c.value]; ok { // fix this
			firstInstIdx := uniques[c.value][1]                    // get the 1st instance index
			bestHand = append(bestHand, c, d[firstInstIdx])        // append the 1st and 2nd instances
			bestHand = bestHand.addHighCards(d, []string{c.value}) // append the 3 highest remaining cards
			return handResult{combinationId: 2, playerHand: bestHand}
		}

		uniques[c.value] = []int{1, idx}
	}

	return handResult{combinationId: 0, playerHand: d[:5]}
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
