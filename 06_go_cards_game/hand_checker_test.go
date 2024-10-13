package main

import "testing"

func TestHasPair(t *testing.T) {
	d := deck{
		card{SuitClubs, "3", 3},
		card{SuitDiamonds, "Jack", 11},
		card{SuitClubs, "10", 10},
		card{SuitHearts, "6", 6},
		card{SuitDiamonds, "10", 10},
		card{SuitSpades, "King", 13},
		card{SuitSpades, "8", 8},
	}

	res := d.hasPair()
	exp := handResult{
		combinationId: SinglePairId,
		playerHand: deck{
			card{SuitClubs, "10", 10},
			card{SuitDiamonds, "10", 10},
			card{SuitSpades, "King", 13},
			card{SuitDiamonds, "Jack", 11},
			card{SuitSpades, "8", 8},
		},
		combinationValues: []string{"10"},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 2)
	helperCombinationValues(res, exp, t)

	d2 := deck{
		card{SuitClubs, "3", 3},
		card{SuitDiamonds, "Jack", 11},
		card{SuitClubs, "5", 5},
		card{SuitHearts, "6", 6},
		card{SuitDiamonds, "10", 10},
		card{SuitSpades, "King", 13},
		card{SuitSpades, "8", 8},
	}
	res2 := d2.hasPair()
	exp2 := NotMatchId

	if res2.combinationId != exp2 {
		t.Errorf("Expected combinationId %d, got %d",
			exp2, res2.combinationId)
	}

	res3 := d.evaluateHand()
	helperComboId(res3, exp, t)
	helperPlayerHand(res3, exp, t, 5)
	helperCombinationValues(res3, exp, t)
}

func TestHasTwoPair(t *testing.T) {
	d := deck{
		card{SuitClubs, "3", 3},
		card{SuitDiamonds, "Jack", 11},
		card{SuitClubs, "10", 10},
		card{SuitHearts, "6", 6},
		card{SuitDiamonds, "King", 13},
		card{SuitSpades, "3", 3},
		card{SuitHearts, "10", 10},
	}

	res := d.hasTwoPair()
	exp := handResult{
		combinationId: TwoPairId,
		playerHand: deck{
			card{SuitClubs, "10", 10},
			card{SuitHearts, "10", 10},
			card{SuitClubs, "3", 3},
			card{SuitSpades, "3", 3},
		},
		combinationValues: []string{"10", "3"},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 4)
	helperCombinationValues(res, exp, t)

	d[2] = card{SuitHearts, "7", 7}
	res2 := d.hasTwoPair()
	exp2 := NotMatchId

	if res2.combinationId != exp2 {
		t.Errorf("Expected combinationId %d, got %d",
			exp2, res2.combinationId)
	}

	d[0] = card{SuitHearts, "4", 4}
	res3 := d.hasTwoPair()
	exp3 := NotMatchId

	if res3.combinationId != exp3 {
		t.Errorf("Expected combinationId %d, got %d",
			exp3, res3.combinationId)
	}

	d[0] = card{SuitClubs, "3", 3}
	d[2] = card{SuitClubs, "10", 10}

	res4 := d.evaluateHand()
	exp.playerHand = append(exp.playerHand,
		card{SuitDiamonds, "King", 13})
	helperComboId(res4, exp, t)
	helperPlayerHand(res4, exp, t, 5)
	helperCombinationValues(res4, exp, t)
}

func TestHasThree(t *testing.T) {
	d := deck{
		card{SuitClubs, "3", 3},
		card{SuitDiamonds, "Jack", 11},
		card{SuitClubs, "10", 10},
		card{SuitHearts, "6", 6},
		card{SuitDiamonds, "10", 10},
		card{SuitSpades, "King", 13},
		card{SuitHearts, "10", 10},
	}

	res := d.hasThree()
	exp := handResult{
		combinationId: ThreeKindId,
		playerHand: deck{
			card{SuitClubs, "10", 10},
			card{SuitDiamonds, "10", 10},
			card{SuitHearts, "10", 10},
		},
		combinationValues: []string{"10"},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 3)
	helperCombinationValues(res, exp, t)

	d2 := deck{
		card{SuitClubs, "3", 3},
		card{SuitDiamonds, "Jack", 11},
		card{SuitClubs, "5", 5},
		card{SuitHearts, "6", 6},
		card{SuitDiamonds, "10", 10},
		card{SuitSpades, "King", 13},
		card{SuitSpades, "8", 8},
	}
	res2 := d2.hasThree()
	exp2 := NotMatchId

	if res2.combinationId != exp2 {
		t.Errorf("Expected combinationId %d, got %d",
			exp2, res2.combinationId)
	}

	res3 := d.evaluateHand()
	exp.playerHand = append(exp.playerHand,
		card{SuitSpades, "King", 13},
		card{SuitDiamonds, "Jack", 11})
	helperComboId(res3, exp, t)
	helperPlayerHand(res3, exp, t, 5)
	helperCombinationValues(res3, exp, t)
}

