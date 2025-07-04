package common

import (
	"fmt"
	"strconv"
)

// TrumpCard constants
const (
	// NumbersCount 通常数字の個数
	NumbersCount = 13
	// NumberNA 数値未割当
	NumberNA = -1
	// NumberMin 最小数値
	NumberMin = 1
	// NumberMax 最大数値
	NumberMax = 13
	// NumberAny 任意数値
	NumberAny = 0
	// NumberJoker ジョーカー数値
	NumberJoker = 99
)

// Mark constants for number representation
const (
	MarkNumber1  = "A"
	MarkNumber11 = "J"
	MarkNumber12 = "Q"
	MarkNumber13 = "K"
)

var (
	// Suits 全てのスート
	Suits = []TrumpCardSuit{
		TrumpCardSuitClub,
		TrumpCardSuitDiamond,
		TrumpCardSuitHeart,
		TrumpCardSuitSpade,
	}

	// Numbers 全ての通常数値
	Numbers = make([]int, NumbersCount)

	// CardNA 未割当カード
	CardNA *TrumpCard
)

// init initializes global variables
func init() {
	for i := 0; i < NumbersCount; i++ {
		Numbers[i] = NumberMin + i
	}
	CardNA = &TrumpCard{
		Card:      NewCard(),
		cardType:  TrumpCardTypeNormal,
		suit:      TrumpCardSuitNotAssigned,
		number:    NumberNA,
		altSuit:   TrumpCardSuitNotAssigned,
		altNumber: NumberNA,
	}
}

// TrumpCard トランプカード
type TrumpCard struct {
	*Card
	// cardType 種別
	cardType TrumpCardType
	// suit スート
	suit TrumpCardSuit
	// number 数値
	number int
	// altSuit 代役スート
	altSuit TrumpCardSuit
	// altNumber 代役数値
	altNumber int
}

// NewTrumpCard creates a new TrumpCard instance (private constructor)
func newTrumpCard(cardType TrumpCardType, suit TrumpCardSuit, number int) *TrumpCard {
	return &TrumpCard{
		Card:      NewCard(),
		cardType:  cardType,
		suit:      suit,
		number:    number,
		altSuit:   TrumpCardSuitNotAssigned,
		altNumber: NumberNA,
	}
}

// NewNormalTrumpCard 通常カードを生成する
func NewNormalTrumpCard(suit TrumpCardSuit, number int) (*TrumpCard, error) {
	if !isValidNormalSuitAndNumber(suit, number) {
		return nil, fmt.Errorf("スートまたは数値が不正です")
	}
	return newTrumpCard(TrumpCardTypeNormal, suit, number), nil
}

// NewJokerTrumpCard ジョーカーを生成する
func NewJokerTrumpCard() *TrumpCard {
	return newTrumpCard(TrumpCardTypeJoker, TrumpCardSuitJoker, NumberJoker)
}

// NewJokerTrumpCardWithAltSuit 代役スート設定済みのジョーカーを生成する
func NewJokerTrumpCardWithAltSuit(altSuit TrumpCardSuit) (*TrumpCard, error) {
	card := NewJokerTrumpCard()
	err := card.SetAltSuit(altSuit)
	if err != nil {
		return nil, err
	}
	return card, nil
}

// NewJokerTrumpCardWithAltNumber 代役数値設定済みのジョーカーを生成する
func NewJokerTrumpCardWithAltNumber(altNumber int) (*TrumpCard, error) {
	card := NewJokerTrumpCard()
	err := card.SetAltNumber(altNumber)
	if err != nil {
		return nil, err
	}
	return card, nil
}

// NewJokerTrumpCardWithAlt 代役スート・数値設定済みのジョーカーを生成する
func NewJokerTrumpCardWithAlt(altSuit TrumpCardSuit, altNumber int) (*TrumpCard, error) {
	card := NewJokerTrumpCard()
	err := card.SetAlt(altSuit, altNumber)
	if err != nil {
		return nil, err
	}
	return card, nil
}

// Type returns the card type
func (tc *TrumpCard) Type() TrumpCardType {
	return tc.cardType
}

// Suit returns the card suit
func (tc *TrumpCard) Suit() TrumpCardSuit {
	return tc.suit
}

// Number returns the card number
func (tc *TrumpCard) Number() int {
	return tc.number
}

// AltSuit returns the alternative suit
func (tc *TrumpCard) AltSuit() TrumpCardSuit {
	return tc.altSuit
}

// AltNumber returns the alternative number
func (tc *TrumpCard) AltNumber() int {
	return tc.altNumber
}

// IsJoker ジョーカーかどうか
func (tc *TrumpCard) IsJoker() bool {
	return tc.cardType == TrumpCardTypeJoker
}

// IsNA カードなしを表すかどうか
func (tc *TrumpCard) IsNA() bool {
	return tc.Card.Equals(CardNA.Card)
}

// HasAltSuit 代役スートが設定されているか
func (tc *TrumpCard) HasAltSuit() bool {
	return tc.altSuit != TrumpCardSuitNotAssigned
}

// HasAltNumber 代役数値が設定されているか
func (tc *TrumpCard) HasAltNumber() bool {
	return tc.altNumber != NumberNA
}

