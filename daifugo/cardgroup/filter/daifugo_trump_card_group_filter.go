package filter

import (
	"github.com/rpaka/umefugo-server/common"
	"github.com/rpaka/umefugo-server/daifugo/cardgroup"
)

// DaifugoTrumpCardGroupFilter カードグループフィルターのインターフェース
type DaifugoTrumpCardGroupFilter interface {
	// GetFilteredCardsWithNoCondition 条件なしでフィルターされたカードを取得する
	GetFilteredCardsWithNoCondition() *cardgroup.DaifugoTrumpCardGroup
	// GetFilteredCardWithCondition 条件ありでフィルターされたカードを取得する
	GetFilteredCardWithCondition(numCards int) *cardgroup.DaifugoTrumpCardGroup
}

// DaifugoTrumpCardGroupFilterBase カードグループフィルターの基本実装
type DaifugoTrumpCardGroupFilterBase struct {
	// SourceCards 元カード群
	SourceCards *cardgroup.DaifugoTrumpCardGroup
	// RequiredCards 選択済み（後付け）カード群
	RequiredCards *cardgroup.DaifugoTrumpCardGroup
	// Conditions フィルター条件のリスト
	Conditions []DaifugoTrumpCardGroupFilterCondition
}

// NewDaifugoTrumpCardGroupFilterBase creates a new filter base
func NewDaifugoTrumpCardGroupFilterBase(sourceCards, requiredCards *cardgroup.DaifugoTrumpCardGroup) *DaifugoTrumpCardGroupFilterBase {
	return &DaifugoTrumpCardGroupFilterBase{
		SourceCards:   sourceCards,
		RequiredCards: requiredCards,
		Conditions:    make([]DaifugoTrumpCardGroupFilterCondition, 0),
	}
}

// SourceJokerCards 元カード群のジョーカーカードを取得する
func (f *DaifugoTrumpCardGroupFilterBase) SourceJokerCards() []*common.TrumpCard {
	return f.SourceCards.JokerCards()
}

// ResetAllConditions 全条件をリセットする
func (f *DaifugoTrumpCardGroupFilterBase) ResetAllConditions() {
	for _, condition := range f.Conditions {
		condition.Reset()
	}
}

// AllConditionsSatisfied 全条件が満たされているかを判定する
func (f *DaifugoTrumpCardGroupFilterBase) AllConditionsSatisfied() bool {
	for _, condition := range f.Conditions {
		if condition.Status() != DaifugoTrumpCardGroupFilterConditionStatusSatisfied {
			return false
		}
	}
	return true
}