func TestHasStraight(t *testing.T) {
	d := deck{
		card{SuitClubs, "3", 3},
		card{SuitDiamonds, "Jack", 11},
		card{SuitClubs, "10", 10},
		card{SuitHearts, "6", 6},
		card{SuitDiamonds, "4", 4},
		card{SuitSpades, "5", 5},
		card{SuitHearts, "7", 7},
	}

	res := d.hasStraight()
	exp := handResult{
		combinationId: StraightId,
		playerHand: deck{
			card{SuitHearts, "7", 7},
			card{SuitHearts, "6", 6},
			card{SuitSpades, "5", 5},
			card{SuitDiamonds, "4", 4},
			card{SuitClubs, "3", 3},
		},
		combinationValues: []string{},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 5)
	helperCombinationValues(res, exp, t)

	d[0] = card{SuitSpades, "Ace", 14}
	res2 := d.hasStraight()
	exp2 := NotMatchId

	if res2.combinationId != exp2 {
		t.Errorf("Expected combinationId %d, got %d",
			exp2, res2.combinationId)
	}

	// Low straight A 2 3 4 5
	d = deck{
		card{SuitDiamonds, "10", 10},
		card{SuitHearts, "5", 5},
		card{SuitDiamonds, "Ace", 14},
		card{SuitClubs, "4", 4},
		card{SuitSpades, "Queen", 12},
		card{SuitHearts, "2", 2},
		card{SuitClubs, "3", 3},
	}

	res3 := d.hasStraight()
	exp3 := handResult{
		combinationId: StraightId,
		playerHand: deck{
			card{SuitHearts, "5", 5},
			card{SuitClubs, "4", 4},
			card{SuitClubs, "3", 3},
			card{SuitHearts, "2", 2},
			card{SuitDiamonds, "Ace", 14},
		},
		combinationValues: []string{},
	}

	helperComboId(res3, exp3, t)
	helperPlayerHand(res3, exp3, t, 5)
	helperCombinationValues(res3, exp3, t)
}

func TestHasFlush(t *testing.T) {
	d := deck{
		card{SuitHearts, "3", 3},
		card{SuitHearts, "Jack", 11},
		card{SuitHearts, "10", 10},
		card{SuitDiamonds, "6", 6},
		card{SuitHearts, "4", 4},
		card{SuitSpades, "5", 5},
		card{SuitHearts, "7", 7},
	}

	res := d.hasFlush()
	exp := handResult{
		combinationId: FlushId,
		playerHand: deck{
			card{SuitHearts, "Jack", 11},
			card{SuitHearts, "10", 10},
			card{SuitHearts, "7", 7},
			card{SuitHearts, "4", 4},
			card{SuitHearts, "3", 3},
		},
		combinationValues: []string{},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 5)
	helperCombinationValues(res, exp, t)

	d[0] = card{SuitSpades, "Ace", 14}
	res2 := d.hasFlush()
	exp2 := NotMatchId

	if res2.combinationId != exp2 {
		t.Errorf("Expected combinationId %d, got %d",
			exp2, res2.combinationId)
	}
}

func TestHasFullHouse(t *testing.T) {
	d := deck{
		card{SuitClubs, "3", 3},
		card{SuitDiamonds, "10", 10},
		card{SuitClubs, "10", 10},
		card{SuitHearts, "3", 3},
		card{SuitDiamonds, "4", 4},
		card{SuitSpades, "5", 5},
		card{SuitDiamonds, "3", 3},
	}

	res := d.hasFullHouse()
	exp := handResult{
		combinationId: FullHouseId,
		playerHand: deck{
			card{SuitClubs, "3", 3},
			card{SuitHearts, "3", 3},
			card{SuitDiamonds, "3", 3},
			card{SuitDiamonds, "10", 10},
			card{SuitClubs, "10", 10},
		},
		combinationValues: []string{},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 5)
	helperCombinationValues(res, exp, t)

	d[1] = card{SuitSpades, "Ace", 14}
	res2 := d.hasFullHouse()
	exp2 := NotMatchId

	if res2.combinationId != exp2 {
		t.Errorf("Expected combinationId %d, got %d",
			exp2, res2.combinationId)
	}

	d[0] = card{SuitSpades, "Queen", 12}
	res3 := d.hasFullHouse()
	exp3 := NotMatchId

	if res3.combinationId != exp3 {
		t.Errorf("Expected combinationId %d, got %d",
			exp3, res3.combinationId)
	}
}

