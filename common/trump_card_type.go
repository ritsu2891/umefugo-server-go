package common

// TrumpCardType トランプカード種別
type TrumpCardType int

const (
	// TrumpCardTypeNormal 通常カード
	TrumpCardTypeNormal TrumpCardType = iota
	// TrumpCardTypeJoker ジョーカーカード
	TrumpCardTypeJoker
)

// String returns the string representation of TrumpCardType
func (t TrumpCardType) String() string {
	switch t {
	case TrumpCardTypeNormal:
		return "通常カード"
	case TrumpCardTypeJoker:
		return "ジョーカーカード"
	default:
		return "不明"
	}
}