// HasAltBoth 代役スート・数値両方が設定されているか
func (tc *TrumpCard) HasAltBoth() bool {
	return tc.HasAltSuit() && tc.HasAltNumber()
}

// HasAltAny 代役スート・数値いずれかが設定されているか
func (tc *TrumpCard) HasAltAny() bool {
	return tc.HasAltSuit() || tc.HasAltNumber()
}

// IsAltJoker 代役情報ありジョーカーかどうか
func (tc *TrumpCard) IsAltJoker() bool {
	return tc.IsJoker() && tc.HasAltAny()
}

// IsNoAltJoker 代役情報なしジョーカーかどうか
func (tc *TrumpCard) IsNoAltJoker() bool {
	return tc.IsJoker() && !tc.HasAltAny()
}

// IsAltNumberJoker 代役数値情報ありジョーカーかどうか
func (tc *TrumpCard) IsAltNumberJoker() bool {
	return tc.IsJoker() && tc.HasAltNumber()
}

// IsNoAltNumberJoker 代役数値情報なしジョーカーかどうか
func (tc *TrumpCard) IsNoAltNumberJoker() bool {
	return tc.IsJoker() && !tc.HasAltNumber()
}

// IsAltSuitJoker 代役スートありジョーカーかどうか
func (tc *TrumpCard) IsAltSuitJoker() bool {
	return tc.IsJoker() && tc.HasAltSuit()
}

// IsNoAltSuitJoker 代役スートなしジョーカーかどうか
func (tc *TrumpCard) IsNoAltSuitJoker() bool {
	return tc.IsJoker() && !tc.HasAltSuit()
}

// AltSuitOrSuit 代役スートまたはスート
func (tc *TrumpCard) AltSuitOrSuit() TrumpCardSuit {
	if tc.HasAltSuit() {
		return tc.altSuit
	}
	return tc.suit
}

// AltNumberOrNumber 代役数値または数値
func (tc *TrumpCard) AltNumberOrNumber() int {
	if tc.HasAltNumber() {
		return tc.altNumber
	}
	return tc.number
}

// SetAltSuit 代役スートを設定する
func (tc *TrumpCard) SetAltSuit(suit TrumpCardSuit) error {
	if !tc.IsJoker() {
		return fmt.Errorf("ジョーカーでないカードに代役情報は設定できません")
	}
	if !isValidNormalSuit(suit) {
		return fmt.Errorf("スートが不正です")
	}
	tc.altSuit = suit
	return nil
}

// SetAltNumber 代役数値を設定する
func (tc *TrumpCard) SetAltNumber(number int) error {
	if !tc.IsJoker() {
		return fmt.Errorf("ジョーカーでないカードに代役情報は設定できません")
	}
	if !isValidNormalNumber(number) {
		return fmt.Errorf("数値が不正です")
	}
	tc.altNumber = number
	return nil
}

// SetAlt 代役スート・数値を設定する
func (tc *TrumpCard) SetAlt(suit TrumpCardSuit, number int) error {
	if err := tc.SetAltSuit(suit); err != nil {
		return err
	}
	return tc.SetAltNumber(number)
}

// ClearAltSuit 代役スートをクリアする
func (tc *TrumpCard) ClearAltSuit() error {
	if !tc.IsJoker() {
		return fmt.Errorf("ジョーカーでないカードに代役情報は設定できません")
	}
	tc.altSuit = TrumpCardSuitNotAssigned
	return nil
}

// ClearAltNumber 代役数値をクリアする
func (tc *TrumpCard) ClearAltNumber() error {
	if !tc.IsJoker() {
		return fmt.Errorf("ジョーカーでないカードに代役情報は設定できません")
	}
	tc.altNumber = NumberNA
	return nil
}

// ClearAlt 代役スート・数値をクリアする
func (tc *TrumpCard) ClearAlt() error {
	if err := tc.ClearAltSuit(); err != nil {
		return err
	}
	return tc.ClearAltNumber()
}

// Helper functions for validation
func isValidNormalSuit(suit TrumpCardSuit) bool {
	return suit != TrumpCardSuitAny && suit != TrumpCardSuitNotAssigned && suit != TrumpCardSuitJoker
}

func isValidNormalNumber(number int) bool {
	return number >= NumberMin && number <= NumberMax
}

func isValidNormalSuitAndNumber(suit TrumpCardSuit, number int) bool {
	return isValidNormalSuit(suit) && isValidNormalNumber(number)
}

// ToNumberString 数値ラベル文字列を取得する
func (tc *TrumpCard) ToNumberString() string {
	switch tc.number {
	case 1:
		return MarkNumber1
	case 11:
		return MarkNumber11
	case 12:
		return MarkNumber12
	case 13:
		return MarkNumber13
	default:
		return strconv.Itoa(tc.number)
	}
}

// String 文字列ラベルを取得する
func (tc *TrumpCard) String() string {
	switch tc.cardType {
	case TrumpCardTypeNormal:
		return fmt.Sprintf("%s%s", tc.suit.ToMarkString(), tc.ToNumberString())
	case TrumpCardTypeJoker:
		return tc.suit.ToMarkString()
	default:
		panic(fmt.Sprintf("Invalid type: %v", tc.cardType))
	}
}
