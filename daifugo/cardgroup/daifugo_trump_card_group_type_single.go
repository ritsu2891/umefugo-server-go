package cardgroup

import (
	"github.com/rpaka/umefugo-server/common"
)

// DaifugoTrumpCardGroupTypeSingle 大富豪トランプカードシングルグループ構成
type DaifugoTrumpCardGroupTypeSingle struct {
	class DaifugoTrumpCardGroupTypeClass
}

// NewDaifugoTrumpCardGroupTypeSingle creates a new single group type
func NewDaifugoTrumpCardGroupTypeSingle() *DaifugoTrumpCardGroupTypeSingle {
	return &DaifugoTrumpCardGroupTypeSingle{
		class: DaifugoTrumpCardGroupTypeClassSingle,
	}
}

// Class returns the group type class
func (t *DaifugoTrumpCardGroupTypeSingle) Class() DaifugoTrumpCardGroupTypeClass {
	return t.class
}

// ConsistableCards 手札のうち構成可能なカードを取得する
func (t *DaifugoTrumpCardGroupTypeSingle) ConsistableCards(fieldCards, handCards, handSelectedCards *DaifugoTrumpCardGroup) []*common.TrumpCard {
	condition := NewDaifugoTrumpCardGroupTypeSingleCondition()

	// 場または手札で既に選択済みのカードがシングルの条件を満たさない場合はシングルの候補はない
	if !condition.Match(fieldCards) || !condition.Match(handSelectedCards) {
		return []*common.TrumpCard{}
	}

	// 手札で既に選択されているカードがあれば、そのカードのみを返す
	if handSelectedCards.Count() > 0 {
		return handSelectedCards.ToSlice()
	}

	// 手札の全カードがシングルの候補
	return handCards.ToSlice()
}
