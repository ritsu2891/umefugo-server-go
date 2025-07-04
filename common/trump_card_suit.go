package common

import "fmt"

// TrumpCardSuit トランプカードスート
type TrumpCardSuit int

const (
	// TrumpCardSuitClub クラブ
	TrumpCardSuitClub TrumpCardSuit = iota
	// TrumpCardSuitDiamond ダイアモンド
	TrumpCardSuitDiamond
	// TrumpCardSuitHeart ハート
	TrumpCardSuitHeart
	// TrumpCardSuitSpade スペード
	TrumpCardSuitSpade
	// TrumpCardSuitJoker ジョーカー
	TrumpCardSuitJoker
	// TrumpCardSuitAny 任意（通常スート全てと同じ扱い）
	TrumpCardSuitAny
	// TrumpCardSuitNotAssigned 未割当
	TrumpCardSuitNotAssigned
)

// ToLabelString ラベル文字列を取得する
func (s TrumpCardSuit) ToLabelString() string {
	switch s {
	case TrumpCardSuitClub:
		return "クラブ"
	case TrumpCardSuitDiamond:
		return "ダイアモンド"
	case TrumpCardSuitHeart:
		return "ハート"
	case TrumpCardSuitSpade:
		return "スペード"
	case TrumpCardSuitJoker:
		return "ジョーカー"
	case TrumpCardSuitAny:
		return "任意"
	case TrumpCardSuitNotAssigned:
		return "未割当"
	default:
		panic(fmt.Sprintf("Invalid suit: %d", s))
	}
}

// ToMarkString マーク文字列を取得する
func (s TrumpCardSuit) ToMarkString() string {
	switch s {
	case TrumpCardSuitClub:
		return "♣"
	case TrumpCardSuitDiamond:
		return "♦"
	case TrumpCardSuitHeart:
		return "♥"
	case TrumpCardSuitSpade:
		return "♠"
	case TrumpCardSuitJoker:
		return "[J]"
	case TrumpCardSuitAny:
		return "[*]"
	case TrumpCardSuitNotAssigned:
		return "[/]"
	default:
		panic(fmt.Sprintf("Invalid suit: %d", s))
	}
}

// String returns the string representation of TrumpCardSuit
func (s TrumpCardSuit) String() string {
	return s.ToLabelString()
}
