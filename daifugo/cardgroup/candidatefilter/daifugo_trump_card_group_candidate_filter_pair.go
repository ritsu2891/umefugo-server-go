package candidatefilter

import (
	"github.com/rpaka/umefugo-server/common"
	"github.com/rpaka/umefugo-server/daifugo/cardgroup"
)

// DaifugoTrumpCardGroupCandidateFilterPair ペアのカードグループ候補抽出フィルタ
// 手札のジョーカーには代役情報はついていないものと考え、代役情報は無視する
type DaifugoTrumpCardGroupCandidateFilterPair struct {
	*DaifugoTrumpCardGroupCandidateFilterBase
}

// NewDaifugoTrumpCardGroupCandidateFilterPair creates a new pair candidate filter
func NewDaifugoTrumpCardGroupCandidateFilterPair() *DaifugoTrumpCardGroupCandidateFilterPair {
	return &DaifugoTrumpCardGroupCandidateFilterPair{
		DaifugoTrumpCardGroupCandidateFilterBase: NewDaifugoTrumpCardGroupCandidateFilterBase(),
	}
}

// GetCandidateGroups カードグループ候補を取得する
func (f *DaifugoTrumpCardGroupCandidateFilterPair) GetCandidateGroups(cards []*common.TrumpCard) [][][]*common.TrumpCard {
	var res [][][]*common.TrumpCard

	// ジョーカーと各数字のカードを抽出
	dcards := cardgroup.NewDaifugoTrumpCardGroupWithCards(cards)
	jokers := dcards.JokerCards()
	numDict := dcards.SeparateByNumber()

	// 各数字についてペアとなるカード候補を追加
	for _, n := range common.Numbers {
		// 枚数が不足していたらこの数字についてはスキップ
		if len(jokers)+len(numDict[n]) < 2 {
			continue
		}

		// ペアを構成するカードを抽出
		var carNumCards []*common.TrumpCard
		carNumCards = append(carNumCards, numDict[n]...)
		carNumCards = append(carNumCards, jokers...)

		// ペアを構成するカードの位置は任意性があるため全ての位置で候補として追加
		var carNumRes [][]*common.TrumpCard
		k := len(carNumCards)
		for k > 0 {
			// 各位置でカード候補のコピーを追加
			carNumCardsCopy := make([]*common.TrumpCard, len(carNumCards))
			copy(carNumCardsCopy, carNumCards)
			carNumRes = append(carNumRes, carNumCardsCopy)
			k--
		}

		// 現在走査中の数字の候補として追加
		res = append(res, carNumRes)
	}

	return res
}
