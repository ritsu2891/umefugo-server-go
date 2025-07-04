package common

// TrumpCard comparison and utility methods

// getSuitForCompare 比較用のスートを取得する
func getSuitForCompare(method TrumpCardCompareMethod, cardA, cardB *TrumpCard) (TrumpCardSuit, TrumpCardSuit) {
	switch method {
	case TrumpCardCompareMethodStrict:
		if cardA.IsJoker() && cardB.IsJoker() {
			return cardA.AltSuitOrSuit(), cardB.AltSuitOrSuit()
		}
		return cardA.Suit(), cardB.Suit()
	case TrumpCardCompareMethodAllowJokerAltStrict:
		return cardA.AltSuitOrSuit(), cardB.AltSuitOrSuit()
	case TrumpCardCompareMethodAllowJokerAlt:
		suitA := cardA.AltSuitOrSuit()
		if cardA.IsNoAltJoker() {
			suitA = TrumpCardSuitAny
		}
		suitB := cardB.AltSuitOrSuit()
		if cardB.IsNoAltJoker() {
			suitB = TrumpCardSuitAny
		}
		return suitA, suitB
	case TrumpCardCompareMethodAllowJoker:
		suitA := cardA.AltSuitOrSuit()
		if cardA.IsJoker() {
			suitA = TrumpCardSuitAny
		}
		suitB := cardB.AltSuitOrSuit()
		if cardB.IsJoker() {
			suitB = TrumpCardSuitAny
		}
		return suitA, suitB
	default:
		return TrumpCardSuitNotAssigned, TrumpCardSuitNotAssigned
	}
}

// GetNumberForCompare 比較用の数値を取得する
func GetNumberForCompare(method TrumpCardCompareMethod, cardA, cardB *TrumpCard) (int, int) {
	switch method {
	case TrumpCardCompareMethodStrict:
		if cardA.IsJoker() && cardB.IsJoker() {
			return cardA.AltNumberOrNumber(), cardB.AltNumberOrNumber()
		}
		return cardA.Number(), cardB.Number()
	case TrumpCardCompareMethodAllowJokerAltStrict:
		return cardA.AltNumberOrNumber(), cardB.AltNumberOrNumber()
	case TrumpCardCompareMethodAllowJokerAlt:
		numberA := cardA.AltNumberOrNumber()
		if cardA.IsNoAltJoker() {
			numberA = NumberAny
		}
		numberB := cardB.AltNumberOrNumber()
		if cardB.IsNoAltJoker() {
			numberB = NumberAny
		}
		return numberA, numberB
	case TrumpCardCompareMethodAllowJoker:
		numberA := cardA.AltNumberOrNumber()
		if cardA.IsJoker() {
			numberA = NumberAny
		}
		numberB := cardB.AltNumberOrNumber()
		if cardB.IsJoker() {
			numberB = NumberAny
		}
		return numberA, numberB
	default:
		return NumberNA, NumberNA
	}
}

// EqualsSuit 同じスートか
func EqualsSuit(method TrumpCardCompareMethod, cardL, cardR *TrumpCard) bool {
	suitL, suitR := getSuitForCompare(method, cardL, cardR)
	return suitL == TrumpCardSuitAny || suitR == TrumpCardSuitAny || suitL == suitR
}

// EqualsNumber 同じ数値か
func EqualsNumber(method TrumpCardCompareMethod, cardL, cardR *TrumpCard) bool {
	numberL, numberR := GetNumberForCompare(method, cardL, cardR)
	return numberL == NumberAny || numberR == NumberAny || numberL == numberR
}

// EqualsCard 同じ内容のカードか
func EqualsCard(method TrumpCardCompareMethod, cardL, cardR *TrumpCard) bool {
	return EqualsSuit(method, cardL, cardR) && EqualsNumber(method, cardL, cardR)
}

// GreaterThan 数値が大なりか
// cardL > cardR (L is Greater Than R) をtrueとする
func GreaterThan(method TrumpCardCompareMethod, cardL, cardR *TrumpCard) bool {
	numberL, numberR := GetNumberForCompare(method, cardL, cardR)
	if numberL == NumberAny && numberR == NumberAny {
		return false // Infinity > Infinity は不成立とする
	}
	if numberL == NumberAny {
		return true // Infinity > (Any) は常に成立
	}
	if numberR == NumberAny {
		return false // (Any) > Infinity は常に不成立
	}
	return numberL > numberR
}

