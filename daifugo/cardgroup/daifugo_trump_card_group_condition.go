package cardgroup

// DaifugoTrumpCardGroupCondition 大富豪トランプカードグループ条件
type DaifugoTrumpCardGroupCondition interface {
	// Match 対象のカードグループが条件を満たすかどうかを判定する
	Match(target *DaifugoTrumpCardGroup) bool
}
