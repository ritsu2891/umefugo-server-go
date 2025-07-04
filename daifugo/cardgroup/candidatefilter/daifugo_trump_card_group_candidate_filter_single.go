package candidatefilter

import (
	"github.com/rpaka/umefugo-server/common"
)

// DaifugoTrumpCardGroupCandidateFilterSingle カード1枚のカードグループ候補抽出フィルタ
type DaifugoTrumpCardGroupCandidateFilterSingle struct {
	*DaifugoTrumpCardGroupCandidateFilterBase
}

// NewDaifugoTrumpCardGroupCandidateFilterSingle creates a new single candidate filter
func NewDaifugoTrumpCardGroupCandidateFilterSingle() *DaifugoTrumpCardGroupCandidateFilterSingle {
	return &DaifugoTrumpCardGroupCandidateFilterSingle{
		DaifugoTrumpCardGroupCandidateFilterBase: NewDaifugoTrumpCardGroupCandidateFilterBase(),
	}
}

// GetCandidateGroups カードグループ候補を取得する
func (f *DaifugoTrumpCardGroupCandidateFilterSingle) GetCandidateGroups(cards []*common.TrumpCard) [][][]*common.TrumpCard {
	var res [][][]*common.TrumpCard
	for _, card := range cards {
		candidateGroup := [][]*common.TrumpCard{{card}}
		res = append(res, candidateGroup)
	}
	return res
}
