package cardgroup

import (
	"github.com/rpaka/umefugo-server/common"
)

// DaifugoTrumpCardGroupType 大富豪トランプカードグループ構成
type DaifugoTrumpCardGroupType interface {
	// Class returns the class of the group type
	Class() DaifugoTrumpCardGroupTypeClass
	// ConsistableCards 手札のうち構成可能なカードを取得する
	ConsistableCards(fieldCards, handCards, handSelectedCards *DaifugoTrumpCardGroup) []*common.TrumpCard
}
