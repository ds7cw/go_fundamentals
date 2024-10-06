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
