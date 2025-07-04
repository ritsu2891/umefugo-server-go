package cardgroup

import (
	"github.com/rpaka/umefugo-server/common"
)

// DaifugoTrumpCardGroupTypePair 大富豪トランプカードペアグループ構成
type DaifugoTrumpCardGroupTypePair struct {
	class DaifugoTrumpCardGroupTypeClass
}

const (
	// DefaultMinConsistCountPair デフォルトの構成可能最低枚数（ペア）
	DefaultMinConsistCountPair = 2
)

// NewDaifugoTrumpCardGroupTypePair creates a new pair group type
func NewDaifugoTrumpCardGroupTypePair() *DaifugoTrumpCardGroupTypePair {
	return &DaifugoTrumpCardGroupTypePair{
		class: DaifugoTrumpCardGroupTypeClassPair,
	}
}

// Class returns the group type class
func (t *DaifugoTrumpCardGroupTypePair) Class() DaifugoTrumpCardGroupTypeClass {
	return t.class
}

// ConsistableCards 手札のうち構成可能なカードを取得する
func (t *DaifugoTrumpCardGroupTypePair) ConsistableCards(fieldCards, handCards, handSelectedCards *DaifugoTrumpCardGroup) []*common.TrumpCard {
	result := make(map[*common.TrumpCard]bool)

	condition := NewDaifugoTrumpCardGroupTypePairCondition()
	handCardNumbers := handCards.Numbers()
	handJokers := handCards.JokerCards()
	handNumberDic := handCards.SeparateByNumber()

	// 場または手札で既に選択済みのカードがペアの条件を満たさない場合はペアの候補はない
	if !condition.Match(fieldCards) || !condition.Match(handSelectedCards) {
		return []*common.TrumpCard{}
	}

	// 最低構成枚数は場にカードがなければ2枚、あれば場のカードの枚数となる
	minConsistCount := DefaultMinConsistCountPair
	if fieldCards.Count() > 0 {
		minConsistCount = fieldCards.Count()
	}

	// 手札で既に選択されているカードがあれば、その数字のみを候補抽出の対象とする
	searchNumbers := common.Numbers
	if handSelectedCards.Count() > 0 {
		searchNumbers = handCardNumbers
	}

	// 各数字で候補を抽出
	for _, number := range searchNumbers {
		// TODO: 縛りに有効なカードを問い合わせ、縛り条件成立を確認

		// 当該数字+ジョーカーの枚数が最低構成枚数を満たしていない場合はペアの候補でない
		numberCards := handNumberDic[number]
		if len(numberCards)+len(handJokers) < minConsistCount {
			continue
		}

		// 当該数字のカード・およびジョーカーを候補として追加
		for _, card := range numberCards {
			result[card] = true
		}
		for _, joker := range handJokers {
			result[joker] = true
		}
	}

	// map から slice に変換
	var resultSlice []*common.TrumpCard
	for card := range result {
		resultSlice = append(resultSlice, card)
	}

	return resultSlice
}
