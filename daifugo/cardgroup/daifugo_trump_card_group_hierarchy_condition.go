package cardgroup

import (
	"github.com/rpaka/umefugo-server/common"
)

// DaifugoTrumpCardGroupHierarchyCondition 数字縛りカードグループ条件
type DaifugoTrumpCardGroupHierarchyCondition struct {
	requiredHierarchyDistance int
	field                     *DaifugoTrumpCardGroup
	hierarchy                 TrumpCardHierarchy
}

// NewDaifugoTrumpCardGroupHierarchyCondition creates a new hierarchy condition
func NewDaifugoTrumpCardGroupHierarchyCondition(field *DaifugoTrumpCardGroup, hierarchy TrumpCardHierarchy) *DaifugoTrumpCardGroupHierarchyCondition {
	return &DaifugoTrumpCardGroupHierarchyCondition{
		requiredHierarchyDistance: 1,
		field:                     field,
		hierarchy:                 hierarchy,
	}
}

// Match 対象のカードグループが数字縛り条件を満たすかどうかを判定する
func (c *DaifugoTrumpCardGroupHierarchyCondition) Match(target *DaifugoTrumpCardGroup) bool {
	// 確認するカードグループの中に代役数値情報がないジョーカーを含んでいる場合、条件の判定ができない
	// 確認処理としてはfalseを返すが、ログに残す
	if target.HasNoAltNumberJoker() {
		// TODO: ログ実装
		return false
	}

	method := common.TrumpCardCompareMethodAllowJokerAlt
	distance := c.field.StrongOrderDistance(method, c.hierarchy, target)

	// 場のカードグループの強さ（最弱カードの強さ）と対象のカードグループの強さ（最弱カードの強さ）の差が縛りで指定されている強さの差と一致する場合は条件を満たす
	if distance == c.requiredHierarchyDistance {
		return true
	}

	// そうでなくとも、場のカードグループの強さ（最弱カードの強さ）よりも対象のカードグループの強さが上回っており
	// かつ対象のカードグループに代役情報無しジョーカーが存在している場合は条件を満たす（ジョーカーが指定の強さのカードと見なす）
	if distance > 0 && target.HasNoAltJoker() {
		return true
	}

	return false
}
