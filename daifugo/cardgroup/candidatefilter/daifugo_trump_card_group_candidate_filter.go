package candidatefilter

import (
	"github.com/rpaka/umefugo-server/common"
)

// DaifugoTrumpCardGroupCandidateFilter カードグループ候補のカードを抽出するインターフェース
// 主にクライアントサイドで利用
type DaifugoTrumpCardGroupCandidateFilter interface {
	// Filter 条件なしでフィルターを実行する
	Filter(cards []*common.TrumpCard) []*common.TrumpCard
	// FilterWithConditions 条件ありでフィルターを実行する
	FilterWithConditions(conditions []DaifugoTrumpCardGroupCandidateFilterCondition, cards []*common.TrumpCard) []*common.TrumpCard
	// GetCandidateGroups カードグループ候補を取得する
	GetCandidateGroups(cards []*common.TrumpCard) [][][]*common.TrumpCard
}

// DaifugoTrumpCardGroupCandidateFilterBase 候補フィルターの基本実装
type DaifugoTrumpCardGroupCandidateFilterBase struct{}

// NewDaifugoTrumpCardGroupCandidateFilterBase creates a new candidate filter base
func NewDaifugoTrumpCardGroupCandidateFilterBase() *DaifugoTrumpCardGroupCandidateFilterBase {
	return &DaifugoTrumpCardGroupCandidateFilterBase{}
}

// Filter 条件なしでフィルターを実行する
func (f *DaifugoTrumpCardGroupCandidateFilterBase) Filter(cards []*common.TrumpCard) []*common.TrumpCard {
	return f.FilterWithConditions([]DaifugoTrumpCardGroupCandidateFilterCondition{}, cards)
}

// FilterWithConditions 条件ありでフィルターを実行する
func (f *DaifugoTrumpCardGroupCandidateFilterBase) FilterWithConditions(conditions []DaifugoTrumpCardGroupCandidateFilterCondition, cards []*common.TrumpCard) []*common.TrumpCard {
	candidateGroups := f.GetCandidateGroups(cards)
	resSet := make(map[*common.TrumpCard]bool)

	// 各カードグループ候補について調べる
	for _, candidateGroup := range candidateGroups {
		// カードグループ候補内の各位置の条件を満たすカード
		var filteredCandidateGroup [][]*common.TrumpCard
		unsatisfied := false

		// 縛り条件判定リセット
		for _, condition := range conditions {
			condition.Reset()
		}

		// カードグループ候補内の各位置について調べる
		for idx := 0; idx < len(candidateGroup); idx++ {
			// 現在の位置のカード全て（ワーク）
			curIndexCards := make([]*common.TrumpCard, len(candidateGroup[idx]))
			copy(curIndexCards, candidateGroup[idx])

			// 各縛り条件について調べる
			for _, condition := range conditions {
				curIndexCards = condition.Scan(idx, curIndexCards)

				// 現在の位置・縛り条件で既に条件を満たさない事が分かったらこれ以上調べる必要は無い
				if condition.Status() == ConditionSatisfyStatusUnsatisfy {
					unsatisfied = true
					break
				}
			}

			// 条件を満たしていない場合はこのカードグループ候補については終了
			if unsatisfied {
				break
			}

			filteredCandidateGroup = append(filteredCandidateGroup, curIndexCards)
		}

		// 全ての条件を満たしていることを確認する
		if len(conditions) != 0 {
			for _, cond := range conditions {
				if cond.Status() != ConditionSatisfyStatusSatisfy {
					unsatisfied = true
					break
				}
			}
		}

		// 条件を満たしていない場合は候補として追加しない
		if unsatisfied {
			continue
		}

		// 結果に追加
		for _, cardsInGroup := range filteredCandidateGroup {
			for _, card := range cardsInGroup {
				resSet[card] = true
			}
		}
	}

	// setから配列に変換して重複を排除して返す
	var result []*common.TrumpCard
	for card := range resSet {
		result = append(result, card)
	}

	return result
}

// GetCandidateGroups カードグループ候補を取得する（基本実装では空のスライスを返す）
func (f *DaifugoTrumpCardGroupCandidateFilterBase) GetCandidateGroups(cards []*common.TrumpCard) [][][]*common.TrumpCard {
	// 具象クラスでオーバーライドする必要がある
	return [][][]*common.TrumpCard{}
}
