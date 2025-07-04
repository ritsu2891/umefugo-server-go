package cardgroup

import (
	"github.com/rpaka/umefugo-server/common"
)

// DaifugoTrumpCardGroupTypePairCondition ペアカードグループ条件
type DaifugoTrumpCardGroupTypePairCondition struct{}

// NewDaifugoTrumpCardGroupTypePairCondition creates a new pair condition
func NewDaifugoTrumpCardGroupTypePairCondition() *DaifugoTrumpCardGroupTypePairCondition {
	return &DaifugoTrumpCardGroupTypePairCondition{}
}

// Match 対象のカードグループがペア条件を満たすかどうかを判定する
func (c *DaifugoTrumpCardGroupTypePairCondition) Match(target *DaifugoTrumpCardGroup) bool {
	return target.Count() > 1 && target.IsSameNumbers(common.TrumpCardCompareMethodAllowJoker)
}
