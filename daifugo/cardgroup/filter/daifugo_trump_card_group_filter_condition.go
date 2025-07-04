package filter

import (
	"github.com/rpaka/umefugo-server/common"
)

// DaifugoTrumpCardGroupFilterCondition カードグループフィルター条件のインターフェース
type DaifugoTrumpCardGroupFilterCondition interface {
	// CheckIndexOf 指定されたインデックス位置でのカードのチェックを行う
	CheckIndexOf(index int, cards []*common.TrumpCard) []*common.TrumpCard
	// Status 現在の充足状況を取得する
	Status() DaifugoTrumpCardGroupFilterConditionStatus
	// SetStatus 充足状況を設定する
	SetStatus(status DaifugoTrumpCardGroupFilterConditionStatus)
	// Reset 状況をリセットする
	Reset()
}
