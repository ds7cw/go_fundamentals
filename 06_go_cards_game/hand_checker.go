package main

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
	handId       int
	handScore    int
	handHighCard card
}

func (d deck) evaluateHand() handResult {

	if rFlush := d.hasRoyalFlush(); rFlush.handId == 10 {
		return rFlush
	}

	if sFlush := d.hasStraightFlush(); sFlush.handId == 9 {
		return sFlush
	}

	if four := d.hasFour(); four.handId == 8 {
		return four
	}

	if house := d.hasFullHouse(); house.handId == 7 {
		return house
	}

	if flush := d.hasFlush(); flush.handId == 6 {
		return flush
	}

	if straight := d.hasStraight(); straight.handId == 5 {
		return straight
	}

	if three := d.hasThree(); three.handId == 4 {
		return three
	}

	if twoPair := d.hasTwoPair(); twoPair.handId == 3 {
		return twoPair
	}

	if pair := d.hasPair(); pair.handId == 2 {
		return pair
	}

	return handResult{handId: 1, handScore: 0, handHighCard: d.getHighCard()}
}

func (d deck) hasRoyalFlush() handResult {
	hr := handResult{handId: 10, handScore: 0, handHighCard: card{suit: SuitSpades, value: "Ace", rank: 13}}
	return hr
}

func (d deck) hasStraightFlush() handResult {
	hr := handResult{handId: 9, handScore: 0, handHighCard: card{suit: SuitSpades, value: "Ace", rank: 13}}
	return hr
}

func (d deck) hasFour() handResult {
	hr := handResult{handId: 8, handScore: 0, handHighCard: card{suit: SuitSpades, value: "Ace", rank: 13}}
	return hr
}

func (d deck) hasFullHouse() handResult {
	hr := handResult{handId: 7, handScore: 0, handHighCard: card{suit: SuitSpades, value: "Ace", rank: 13}}
	return hr
}

func (d deck) hasFlush() handResult {
	hr := handResult{handId: 6, handScore: 0, handHighCard: card{suit: SuitSpades, value: "Ace", rank: 13}}
	return hr
}

func (d deck) hasStraight() handResult {
	hr := handResult{handId: 5, handScore: 0, handHighCard: card{suit: SuitSpades, value: "Ace", rank: 13}}
	return hr
}

func (d deck) hasThree() handResult {
	hr := handResult{handId: 4, handScore: 0, handHighCard: card{suit: SuitSpades, value: "Ace", rank: 13}}
	return hr
}

func (d deck) hasTwoPair() handResult {
	hr := handResult{handId: 3, handScore: 0, handHighCard: card{suit: SuitSpades, value: "Ace", rank: 13}}
	return hr
}

func (d deck) hasPair() handResult {
	hr := handResult{handId: 2, handScore: 0, handHighCard: card{suit: SuitSpades, value: "Ace", rank: 13}}
	return hr
}

func (d deck) getHighCard() card {
	hc := d[0]
	last_idx := len(d) - 1

	for i := 1; i <= last_idx; i++ {
		if d[i].rank > hc.rank {
			hc = d[i]
		}
	}
	return hc
}