// GreaterThanOrEquals 数値が大なり以上か
// cardL >= cardR (L is Greater Than Or Equals R) をtrueとする
func GreaterThanOrEquals(method TrumpCardCompareMethod, cardL, cardR *TrumpCard) bool {
	return EqualsNumber(method, cardL, cardR) || GreaterThan(method, cardL, cardR)
}

// LessThan 数値が小なりか
// cardL < cardR (L is Less Than R) をtrueとする
func LessThan(method TrumpCardCompareMethod, cardL, cardR *TrumpCard) bool {
	numberL, numberR := GetNumberForCompare(method, cardL, cardR)
	if numberL == NumberAny && numberR == NumberAny {
		return false // Infinity < Infinity は不成立とする
	}
	if numberL == NumberAny {
		return false // Infinity < (Any) は常に不成立
	}
	if numberR == NumberAny {
		return true // (Any) < Infinity は常に成立
	}
	return numberL < numberR
}

// LessThanOrEquals 数値が小なり以上か
// cardL <= cardR (L is Less Than Or Equals R) をtrueとする
func LessThanOrEquals(method TrumpCardCompareMethod, cardL, cardR *TrumpCard) bool {
	return EqualsNumber(method, cardL, cardR) || LessThan(method, cardL, cardR)
}

// CompareTo ICompareable用比較演算
func (tc *TrumpCard) CompareTo(method TrumpCardCompareMethod, other *TrumpCard) int {
	if other == nil {
		return 1
	}
	if EqualsCard(method, tc, other) {
		return 0
	}
	if GreaterThan(method, tc, other) {
		return 1
	}
	return -1
}

// Instance methods for TrumpCard

// EqualsSuit 同じスートか
func (tc *TrumpCard) EqualsSuit(method TrumpCardCompareMethod, other *TrumpCard) bool {
	return EqualsSuit(method, tc, other)
}

// EqualsNumber 同じ数値か
func (tc *TrumpCard) EqualsNumber(method TrumpCardCompareMethod, other *TrumpCard) bool {
	return EqualsNumber(method, tc, other)
}

// EqualsCard 同じ内容のカードか
func (tc *TrumpCard) EqualsCard(method TrumpCardCompareMethod, other *TrumpCard) bool {
	return EqualsCard(method, tc, other)
}

// GreaterThan 数値が大なりか
// this > other (this is Greater Than target) をtrueとする
func (tc *TrumpCard) GreaterThan(method TrumpCardCompareMethod, other *TrumpCard) bool {
	return GreaterThan(method, tc, other)
}

// GreaterThanOrEquals 数値が大なり以上か
// this >= other (this is Greater Than Or Equals target) をtrueとする
func (tc *TrumpCard) GreaterThanOrEquals(method TrumpCardCompareMethod, other *TrumpCard) bool {
	return GreaterThanOrEquals(method, tc, other)
}

// LessThan 数値が小なりか
// this < other (this is Less Than target) をtrueとする
func (tc *TrumpCard) LessThan(method TrumpCardCompareMethod, other *TrumpCard) bool {
	return LessThan(method, tc, other)
}

// LessThanOrEquals 数値が小なり以上か
// this <= other (this is Less Than Or Equals target) をtrueとする
func (tc *TrumpCard) LessThanOrEquals(method TrumpCardCompareMethod, other *TrumpCard) bool {
	return LessThanOrEquals(method, tc, other)
}

// RotateNumber 数値を指定個数回転させる
func RotateNumber(number, amount int) int {
	amount %= NumbersCount
	return (number+NumbersCount+amount-1)%NumbersCount + 1
}

// PreviousNumber 数値を一個前に回転させる
func PreviousNumber(number int) int {
	return RotateNumber(number, -1)
}

// NextNumber 数値を一個後に回転させる
func NextNumber(number int) int {
	return RotateNumber(number, 1)
}

// RotateNumber 数値を指定個数回転させる
func (tc *TrumpCard) RotateNumber(amount int) int {
	return RotateNumber(tc.number, amount)
}

// PreviousNumber 数値を一個前に回転させる
func (tc *TrumpCard) PreviousNumber() int {
	return PreviousNumber(tc.number)
}

// NextNumber 数値を一個後に回転させる
func (tc *TrumpCard) NextNumber() int {
	return NextNumber(tc.number)
}
