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

	if res.combinationId != exp.combinationId {
		t.Errorf("Expected combinationId %d, got %d",
			exp.combinationId, res.combinationId)
	}

	if res.combinationValues[0] != exp.combinationValues[0] {
		t.Errorf("Expected combinationValues %v, got %v",
			exp.combinationValues[0], res.combinationValues[0])
	}

	if res.playerHand[0] != exp.playerHand[0] || res.playerHand[1] != exp.playerHand[1] {
		t.Errorf("Expected pair %v %v, got %v %v",
			exp.combinationValues[0], res.combinationValues[0], exp.combinationValues[1], res.combinationValues[1])
	}

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

	if res.combinationId != exp.combinationId {
		t.Errorf("Expected combinationId %d, got %d",
			exp.combinationId, res.combinationId)
	}

	for i := 0; i < 4; i++ {
		if res.playerHand[i] != exp.playerHand[i] {
			t.Errorf("Expected card at idx #%d: %v of %v, got %v %v",
				i, exp.playerHand[i].value, exp.playerHand[i].suit,
				res.playerHand[i].value, res.playerHand[i].suit)
		}
	}

	if exp.combinationValues[0] != res.combinationValues[0] ||
		exp.combinationValues[1] != res.combinationValues[1] {
		t.Errorf("Expected combinationValues %q, got %q",
			exp.combinationValues, res.combinationValues)
	}

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
			card{SuitSpades, "King", 13},
			card{SuitDiamonds, "Jack", 11},
		},
		combinationValues: []string{"10"},
	}

	if res.combinationId != exp.combinationId {
		t.Errorf("Expected combinationId %d, got %d",
			exp.combinationId, res.combinationId)
	}

	if res.combinationValues[0] != exp.combinationValues[0] {
		t.Errorf("Expected combinationValues %v, got %v",
			exp.combinationValues[0], res.combinationValues[0])
	}

	if res.playerHand[0] != exp.playerHand[0] || res.playerHand[1] != exp.playerHand[1] ||
		res.playerHand[2] != exp.playerHand[2] {
		t.Errorf("Expected pair %v %v %v, got %v %v %v",
			exp.combinationValues[0], res.combinationValues[0],
			exp.combinationValues[1], res.combinationValues[1],
			exp.combinationValues[2], res.combinationValues[2])
	}

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

	if res.combinationId != exp.combinationId {
		t.Errorf("Expected combinationId %d, got %d",
			exp.combinationId, res.combinationId)
	}

	for i := 0; i < 4; i++ {
		if res.playerHand[i] != exp.playerHand[i] {
			t.Errorf("Expected card at idx #%d: %v of %v, got %v %v",
				i, exp.playerHand[i].value, exp.playerHand[i].suit,
				res.playerHand[i].value, res.playerHand[i].suit)
		}
	}

	if len(res.combinationValues) != len(exp.combinationValues) {
		t.Errorf("Expected combinationValues length %d, got %d",
			len(exp.combinationValues), len(res.combinationValues))
	}

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
			card{SuitDiamonds, "Ace", 14},
			card{SuitHearts, "2", 2},
			card{SuitClubs, "3", 3},
			card{SuitClubs, "4", 4},
			card{SuitHearts, "5", 5},
		},
		combinationValues: []string{},
	}

	if res.combinationId != exp3.combinationId {
		t.Errorf("Expected combinationId %d, got %d",
			exp3.combinationId, res3.combinationId)
	}

	for i := 0; i < 4; i++ {
		if res3.playerHand[i] != exp3.playerHand[i] {
			t.Errorf("Expected card at idx #%d: %v of %v, got %v %v",
				i, exp3.playerHand[i].value, exp3.playerHand[i].suit,
				res3.playerHand[i].value, res3.playerHand[i].suit)
		}
	}

	if len(res3.combinationValues) != len(exp3.combinationValues) {
		t.Errorf("Expected combinationValues length %d, got %d",
			len(exp3.combinationValues), len(res3.combinationValues))
	}
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

	if res.combinationId != exp.combinationId {
		t.Errorf("Expected combinationId %d, got %d",
			exp.combinationId, res.combinationId)
	}

	for i := 0; i < 4; i++ {
		if res.playerHand[i] != exp.playerHand[i] {
			t.Errorf("Expected card at idx #%d: %v of %v, got %v %v",
				i, exp.playerHand[i].value, exp.playerHand[i].suit,
				res.playerHand[i].value, res.playerHand[i].suit)
		}
	}

	if len(res.combinationValues) != len(exp.combinationValues) {
		t.Errorf("Expected combinationValues length %d, got %d",
			len(exp.combinationValues), len(res.combinationValues))
	}

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

	if res.combinationId != exp.combinationId {
		t.Errorf("Expected combinationId %d, got %d",
			exp.combinationId, res.combinationId)
	}

	for i := 0; i < 4; i++ {
		if res.playerHand[i] != exp.playerHand[i] {
			t.Errorf("Expected card at idx #%d: %v of %v, got %v %v",
				i, exp.playerHand[i].value, exp.playerHand[i].suit,
				res.playerHand[i].value, res.playerHand[i].suit)
		}
	}

	if len(res.combinationValues) != len(exp.combinationValues) {
		t.Errorf("Expected combinationValues length %d, got %d",
			len(exp.combinationValues), len(res.combinationValues))
	}

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
			card{SuitDiamonds, "Jack", 11},
		},
		combinationValues: []string{"10"},
	}

	if res.combinationId != exp.combinationId {
		t.Errorf("Expected combinationId %d, got %d",
			exp.combinationId, res.combinationId)
	}

	if res.combinationValues[0] != exp.combinationValues[0] {
		t.Errorf("Expected combinationValues %v, got %v",
			exp.combinationValues[0], res.combinationValues[0])
	}

	if res.playerHand[0] != exp.playerHand[0] || res.playerHand[1] != exp.playerHand[1] ||
		res.playerHand[2] != exp.playerHand[2] || res.playerHand[3] != exp.playerHand[3] {
		t.Errorf("Expected pair %v %v %v %v, got %v %v %v %v",
			exp.combinationValues[0], exp.combinationValues[1],
			exp.combinationValues[2], exp.combinationValues[3],
			res.combinationValues[0], res.combinationValues[1],
			res.combinationValues[2], res.combinationValues[3])
	}

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
}
