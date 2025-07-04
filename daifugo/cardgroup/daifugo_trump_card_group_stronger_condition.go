package cardgroup

import (
	"github.com/rpaka/umefugo-server/common"
)

// DaifugoTrumpCardGroupStrongerCondition より強いカードグループ条件
type DaifugoTrumpCardGroupStrongerCondition struct {
	field     *DaifugoTrumpCardGroup
	hierarchy TrumpCardHierarchy
}

// NewDaifugoTrumpCardGroupStrongerCondition creates a new stronger condition
func NewDaifugoTrumpCardGroupStrongerCondition(field *DaifugoTrumpCardGroup, hierarchy TrumpCardHierarchy) *DaifugoTrumpCardGroupStrongerCondition {
	return &DaifugoTrumpCardGroupStrongerCondition{
		field:     field,
		hierarchy: hierarchy,
	}
}

// Match 対象のカードグループが場のカードより強いかどうかを判定する
func (c *DaifugoTrumpCardGroupStrongerCondition) Match(target *DaifugoTrumpCardGroup) bool {
	method := common.TrumpCardCompareMethodAllowJokerAlt
	fieldWeakest := c.field.WeakestCard(method, c.hierarchy)
	if fieldWeakest.IsNA() {
		return true
	}
	targetWeakest := target.WeakestCard(method, c.hierarchy)
	if targetWeakest.IsNA() {
		return false
	}
	return c.hierarchy.StrongerThanCard(method, targetWeakest, fieldWeakest)
}
