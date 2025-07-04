package cardgroup

import (
	"testing"

	"github.com/rpaka/umefugo-server/common"
)

func TestDaifugoTrumpCardGroupTypeClass(t *testing.T) {
	typeClass := DaifugoTrumpCardGroupTypeClassSingle
	if typeClass.ToLabelString() != "シングル" {
		t.Errorf("Expected 'シングル', got %s", typeClass.ToLabelString())
	}

	if typeClass.String() != "シングル" {
		t.Errorf("Expected 'シングル', got %s", typeClass.String())
	}

	pairType := DaifugoTrumpCardGroupTypeClassPair
	if pairType.ToLabelString() != "ペア" {
		t.Errorf("Expected 'ペア', got %s", pairType.ToLabelString())
	}
}

func TestDaifugoTrumpCardGroup(t *testing.T) {
	// 空のグループを作成
	group := NewDaifugoTrumpCardGroup()

	if group.Count() != 0 {
		t.Errorf("Expected empty group, got count %d", group.Count())
	}

	// カードを追加
	card1, _ := common.NewNormalTrumpCard(common.TrumpCardSuitSpade, 10)
	card2, _ := common.NewNormalTrumpCard(common.TrumpCardSuitHeart, 10)
	joker := common.NewJokerTrumpCard()

	group.Add(card1)
	group.Add(card2)
	group.Add(joker)

	if group.Count() != 3 {
		t.Errorf("Expected 3 cards, got %d", group.Count())
	}

	// ジョーカーの確認
	if !group.HasJoker() {
		t.Error("Expected to have joker")
	}

	if group.CountJoker() != 1 {
		t.Errorf("Expected 1 joker, got %d", group.CountJoker())
	}

	// 非ジョーカーカードの確認
	notJokers := group.NotJokerCards()
	if len(notJokers) != 2 {
		t.Errorf("Expected 2 non-joker cards, got %d", len(notJokers))
	}

	// スート毎の分離
	suitMap := group.SeparateBySuit()
	if len(suitMap[common.TrumpCardSuitSpade]) != 1 {
		t.Errorf("Expected 1 spade card, got %d", len(suitMap[common.TrumpCardSuitSpade]))
	}
	if len(suitMap[common.TrumpCardSuitHeart]) != 1 {
		t.Errorf("Expected 1 heart card, got %d", len(suitMap[common.TrumpCardSuitHeart]))
	}

	// 数字毎の分離
	numberMap := group.SeparateByNumber()
	if len(numberMap[10]) != 2 {
		t.Errorf("Expected 2 cards with number 10, got %d", len(numberMap[10]))
	}

	// 同じ数字かどうかの判定
	if !group.IsSameNumbers(common.TrumpCardCompareMethodAllowJoker) {
		t.Error("Expected same numbers with AllowJoker method")
	}

	// 含まれるスート・数字の取得
	suits := group.Suits()
	if len(suits) != 2 {
		t.Errorf("Expected 2 different suits, got %d", len(suits))
	}

	numbers := group.Numbers()
	if len(numbers) != 1 || numbers[0] != 10 {
		t.Errorf("Expected one number 10, got %v", numbers)
	}
}
