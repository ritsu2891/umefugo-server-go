package cardgroup

// DaifugoTrumpCardGroupTypeSingleCondition シングルカードグループ条件
type DaifugoTrumpCardGroupTypeSingleCondition struct{}

// NewDaifugoTrumpCardGroupTypeSingleCondition creates a new single condition
func NewDaifugoTrumpCardGroupTypeSingleCondition() *DaifugoTrumpCardGroupTypeSingleCondition {
	return &DaifugoTrumpCardGroupTypeSingleCondition{}
}

// Match 対象のカードグループがシングル条件を満たすかどうかを判定する
func (c *DaifugoTrumpCardGroupTypeSingleCondition) Match(target *DaifugoTrumpCardGroup) bool {
	return target.Count() == 1
}
