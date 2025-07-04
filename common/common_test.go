package common

import (
	"testing"
)

func TestCard(t *testing.T) {
	card1 := NewCard()
	card2 := NewCard()

	// Test that different cards have different IDs
	if card1.ID() == card2.ID() {
		t.Error("Different cards should have different IDs")
	}

	// Test that cards are not equal by default
	if card1.Equals(card2) {
		t.Error("Different cards should not be equal")
	}

	// Test that a card equals itself
	if !card1.Equals(card1) {
		t.Error("A card should equal itself")
	}

	// Test public status
	if card1.PublicStatus() != CardPublicStatusPublic {
		t.Error("Default public status should be Public")
	}

	// Test setting public status
	card1.SetPublicStatus(CardPublicStatusPrivate)
	if card1.PublicStatus() != CardPublicStatusPrivate {
		t.Error("Public status should be Private after setting")
	}
}

func TestPlayer(t *testing.T) {
	player1 := NewPlayer("Player 1")
	player2 := NewPlayer("Player 2")

	// Test names
	if player1.Name() != "Player 1" {
		t.Errorf("Expected Player 1, got %s", player1.Name())
	}

	// Test that different players have different IDs
	if player1.ID() == player2.ID() {
		t.Error("Different players should have different IDs")
	}

	// Test that players are not equal by default
	if player1.Equals(player2) {
		t.Error("Different players should not be equal")
	}

	// Test setting name
	player1.SetName("Updated Player 1")
	if player1.Name() != "Updated Player 1" {
		t.Error("Player name should be updated")
	}
}

func TestTrumpCard(t *testing.T) {
	// Test normal card creation
	card, err := NewNormalTrumpCard(TrumpCardSuitClub, 1)
	if err != nil {
		t.Fatalf("Failed to create normal trump card: %v", err)
	}

	if card.Type() != TrumpCardTypeNormal {
		t.Error("Card type should be Normal")
	}

	if card.Suit() != TrumpCardSuitClub {
		t.Error("Card suit should be Club")
	}

	if card.Number() != 1 {
		t.Error("Card number should be 1")
	}

	if card.IsJoker() {
		t.Error("Normal card should not be joker")
	}

	// Test string representation
	expectedString := "♣A"
	if card.String() != expectedString {
		t.Errorf("Expected %s, got %s", expectedString, card.String())
	}

	// Test joker creation
	joker := NewJokerTrumpCard()
	if !joker.IsJoker() {
		t.Error("Joker should be joker")
	}

	if joker.HasAltSuit() {
		t.Error("New joker should not have alt suit")
	}

	// Test setting alt suit
	err = joker.SetAltSuit(TrumpCardSuitHeart)
	if err != nil {
		t.Fatalf("Failed to set alt suit: %v", err)
	}

	if !joker.HasAltSuit() {
		t.Error("Joker should have alt suit after setting")
	}

	if joker.AltSuit() != TrumpCardSuitHeart {
		t.Error("Alt suit should be Heart")
	}

	// Test invalid card creation
	_, err = NewNormalTrumpCard(TrumpCardSuitJoker, 1)
	if err == nil {
		t.Error("Should not be able to create normal card with joker suit")
	}

	_, err = NewNormalTrumpCard(TrumpCardSuitClub, 0)
	if err == nil {
		t.Error("Should not be able to create normal card with invalid number")
	}
}

func TestTrumpCardComparison(t *testing.T) {
	card1, _ := NewNormalTrumpCard(TrumpCardSuitClub, 5)
	card2, _ := NewNormalTrumpCard(TrumpCardSuitClub, 10)
	card3, _ := NewNormalTrumpCard(TrumpCardSuitHeart, 5)

	// Test number comparison
	if !EqualsNumber(TrumpCardCompareMethodStrict, card1, card3) {
		t.Error("Cards with same number should be equal in number")
	}

	if !LessThan(TrumpCardCompareMethodStrict, card1, card2) {
		t.Error("5 should be less than 10")
	}

	if !GreaterThan(TrumpCardCompareMethodStrict, card2, card1) {
		t.Error("10 should be greater than 5")
	}

	// Test suit comparison
	if !EqualsSuit(TrumpCardCompareMethodStrict, card1, card2) {
		t.Error("Cards with same suit should be equal in suit")
	}

	if EqualsSuit(TrumpCardCompareMethodStrict, card1, card3) {
		t.Error("Cards with different suits should not be equal in suit")
	}
}

func TestNumberRotation(t *testing.T) {
	// Test next number
	if NextNumber(1) != 2 {
		t.Errorf("Next number of 1 should be 2, got %d", NextNumber(1))
	}

	if NextNumber(13) != 1 {
		t.Errorf("Next number of 13 should be 1, got %d", NextNumber(13))
	}

	// Test previous number
	if PreviousNumber(2) != 1 {
		t.Errorf("Previous number of 2 should be 1, got %d", PreviousNumber(2))
	}

	if PreviousNumber(1) != 13 {
		t.Errorf("Previous number of 1 should be 13, got %d", PreviousNumber(1))
	}
}
