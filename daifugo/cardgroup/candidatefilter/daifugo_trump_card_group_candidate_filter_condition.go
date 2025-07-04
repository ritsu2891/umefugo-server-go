package candidatefilter

import (
	"github.com/rpaka/umefugo-server/common"
)

// ConditionSatisfyStatus 条件満足状況
type ConditionSatisfyStatus int

const (
	// ConditionSatisfyStatusNeutral 中立
	ConditionSatisfyStatusNeutral ConditionSatisfyStatus = iota
	// ConditionSatisfyStatusSatisfy 満足
	ConditionSatisfyStatusSatisfy
	// ConditionSatisfyStatusUnsatisfy 不満足
	ConditionSatisfyStatusUnsatisfy
)

// DaifugoTrumpCardGroupCandidateFilterCondition カードグループ候補のカード抽出の条件インターフェース
// 主にクライアントサイドで利用
type DaifugoTrumpCardGroupCandidateFilterCondition interface {
	// Scan 指定されたインデックス位置でのカードのスキャンを行う
	Scan(index int, cards []*common.TrumpCard) []*common.TrumpCard
	// Status 現在の満足状況を取得する
	Status() ConditionSatisfyStatus
	// SetStatus 満足状況を設定する
	SetStatus(status ConditionSatisfyStatus)
	// Reset 状況をリセットする
	Reset()
}

// DaifugoTrumpCardGroupCandidateFilterConditionBase 候補フィルター条件の基本実装
type DaifugoTrumpCardGroupCandidateFilterConditionBase struct {
	status ConditionSatisfyStatus
}

// NewDaifugoTrumpCardGroupCandidateFilterConditionBase creates a new candidate filter condition base
func NewDaifugoTrumpCardGroupCandidateFilterConditionBase() *DaifugoTrumpCardGroupCandidateFilterConditionBase {
	return &DaifugoTrumpCardGroupCandidateFilterConditionBase{
		status: ConditionSatisfyStatusNeutral,
	}
}

// Status 現在の満足状況を取得する
func (c *DaifugoTrumpCardGroupCandidateFilterConditionBase) Status() ConditionSatisfyStatus {
	return c.status
}

// SetStatus 満足状況を設定する
func (c *DaifugoTrumpCardGroupCandidateFilterConditionBase) SetStatus(status ConditionSatisfyStatus) {
	c.status = status
}

// Reset 状況をリセットする
func (c *DaifugoTrumpCardGroupCandidateFilterConditionBase) Reset() {
	c.status = ConditionSatisfyStatusNeutral
}
