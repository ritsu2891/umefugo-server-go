package filter

import (
	"testing"

	"github.com/rpaka/umefugo-server/common"
	"github.com/rpaka/umefugo-server/daifugo/cardgroup"
)

func TestDaifugoTrumpCardGroupFilterPair_GetFilteredCardsWithNoCondition(t *testing.T) {
	// テスト用のカードを作成
	card1, _ := common.NewNormalTrumpCard(common.TrumpCardSuitSpade, 3)
	card2, _ := common.NewNormalTrumpCard(common.TrumpCardSuitHeart, 3)
	card3, _ := common.NewNormalTrumpCard(common.TrumpCardSuitSpade, 4)
	card4 := common.NewJokerTrumpCard()

	cards := []*common.TrumpCard{card1, card2, card3, card4}

	sourceCards := cardgroup.NewDaifugoTrumpCardGroupWithCards(cards)
	requiredCards := cardgroup.NewDaifugoTrumpCardGroup()

	filter := NewDaifugoTrumpCardGroupFilterPair(sourceCards, requiredCards)
	result := filter.GetFilteredCardsWithNoCondition()

	// ペアを構成可能なカードが含まれているかチェック
	// 3のペア（♠3, ♥3）とジョーカーが含まれているはず
	if result.Count() == 0 {
		t.Error("Expected some cards, but got none")
	}

	// 3のスペードとハートが含まれているはず
	found3Spade := false
	found3Heart := false
	foundJoker := false

	for _, card := range result.ToSlice() {
		if card.Suit() == common.TrumpCardSuitSpade && card.Number() == 3 {
			found3Spade = true
		}
		if card.Suit() == common.TrumpCardSuitHeart && card.Number() == 3 {
			found3Heart = true
		}
		if card.IsJoker() {
			foundJoker = true
		}
	}

	if !found3Spade {
		t.Error("Expected to find 3 of Spades")
	}
	if !found3Heart {
		t.Error("Expected to find 3 of Hearts")
	}
	if !foundJoker {
		t.Error("Expected to find Joker")
	}
}

func TestDaifugoTrumpCardGroupFilterConditionStatus(t *testing.T) {
	// 基本的なenum値の確認
	if DaifugoTrumpCardGroupFilterConditionStatusDefault != 0 {
		t.Error("Expected Default status to be 0")
	}
	if DaifugoTrumpCardGroupFilterConditionStatusSatisfied != 1 {
		t.Error("Expected Satisfied status to be 1")
	}
	if DaifugoTrumpCardGroupFilterConditionStatusUnSatisfied != 2 {
		t.Error("Expected UnSatisfied status to be 2")
	}
}
