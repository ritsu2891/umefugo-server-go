package daifugo

import (
	"testing"

	"github.com/rpaka/umefugo-server/common"
)

func TestDaifugoPlayerExchangeStatus(t *testing.T) {
	status := DaifugoPlayerExchangeStatusNoExchange
	if status.String() != "交換対象外" {
		t.Errorf("Expected '交換対象外', got %s", status.String())
	}
}

func TestDaifugoPlayerGrade(t *testing.T) {
	grade := DaifugoPlayerGradeGreatMillionaire
	if grade.String() != "大富豪" {
		t.Errorf("Expected '大富豪', got %s", grade.String())
	}
}

func TestDaifugoPlayerPassStatus(t *testing.T) {
	status := DaifugoPlayerPassStatusNone
	if status.String() != "未パス" {
		t.Errorf("Expected '未パス', got %s", status.String())
	}
}

func TestDaifugoPlayer(t *testing.T) {
	player := NewDaifugoPlayer("テストプレイヤー")

	if player.Name() != "テストプレイヤー" {
		t.Errorf("Expected 'テストプレイヤー', got %s", player.Name())
	}

	if player.PassStatus() != DaifugoPlayerPassStatusNone {
		t.Errorf("Expected %v, got %v", DaifugoPlayerPassStatusNone, player.PassStatus())
	}

	player.SetPassStatus(DaifugoPlayerPassStatusPassed)
	if player.PassStatus() != DaifugoPlayerPassStatusPassed {
		t.Errorf("Expected %v, got %v", DaifugoPlayerPassStatusPassed, player.PassStatus())
	}
}

func TestDaifugoTrumpCardHierarchy(t *testing.T) {
	hierarchy := NewDaifugoTrumpCardHierarchy()

	// 初期状態の確認
	if hierarchy.WeakestNumber() != DefaultWeakestNumber {
		t.Errorf("Expected %d, got %d", DefaultWeakestNumber, hierarchy.WeakestNumber())
	}

	if hierarchy.StrongestNumber() != DefaultStrongestNumber {
		t.Errorf("Expected %d, got %d", DefaultStrongestNumber, hierarchy.StrongestNumber())
	}

	if hierarchy.IsRevolution() {
		t.Error("Expected normal state, but got revolution state")
	}

	// 革命状態のテスト
	hierarchy.ActivateRevolution()
	if !hierarchy.IsRevolution() {
		t.Error("Expected revolution state, but got normal state")
	}

	hierarchy.DeactivateRevolution()
	if hierarchy.IsRevolution() {
		t.Error("Expected normal state, but got revolution state")
	}

	// 序列のテスト
	order3 := hierarchy.StrongOrder(3)
	order2 := hierarchy.StrongOrder(2)
	if order3 >= order2 {
		t.Errorf("Expected 3 to be weaker than 2, but got order3=%d, order2=%d", order3, order2)
	}

	// カード比較のテスト
	card3, _ := common.NewNormalTrumpCard(common.TrumpCardSuitSpade, 3)
	card2, _ := common.NewNormalTrumpCard(common.TrumpCardSuitSpade, 2)

	if !hierarchy.StrongerThanCard(common.TrumpCardCompareMethodStrict, card2, card3) {
		t.Error("Expected card2 to be stronger than card3")
	}

	if hierarchy.StrongerThanCard(common.TrumpCardCompareMethodStrict, card3, card2) {
		t.Error("Expected card3 to be weaker than card2")
	}
}

func TestDaifugoTrumpCardHierarchyNumbers(t *testing.T) {
	hierarchy := NewDaifugoTrumpCardHierarchy()
	numbers := hierarchy.Numbers()

	// デフォルトでは3から始まり2で終わる
	if len(numbers) != common.NumbersCount {
		t.Errorf("Expected %d numbers, got %d", common.NumbersCount, len(numbers))
	}

	if numbers[0] != 3 {
		t.Errorf("Expected first number to be 3, got %d", numbers[0])
	}

	if numbers[len(numbers)-1] != 2 {
		t.Errorf("Expected last number to be 2, got %d", numbers[len(numbers)-1])
	}
}
