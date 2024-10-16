package main

import "testing"

func TestHasPair(t *testing.T) {
	d := Deck{
		Card{SuitClubs, "3", 3},
		Card{SuitDiamonds, "Jack", 11},
		Card{SuitClubs, "10", 10},
		Card{SuitHearts, "6", 6},
		Card{SuitDiamonds, "10", 10},
		Card{SuitSpades, "King", 13},
		Card{SuitSpades, "8", 8},
	}

	res := d.hasPair()
	exp := HandResult{
		CombinationId: SinglePairId,
		PlayerHand: Deck{
			Card{SuitClubs, "10", 10},
			Card{SuitDiamonds, "10", 10},
			Card{SuitSpades, "King", 13},
			Card{SuitDiamonds, "Jack", 11},
			Card{SuitSpades, "8", 8},
		},
		CombinationValues: []string{"10"},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 2)
	helperCombinationValues(res, exp, t)

	d2 := Deck{
		Card{SuitClubs, "3", 3},
		Card{SuitDiamonds, "Jack", 11},
		Card{SuitClubs, "5", 5},
		Card{SuitHearts, "6", 6},
		Card{SuitDiamonds, "10", 10},
		Card{SuitSpades, "King", 13},
		Card{SuitSpades, "8", 8},
	}
	res2 := d2.hasPair()
	exp2 := NotMatchId

	if res2.CombinationId != exp2 {
		t.Errorf("Expected combinationId %d, got %d",
			exp2, res2.CombinationId)
	}

	res3 := d.evaluateHand()
	helperComboId(res3, exp, t)
	helperPlayerHand(res3, exp, t, 5)
	helperCombinationValues(res3, exp, t)
}

func TestHasTwoPair(t *testing.T) {
	d := Deck{
		Card{SuitClubs, "3", 3},
		Card{SuitDiamonds, "Jack", 11},
		Card{SuitClubs, "10", 10},
		Card{SuitHearts, "6", 6},
		Card{SuitDiamonds, "King", 13},
		Card{SuitSpades, "3", 3},
		Card{SuitHearts, "10", 10},
	}

	res := d.hasTwoPair()
	exp := HandResult{
		CombinationId: TwoPairId,
		PlayerHand: Deck{
			Card{SuitClubs, "10", 10},
			Card{SuitHearts, "10", 10},
			Card{SuitClubs, "3", 3},
			Card{SuitSpades, "3", 3},
		},
		CombinationValues: []string{"10", "3"},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 4)
	helperCombinationValues(res, exp, t)

	d[2] = Card{SuitHearts, "7", 7}
	res2 := d.hasTwoPair()
	exp2 := NotMatchId

	if res2.CombinationId != exp2 {
		t.Errorf("Expected combinationId %d, got %d",
			exp2, res2.CombinationId)
	}

	d[0] = Card{SuitHearts, "4", 4}
	res3 := d.hasTwoPair()
	exp3 := NotMatchId

	if res3.CombinationId != exp3 {
		t.Errorf("Expected combinationId %d, got %d",
			exp3, res3.CombinationId)
	}

	d[0] = Card{SuitClubs, "3", 3}
	d[2] = Card{SuitClubs, "10", 10}

	res4 := d.evaluateHand()
	exp.PlayerHand = append(exp.PlayerHand,
		Card{SuitDiamonds, "King", 13})
	helperComboId(res4, exp, t)
	helperPlayerHand(res4, exp, t, 5)
	helperCombinationValues(res4, exp, t)
}

