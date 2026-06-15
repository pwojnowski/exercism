package blackjack

var cardValues = map[string]int{
	"ace":     11,
	"two":     2,
	"three":   3,
	"four":    4,
	"five":    5,
	"six":     6,
	"seven":   7,
	"eight":   8,
	"nine":    9,
	"ten":     10,
	"jack":    10,
	"queen":   10,
	"king":    10,
	"*other*": 0,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cardValues[card]
}

func isBlackjack(card1Val, card2Val int) bool {
	return (card1Val + card2Val) == 21
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var card1Val = ParseCard(card1)
	var card2Val = ParseCard(card2)
	var sum = card1Val + card2Val
	var dCardVal = ParseCard(dealerCard)
	var decision string
	switch {
	case card1 == "ace" && card2 == "ace":
		decision = "P"
	case sum == 21 && dCardVal < 10:
		decision = "W"
	case (12 <= sum && sum <= 16) && 7 <= dCardVal:
		decision = "H"
	case (12 <= sum && sum <= 16):
		decision = "S"
	case sum <= 11:
		decision = "H"
	case dCardVal >= 10:
		decision = "S"
	case 17 <= sum && sum <= 20:
		decision = "S"
	}
	return decision
}
