package cardgroup

import "fmt"

// DaifugoTrumpCardGroupTypeClass 大富豪トランプカードグループ種別
type DaifugoTrumpCardGroupTypeClass int

const (
	// DaifugoTrumpCardGroupTypeClassNotAssigned 未割当
	DaifugoTrumpCardGroupTypeClassNotAssigned DaifugoTrumpCardGroupTypeClass = iota
	// DaifugoTrumpCardGroupTypeClassSingle シングル（カード1枚）
	DaifugoTrumpCardGroupTypeClassSingle
	// DaifugoTrumpCardGroupTypeClassPair ペア（同じ数字複数枚）
	DaifugoTrumpCardGroupTypeClassPair
	// DaifugoTrumpCardGroupTypeClassSequence 階段（連続した数字3枚以上）
	DaifugoTrumpCardGroupTypeClassSequence
)

// ToLabelString ラベル文字列の取得
func (t DaifugoTrumpCardGroupTypeClass) ToLabelString() string {
	switch t {
	case DaifugoTrumpCardGroupTypeClassNotAssigned:
		return "未割当"
	case DaifugoTrumpCardGroupTypeClassSingle:
		return "シングル"
	case DaifugoTrumpCardGroupTypeClassPair:
		return "ペア"
	case DaifugoTrumpCardGroupTypeClassSequence:
		return "階段"
	default:
		panic(fmt.Sprintf("未定義のカードグループ種別: %v", t))
	}
}

// String returns the string representation of DaifugoTrumpCardGroupTypeClass
func (t DaifugoTrumpCardGroupTypeClass) String() string {
	return t.ToLabelString()
}