func TestHasThree(t *testing.T) {
	d := Deck{
		Card{SuitClubs, "3", 3},
		Card{SuitDiamonds, "Jack", 11},
		Card{SuitClubs, "10", 10},
		Card{SuitHearts, "6", 6},
		Card{SuitDiamonds, "10", 10},
		Card{SuitSpades, "King", 13},
		Card{SuitHearts, "10", 10},
	}

	res := d.hasThree()
	exp := HandResult{
		CombinationId: ThreeKindId,
		PlayerHand: Deck{
			Card{SuitClubs, "10", 10},
			Card{SuitDiamonds, "10", 10},
			Card{SuitHearts, "10", 10},
		},
		CombinationValues: []string{"10"},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 3)
	helperCombinationValues(res, exp, t)

	d2 := Deck{
		Card{SuitClubs, "3", 3},
		Card{SuitDiamonds, "Jack", 11},
		Card{SuitClubs, "5", 5},
		Card{SuitHearts, "6", 6},
		Card{SuitDiamonds, "10", 10},
		Card{SuitSpades, "King", 13},
		Card{SuitSpades, "8", 8},
	}
	res2 := d2.hasThree()
	exp2 := NotMatchId

	if res2.CombinationId != exp2 {
		t.Errorf("Expected combinationId %d, got %d",
			exp2, res2.CombinationId)
	}

	res3 := d.evaluateHand()
	exp.PlayerHand = append(exp.PlayerHand,
		Card{SuitSpades, "King", 13},
		Card{SuitDiamonds, "Jack", 11})
	helperComboId(res3, exp, t)
	helperPlayerHand(res3, exp, t, 5)
	helperCombinationValues(res3, exp, t)
}

func TestHasStraight(t *testing.T) {
	d := Deck{
		Card{SuitClubs, "3", 3},
		Card{SuitDiamonds, "Jack", 11},
		Card{SuitClubs, "10", 10},
		Card{SuitHearts, "6", 6},
		Card{SuitDiamonds, "4", 4},
		Card{SuitSpades, "5", 5},
		Card{SuitHearts, "7", 7},
	}

	res := d.hasStraight()
	exp := HandResult{
		CombinationId: StraightId,
		PlayerHand: Deck{
			Card{SuitHearts, "7", 7},
			Card{SuitHearts, "6", 6},
			Card{SuitSpades, "5", 5},
			Card{SuitDiamonds, "4", 4},
			Card{SuitClubs, "3", 3},
		},
		CombinationValues: []string{},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 5)
	helperCombinationValues(res, exp, t)

	d[0] = Card{SuitSpades, "Ace", 14}
	res2 := d.hasStraight()
	exp2 := NotMatchId

	if res2.CombinationId != exp2 {
		t.Errorf("Expected combinationId %d, got %d",
			exp2, res2.CombinationId)
	}

	// Low straight A 2 3 4 5
	d = Deck{
		Card{SuitDiamonds, "10", 10},
		Card{SuitHearts, "5", 5},
		Card{SuitDiamonds, "Ace", 14},
		Card{SuitClubs, "4", 4},
		Card{SuitSpades, "Queen", 12},
		Card{SuitHearts, "2", 2},
		Card{SuitClubs, "3", 3},
	}

	res3 := d.hasStraight()
	exp3 := HandResult{
		CombinationId: StraightId,
		PlayerHand: Deck{
			Card{SuitHearts, "5", 5},
			Card{SuitClubs, "4", 4},
			Card{SuitClubs, "3", 3},
			Card{SuitHearts, "2", 2},
			Card{SuitDiamonds, "Ace", 14},
		},
		CombinationValues: []string{},
	}

	helperComboId(res3, exp3, t)
	helperPlayerHand(res3, exp3, t, 5)
	helperCombinationValues(res3, exp3, t)

	res4 := d.evaluateHand()
	helperComboId(res4, exp3, t)
	helperPlayerHand(res4, exp3, t, 5)
	helperCombinationValues(res4, exp3, t)
}