func TestHasFour(t *testing.T) {
	d := deck{
		card{SuitClubs, "3", 3},
		card{SuitDiamonds, "Jack", 11},
		card{SuitClubs, "10", 10},
		card{SuitHearts, "6", 6},
		card{SuitDiamonds, "10", 10},
		card{SuitSpades, "10", 10},
		card{SuitHearts, "10", 10},
	}

	res := d.hasFour()
	exp := handResult{
		combinationId: FourKindId,
		playerHand: deck{
			card{SuitClubs, "10", 10},
			card{SuitDiamonds, "10", 10},
			card{SuitSpades, "10", 10},
			card{SuitHearts, "10", 10},
		},
		combinationValues: []string{"10"},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 4)
	helperCombinationValues(res, exp, t)

	d2 := deck{
		card{SuitClubs, "3", 3},
		card{SuitDiamonds, "Jack", 11},
		card{SuitClubs, "5", 5},
		card{SuitHearts, "6", 6},
		card{SuitDiamonds, "10", 10},
		card{SuitSpades, "King", 13},
		card{SuitSpades, "8", 8},
	}
	res2 := d2.hasFour()
	exp2 := NotMatchId

	if res2.combinationId != exp2 {
		t.Errorf("Expected combinationId %d, got %d",
			exp2, res2.combinationId)
	}

	res4 := d.evaluateHand()
	exp.playerHand = append(exp.playerHand,
		card{SuitDiamonds, "Jack", 11})
	helperComboId(res4, exp, t)
	helperPlayerHand(res4, exp, t, 5)
	helperCombinationValues(res4, exp, t)
}

func TestHasStraightFlush(t *testing.T) {
	// Royal Flush
	d := deck{
		card{SuitSpades, "Ace", 14},
		card{SuitSpades, "King", 13},
		card{SuitSpades, "Jack", 11},
		card{SuitSpades, "10", 10},
		card{SuitSpades, "Queen", 12},
		card{SuitSpades, "9", 9},
		card{SuitSpades, "2", 2},
	}

	res := d.hasStraightFlush()
	exp := handResult{
		combinationId: RoyalFlushId,
		playerHand: deck{
			card{SuitSpades, "Ace", 14},
			card{SuitSpades, "King", 13},
			card{SuitSpades, "Queen", 12},
			card{SuitSpades, "Jack", 11},
			card{SuitSpades, "10", 10},
		},
		combinationValues: []string{},
	}

	helperComboId(res, exp, t)
	helperPlayerHand(res, exp, t, 5)
	helperCombinationValues(res, exp, t)

	// King-high Straight Flush
	d[0] = card{SuitSpades, "4", 4}
	res2 := d.hasStraightFlush()
	exp2 := handResult{
		combinationId: StraightFlushId,
		playerHand: deck{
			card{SuitSpades, "King", 13},
			card{SuitSpades, "Queen", 12},
			card{SuitSpades, "Jack", 11},
			card{SuitSpades, "10", 10},
			card{SuitSpades, "9", 9},
		},
		combinationValues: []string{},
	}

	helperComboId(res2, exp2, t)
	helperPlayerHand(res2, exp2, t, 5)
	helperCombinationValues(res2, exp2, t)

	d[3].suit, d[4].suit = SuitHearts, SuitDiamonds
	res3 := d.hasStraightFlush()
	exp3 := NotMatchId

	if res3.combinationId != exp3 {
		t.Errorf("Expected combinationId %d, got %d",
			exp3, res3.combinationId)
	}
}

func helperComboId(res handResult, exp handResult, t *testing.T) {
	if res.combinationId != exp.combinationId {
		t.Errorf("Expected combinationId %d, got %d",
			exp.combinationId, res.combinationId)
	}
}

func helperPlayerHand(res handResult, exp handResult, t *testing.T, end int) {
	for i := 0; i < end; i++ {
		if res.playerHand[i] != exp.playerHand[i] {
			t.Errorf("Expected card at idx #%d: %v of %v, got %v %v",
				i, exp.playerHand[i].value, exp.playerHand[i].suit,
				res.playerHand[i].value, res.playerHand[i].suit)
		}
	}
}

func helperCombinationValues(res handResult, exp handResult, t *testing.T) {
	if len(res.combinationValues) != len(exp.combinationValues) {
		t.Errorf("Expected combinationValues length %d, got %d",
			len(exp.combinationValues), len(res.combinationValues))
	}
}
