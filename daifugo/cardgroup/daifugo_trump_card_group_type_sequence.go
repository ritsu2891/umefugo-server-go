package cardgroup

import (
	"github.com/rpaka/umefugo-server/common"
)

// DaifugoTrumpCardGroupTypeSequence 大富豪トランプカード階段グループ構成
type DaifugoTrumpCardGroupTypeSequence struct {
	class              DaifugoTrumpCardGroupTypeClass
	minConsistCount    int
	maxConsistCount    int
	allowDifferentSuit bool
	hierarchy          TrumpCardHierarchy
}

const (
	// DefaultMinConsistCountSequence デフォルトの構成可能最低枚数（階段）
	DefaultMinConsistCountSequence = 3
	// DefaultMaxConsistCountSequence デフォルトの構成可能最大枚数（階段）
	DefaultMaxConsistCountSequence = common.NumbersCount
)

// NewDaifugoTrumpCardGroupTypeSequence creates a new sequence group type
func NewDaifugoTrumpCardGroupTypeSequence(hierarchy TrumpCardHierarchy) *DaifugoTrumpCardGroupTypeSequence {
	return &DaifugoTrumpCardGroupTypeSequence{
		class:              DaifugoTrumpCardGroupTypeClassSequence,
		minConsistCount:    DefaultMinConsistCountSequence,
		maxConsistCount:    DefaultMaxConsistCountSequence,
		allowDifferentSuit: false,
		hierarchy:          hierarchy,
	}
}

// Class returns the group type class
func (t *DaifugoTrumpCardGroupTypeSequence) Class() DaifugoTrumpCardGroupTypeClass {
	return t.class
}

// MinConsistCount returns the minimum consist count
func (t *DaifugoTrumpCardGroupTypeSequence) MinConsistCount() int {
	return t.minConsistCount
}

// SetMinConsistCount sets the minimum consist count
func (t *DaifugoTrumpCardGroupTypeSequence) SetMinConsistCount(count int) {
	t.minConsistCount = count
}

// MaxConsistCount returns the maximum consist count
func (t *DaifugoTrumpCardGroupTypeSequence) MaxConsistCount() int {
	return t.maxConsistCount
}

// SetMaxConsistCount sets the maximum consist count
func (t *DaifugoTrumpCardGroupTypeSequence) SetMaxConsistCount(count int) {
	t.maxConsistCount = count
}

// AllowDifferentSuit returns whether different suits are allowed
func (t *DaifugoTrumpCardGroupTypeSequence) AllowDifferentSuit() bool {
	return t.allowDifferentSuit
}

// SetAllowDifferentSuit sets whether different suits are allowed
func (t *DaifugoTrumpCardGroupTypeSequence) SetAllowDifferentSuit(allow bool) {
	t.allowDifferentSuit = allow
}

// ConsistableCards 手札のうち構成可能なカードを取得する
func (t *DaifugoTrumpCardGroupTypeSequence) ConsistableCards(fieldCards, handCards, handSelectedCards *DaifugoTrumpCardGroup) []*common.TrumpCard {
	result := make(map[*common.TrumpCard]bool)

	// 選択しているカードが既にある場合
	if handSelectedCards.Count() > 0 {
		// 選択しているカードが階段でない場合は階段構成可能カードはない
		if !handSelectedCards.IsSequence(t.hierarchy, false) {
			return []*common.TrumpCard{}
		}

		// 階段で有、かつ既に最大構成枚数に達している場合は選択中のカードのみ返す
		if t.maxConsistCount <= handSelectedCards.Count() {
			return handSelectedCards.ToSlice()
		}
	}

	// 場にカードが出ている場合はその枚数に制限して調べる
	if fieldCards.Count() > 0 {
		t.minConsistCount = fieldCards.Count()
		t.maxConsistCount = fieldCards.Count()
	}

	// TODO: スートの混合を許可しない場合はスート毎に調べる実装
	// 現在は基本的な実装のみ

	// 簡単な実装: 手札の全カードを返す（実際の階段判定は後で実装）
	for _, card := range handCards.ToSlice() {
		result[card] = true
	}

	// map から slice に変換
	var resultSlice []*common.TrumpCard
	for card := range result {
		resultSlice = append(resultSlice, card)
	}

	return resultSlice
}
