package cardgroup

import (
	"fmt"
	"strings"

	"github.com/rpaka/umefugo-server/common"
)

// DaifugoTrumpCardGroupSuitCondition スート縛りカードグループ条件
type DaifugoTrumpCardGroupSuitCondition struct {
	requiredSuitCounts map[common.TrumpCardSuit]int
}

// NewDaifugoTrumpCardGroupSuitCondition creates a new suit condition
func NewDaifugoTrumpCardGroupSuitCondition(requiredSuitCounts map[common.TrumpCardSuit]int) *DaifugoTrumpCardGroupSuitCondition {
	return &DaifugoTrumpCardGroupSuitCondition{
		requiredSuitCounts: requiredSuitCounts,
	}
}

// String returns the string representation of the condition
func (c *DaifugoTrumpCardGroupSuitCondition) String() string {
	var labels []string
	for _, suit := range common.Suits {
		if count, exists := c.requiredSuitCounts[suit]; exists {
			labels = append(labels, fmt.Sprintf("%s:%d", suit.ToMarkString(), count))
		}
	}
	return fmt.Sprintf("スート縛り %s", strings.Join(labels, " "))
}

// Match 対象のカードグループがスート縛り条件を満たすかどうかを判定する
func (c *DaifugoTrumpCardGroupSuitCondition) Match(target *DaifugoTrumpCardGroup) bool {
	suitCountDic := target.CountBySuit()
	jokerCount := target.CountJoker()
	remCount := 0

	for _, suit := range common.Suits {
		requiredCount, requiredExists := c.requiredSuitCounts[suit]
		if !requiredExists {
			continue
		}

		targetCount, targetExists := suitCountDic[suit]
		if !targetExists {
			targetCount = 0
		}

		// 縛りで指定されている各スートの枚数よりも対象のカードグループの当該スートの枚数が多い場合は条件を満たさない
		if requiredCount < targetCount {
			return false
		}

		// 縛りで指定されている各スートの枚数よりも対象のカードグループの当該スートの枚数が少ない場合は、少ない分をジョーカーで補うべき枚数としてカウント
		remCount += requiredCount - targetCount
	}

	// ジョーカーで補うべき枚数が1枚以上の場合、ジョーカーの枚数が足りない場合は条件を満たさない
	if remCount > 0 && jokerCount < remCount {
		return false
	}

	return true
}
