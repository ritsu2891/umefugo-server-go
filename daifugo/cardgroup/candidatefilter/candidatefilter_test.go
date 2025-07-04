package candidatefilter

import (
	"testing"

	"github.com/rpaka/umefugo-server/common"
)

func TestDaifugoTrumpCardGroupCandidateFilterSingle_GetCandidateGroups(t *testing.T) {
	// テスト用のカードを作成
	card1, _ := common.NewNormalTrumpCard(common.TrumpCardSuitSpade, 3)
	card2, _ := common.NewNormalTrumpCard(common.TrumpCardSuitHeart, 4)

	cards := []*common.TrumpCard{card1, card2}

	filter := NewDaifugoTrumpCardGroupCandidateFilterSingle()
	candidateGroups := filter.GetCandidateGroups(cards)

	// 2枚のカードから2つの候補グループが生成されるはず
	if len(candidateGroups) != 2 {
		t.Errorf("Expected 2 candidate groups, but got %d", len(candidateGroups))
	}

	// 各候補グループは1つの位置を持ち、その位置に1枚のカードがあるはず
	for i, group := range candidateGroups {
		if len(group) != 1 {
			t.Errorf("Expected 1 position in group %d, but got %d", i, len(group))
		}
		if len(group[0]) != 1 {
			t.Errorf("Expected 1 card in position 0 of group %d, but got %d", i, len(group[0]))
		}
	}
}

func TestDaifugoTrumpCardGroupCandidateFilterPair_GetCandidateGroups(t *testing.T) {
	// テスト用のカードを作成（3のペア + ジョーカー）
	card1, _ := common.NewNormalTrumpCard(common.TrumpCardSuitSpade, 3)
	card2, _ := common.NewNormalTrumpCard(common.TrumpCardSuitHeart, 3)
	card3 := common.NewJokerTrumpCard()

	cards := []*common.TrumpCard{card1, card2, card3}

	filter := NewDaifugoTrumpCardGroupCandidateFilterPair()
	candidateGroups := filter.GetCandidateGroups(cards)

	// 3のペアが構成可能なので、最低1つの候補グループが生成されるはず
	if len(candidateGroups) == 0 {
		t.Error("Expected at least 1 candidate group, but got none")
	}

	// 各候補グループは2つの位置を持つはず（ペアなので）
	for i, group := range candidateGroups {
		if len(group) != 3 { // carNumCards.Count で位置数が決まる
			t.Errorf("Expected 3 positions in group %d (3 cards available), but got %d", i, len(group))
		}
	}
}

func TestConditionSatisfyStatus(t *testing.T) {
	// 基本的なenum値の確認
	if ConditionSatisfyStatusNeutral != 0 {
		t.Error("Expected Neutral status to be 0")
	}
	if ConditionSatisfyStatusSatisfy != 1 {
		t.Error("Expected Satisfy status to be 1")
	}
	if ConditionSatisfyStatusUnsatisfy != 2 {
		t.Error("Expected Unsatisfy status to be 2")
	}
}

func TestDaifugoTrumpCardGroupCandidateFilterConditionBase(t *testing.T) {
	condition := NewDaifugoTrumpCardGroupCandidateFilterConditionBase()

	// 初期状態はNeutral
	if condition.Status() != ConditionSatisfyStatusNeutral {
		t.Error("Expected initial status to be Neutral")
	}

	// 状態変更のテスト
	condition.SetStatus(ConditionSatisfyStatusSatisfy)
	if condition.Status() != ConditionSatisfyStatusSatisfy {
		t.Error("Expected status to be Satisfy after setting")
	}

	// リセットのテスト
	condition.Reset()
	if condition.Status() != ConditionSatisfyStatusNeutral {
		t.Error("Expected status to be Neutral after reset")
	}
}
