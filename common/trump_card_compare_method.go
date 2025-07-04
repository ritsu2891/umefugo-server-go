package common

// TrumpCardCompareMethod トランプカード比較方法
type TrumpCardCompareMethod int

const (
	// TrumpCardCompareMethodStrict ジョーカーに関係なく厳密にスート・数値を比較する
	// 代役無ジョーカー ↔ 通常カード：常に違うカード
	// 代役有ジョーカー ↔ 通常カード：常に違うカード
	// 代役無ジョーカー ↔ 代役有ジョーカー：常に違うカード
	// 代役有ジョーカー ↔ 代役有ジョーカー：代役スート・代役数値を比較
	TrumpCardCompareMethodStrict TrumpCardCompareMethod = iota

	// TrumpCardCompareMethodAllowJokerAltStrict ジョーカーは代役情報がある場合のみスート・数値を比較する
	// 代役無ジョーカー ↔ 通常カード：常に違うカード
	// 代役有ジョーカー ↔ 通常カード：代役スート・代役数値を比較
	// 代役無ジョーカー ↔ 代役有ジョーカー：常に違うカード
	// 代役有ジョーカー ↔ 代役有ジョーカー：代役スート・代役数値を比較
	TrumpCardCompareMethodAllowJokerAltStrict

	// TrumpCardCompareMethodAllowJokerAlt ジョーカーは代役情報がある場合のみスート・数値を比較し、ない場合は常に同じスート・数値として比較する
	// 代役無ジョーカー ↔ 通常カード：常に同じカード
	// 代役有ジョーカー ↔ 通常カード：代役スート・代役数値を比較
	// 代役無ジョーカー ↔ 代役有ジョーカー：常に同じカード
	// 代役有ジョーカー ↔ 代役有ジョーカー：代役スート・代役数値を比較
	TrumpCardCompareMethodAllowJokerAlt

	// TrumpCardCompareMethodAllowJoker ジョーカーは代役情報に関係なく常に同じスート・数値として比較する
	// 代役無ジョーカー ↔ 通常カード：常に同じカード
	// 代役有ジョーカー ↔ 通常カード：常に同じカード
	// 代役無ジョーカー ↔ 代役有ジョーカー：常に同じカード
	// 代役有ジョーカー ↔ 代役有ジョーカー：常に同じカード
	TrumpCardCompareMethodAllowJoker
)

// String returns the string representation of TrumpCardCompareMethod
func (m TrumpCardCompareMethod) String() string {
	switch m {
	case TrumpCardCompareMethodStrict:
		return "Strict"
	case TrumpCardCompareMethodAllowJokerAltStrict:
		return "AllowJokerAltStrict"
	case TrumpCardCompareMethodAllowJokerAlt:
		return "AllowJokerAlt"
	case TrumpCardCompareMethodAllowJoker:
		return "AllowJoker"
	default:
		return "Unknown"
	}
}