func TestHasFlush(t *testing.T) {
	d := Deck{
		Card{SuitHearts, "3", 3},
		Card{SuitHearts, "Jack", 11},
		Card{SuitHearts, "10", 10},
		Card{SuitDiamonds, "6", 6},
		Card{SuitHearts, "4", 4},
		Card{SuitSpades, "5", 5},
		Card{SuitHearts, "7", 7},
	}

	res := d.hasFlush()
	exp := HandResult{
		CombinationId: FlushId,
		PlayerHand: Deck{
			Card{SuitHearts, "Jack", 11},
			Card{SuitHearts, "10", 10},
			Card{SuitHearts, "7", 7},
			Card{SuitHearts, "4", 4},
			Card{SuitHearts, "3", 3},
		},
		CombinationValues: []string{},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 5)
	helperCombinationValues(res, exp, t)

	res2 := d.evaluateHand()
	helperComboId(res2, exp, t)
	helperPlayerHand(res2, exp, t, 5)
	helperCombinationValues(res2, exp, t)

	d[0] = Card{SuitSpades, "Ace", 14}
	res3 := d.hasFlush()
	exp3 := NotMatchId

	if res3.CombinationId != exp3 {
		t.Errorf("Expected combinationId %d, got %d",
			exp3, res3.CombinationId)
	}
}

func TestHasFullHouse(t *testing.T) {
	d := Deck{
		Card{SuitClubs, "3", 3},
		Card{SuitDiamonds, "10", 10},
		Card{SuitClubs, "10", 10},
		Card{SuitHearts, "3", 3},
		Card{SuitDiamonds, "4", 4},
		Card{SuitSpades, "5", 5},
		Card{SuitDiamonds, "3", 3},
	}

	res := d.hasFullHouse()
	exp := HandResult{
		CombinationId: FullHouseId,
		PlayerHand: Deck{
			Card{SuitClubs, "3", 3},
			Card{SuitHearts, "3", 3},
			Card{SuitDiamonds, "3", 3},
			Card{SuitDiamonds, "10", 10},
			Card{SuitClubs, "10", 10},
		},
		CombinationValues: []string{},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 5)
	helperCombinationValues(res, exp, t)

	d[1] = Card{SuitSpades, "Ace", 14}
	res2 := d.hasFullHouse()
	exp2 := NotMatchId

	if res2.CombinationId != exp2 {
		t.Errorf("Expected combinationId %d, got %d",
			exp2, res2.CombinationId)
	}

	d[0] = Card{SuitSpades, "Queen", 12}
	res3 := d.hasFullHouse()
	exp3 := NotMatchId

	if res3.CombinationId != exp3 {
		t.Errorf("Expected combinationId %d, got %d",
			exp3, res3.CombinationId)
	}

	d[0] = Card{SuitClubs, "3", 3}
	d[1] = Card{SuitDiamonds, "10", 10}
	res4 := d.evaluateHand()
	helperComboId(res4, exp, t)
	helperPlayerHand(res4, exp, t, 5)
	helperCombinationValues(res4, exp, t)
}

func TestHasFour(t *testing.T) {
	d := Deck{
		Card{SuitClubs, "3", 3},
		Card{SuitDiamonds, "Jack", 11},
		Card{SuitClubs, "10", 10},
		Card{SuitHearts, "6", 6},
		Card{SuitDiamonds, "10", 10},
		Card{SuitSpades, "10", 10},
		Card{SuitHearts, "10", 10},
	}

	res := d.hasFour()
	exp := HandResult{
		CombinationId: FourKindId,
		PlayerHand: Deck{
			Card{SuitClubs, "10", 10},
			Card{SuitDiamonds, "10", 10},
			Card{SuitSpades, "10", 10},
			Card{SuitHearts, "10", 10},
		},
		CombinationValues: []string{"10"},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 4)
	helperCombinationValues(res, exp, t)

	d2 := Deck{
		Card{SuitClubs, "3", 3},
		Card{SuitDiamonds, "Jack", 11},
		Card{SuitClubs, "5", 5},
		Card{SuitHearts, "6", 6},
		Card{SuitDiamonds, "10", 10},
		Card{SuitSpades, "King", 13},
		Card{SuitSpades, "8", 8},
	}
	res2 := d2.hasFour()
	exp2 := NotMatchId

	if res2.CombinationId != exp2 {
		t.Errorf("Expected combinationId %d, got %d",
			exp2, res2.CombinationId)
	}

	res4 := d.evaluateHand()
	exp.PlayerHand = append(exp.PlayerHand,
		Card{SuitDiamonds, "Jack", 11})
	helperComboId(res4, exp, t)
	helperPlayerHand(res4, exp, t, 5)
	helperCombinationValues(res4, exp, t)
}

