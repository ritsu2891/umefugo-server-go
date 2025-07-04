package filter

import (
	"github.com/rpaka/umefugo-server/common"
	"github.com/rpaka/umefugo-server/daifugo/cardgroup"
)

// DaifugoTrumpCardGroupFilterPair ペア型カードグループフィルター
type DaifugoTrumpCardGroupFilterPair struct {
	*DaifugoTrumpCardGroupFilterBase
}

// 最低構成枚数
const consistNumCardMin = 2

// NewDaifugoTrumpCardGroupFilterPair creates a new pair filter
func NewDaifugoTrumpCardGroupFilterPair(sourceCards, requiredCards *cardgroup.DaifugoTrumpCardGroup) *DaifugoTrumpCardGroupFilterPair {
	return &DaifugoTrumpCardGroupFilterPair{
		DaifugoTrumpCardGroupFilterBase: NewDaifugoTrumpCardGroupFilterBase(sourceCards, requiredCards),
	}
}

// GetFilteredCardsWithNoCondition 条件なしでフィルターされたカードを取得する
func (f *DaifugoTrumpCardGroupFilterPair) GetFilteredCardsWithNoCondition() *cardgroup.DaifugoTrumpCardGroup {
	return f.getFilteredCard(true, 0)
}

// GetFilteredCardWithCondition 条件ありでフィルターされたカードを取得する
func (f *DaifugoTrumpCardGroupFilterPair) GetFilteredCardWithCondition(numCards int) *cardgroup.DaifugoTrumpCardGroup {
	return f.getFilteredCard(false, numCards)
}

// getFilteredCard 条件あり・なし共通のフィルター処理
func (f *DaifugoTrumpCardGroupFilterPair) getFilteredCard(noCond bool, numCards int) *cardgroup.DaifugoTrumpCardGroup {
	resSet := make(map[*common.TrumpCard]bool)

	// 対象札全てを数字毎に仕分ける
	numDic := f.SourceCards.SeparateByNumber()
	sourceJokerCards := f.SourceJokerCards()

	// 各数字について
	for _, num := range common.Numbers {
		if noCond {
			// 最低構成枚数を満たすかを調べて結果に加える
			if len(numDic[num])+len(sourceJokerCards) < consistNumCardMin {
				continue
			}
			for _, card := range numDic[num] {
				resSet[card] = true
			}
		} else {
			// 現在の数字のカードとジョーカーのカードを現在の数字の候補として用意
			var curNumCards []*common.TrumpCard
			curNumCards = append(curNumCards, numDic[num]...)
			curNumCards = append(curNumCards, sourceJokerCards...)
			curNumContinue := true
			curNumFilteredCardsSet := make(map[*common.TrumpCard]bool)

			// 各条件リセット
			f.ResetAllConditions()

			// そもそもこのカードグループを構成可能なカードが不足
			// => グループ構成不可
			if len(curNumCards) < consistNumCardMin {
				continue
			}

			// 各位置
			// 条件を満たすか調べ、候補カードを絞り込む
			for idx := 0; idx < numCards; idx++ {
				var curNumCurIndexCards []*common.TrumpCard
				curNumCurIndexCards = append(curNumCurIndexCards, curNumCards...)

				// 各条件
				for _, cond := range f.Conditions {
					curNumCurIndexCards = cond.CheckIndexOf(idx, curNumCurIndexCards)
					curNumContinue = cond.Status() != DaifugoTrumpCardGroupFilterConditionStatusUnSatisfied

					// 条件確認時点でもう条件成立しないことが分かった
					// => 現在の数字についてはグループ構成不可
					if !curNumContinue {
						break
					}
				}
				if !curNumContinue {
					continue
				}

				// 現在の位置の条件を満たすカードを
				// 現在の数字の条件を満たすカードの候補として追加
				for _, card := range curNumCurIndexCards {
					curNumFilteredCardsSet[card] = true
				}
			}
			if !curNumContinue {
				continue
			}

			// 条件が全て満たされていることを確認
			if !f.AllConditionsSatisfied() {
				continue
			}

			// 現在の数字の条件を満たすカードを
			// 結果として追加
			for card := range curNumFilteredCardsSet {
				resSet[card] = true
			}
		}
	}

	// 条件なしの場合
	// ある数字で構成可能か、あるいはジョーカー単体で最低構成枚数を満たすのであれば
	// ジョーカーも結果に加える
	if noCond && (consistNumCardMin <= len(sourceJokerCards) || len(resSet) > 0) {
		for _, card := range sourceJokerCards {
			resSet[card] = true
		}
	}

	// setから配列に変換
	var result []*common.TrumpCard
	for card := range resSet {
		result = append(result, card)
	}

	return cardgroup.NewDaifugoTrumpCardGroupWithCards(result)
}