func TestHasStraightFlush(t *testing.T) {
	// Royal Flush
	d := Deck{
		Card{SuitSpades, "Ace", 14},
		Card{SuitSpades, "King", 13},
		Card{SuitSpades, "Jack", 11},
		Card{SuitSpades, "10", 10},
		Card{SuitSpades, "Queen", 12},
		Card{SuitSpades, "9", 9},
		Card{SuitSpades, "2", 2},
	}

	res := d.hasStraightFlush()
	exp := HandResult{
		CombinationId: RoyalFlushId,
		PlayerHand: Deck{
			Card{SuitSpades, "Ace", 14},
			Card{SuitSpades, "King", 13},
			Card{SuitSpades, "Queen", 12},
			Card{SuitSpades, "Jack", 11},
			Card{SuitSpades, "10", 10},
		},
		CombinationValues: []string{},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 5)
	helperCombinationValues(res, exp, t)

	// King-high Straight Flush
	d[0] = Card{SuitSpades, "4", 4}
	res2 := d.hasStraightFlush()
	exp2 := HandResult{
		CombinationId: StraightFlushId,
		PlayerHand: Deck{
			Card{SuitSpades, "King", 13},
			Card{SuitSpades, "Queen", 12},
			Card{SuitSpades, "Jack", 11},
			Card{SuitSpades, "10", 10},
			Card{SuitSpades, "9", 9},
		},
		CombinationValues: []string{},
	}

	helperComboId(res2, exp2, t)
	helperPlayerHand(res2, exp2, t, 5)
	helperCombinationValues(res2, exp2, t)

	d[3].Suit, d[4].Suit = SuitHearts, SuitDiamonds
	res3 := d.hasStraightFlush()
	exp3 := NotMatchId

	if res3.CombinationId != exp3 {
		t.Errorf("Expected CombinationId %d, got %d",
			exp3, res3.CombinationId)
	}

	d[3].Suit, d[4].Suit = SuitSpades, SuitSpades
	res4 := d.evaluateHand()
	helperComboId(res4, exp2, t)
	helperPlayerHand(res4, exp2, t, 5)
	helperCombinationValues(res4, exp2, t)
}

func TestHighCard(t *testing.T) {
	d := Deck{
		Card{SuitClubs, "3", 3},
		Card{SuitDiamonds, "Jack", 11},
		Card{SuitClubs, "10", 10},
		Card{SuitHearts, "6", 6},
		Card{SuitDiamonds, "5", 5},
		Card{SuitSpades, "King", 13},
		Card{SuitSpades, "8", 8},
	}

	res := d.evaluateHand()
	exp := HandResult{
		CombinationId: HighCardId,
		PlayerHand: Deck{
			Card{SuitSpades, "King", 13},
			Card{SuitDiamonds, "Jack", 11},
			Card{SuitClubs, "10", 10},
			Card{SuitSpades, "8", 8},
			Card{SuitHearts, "6", 6},
		},
		CombinationValues: []string{},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 2)
	helperCombinationValues(res, exp, t)
}

func helperComboId(res HandResult, exp HandResult, t *testing.T) {
	if res.CombinationId != exp.CombinationId {
		t.Errorf("Expected combinationId %d, got %d",
			exp.CombinationId, res.CombinationId)
	}
}

func helperPlayerHand(res HandResult, exp HandResult, t *testing.T, end int) {
	for i := 0; i < end; i++ {
		if res.PlayerHand[i] != exp.PlayerHand[i] {
			t.Errorf("Expected card at idx #%d: %v of %v, got %v %v",
				i, exp.PlayerHand[i].Value, exp.PlayerHand[i].Suit,
				res.PlayerHand[i].Value, res.PlayerHand[i].Suit)
		}
	}
}

func helperCombinationValues(res HandResult, exp HandResult, t *testing.T) {
	if len(res.CombinationValues) != len(exp.CombinationValues) {
		t.Errorf("Expected combinationValues length %d, got %d",
			len(exp.CombinationValues), len(res.CombinationValues))
	}
}
