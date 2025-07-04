package common

import (
	"github.com/google/uuid"
)

// Card カード
type Card struct {
	// id カードID
	id uuid.UUID
	// publicStatus カード公開状態
	publicStatus CardPublicStatus
}

// NewCard creates a new Card instance
func NewCard() *Card {
	return &Card{
		id:           uuid.New(),
		publicStatus: CardPublicStatusPublic,
	}
}

// ID returns the card ID
func (c *Card) ID() uuid.UUID {
	return c.id
}

// PublicStatus returns the card public status
func (c *Card) PublicStatus() CardPublicStatus {
	return c.publicStatus
}

// SetPublicStatus sets the card public status
func (c *Card) SetPublicStatus(status CardPublicStatus) {
	c.publicStatus = status
}

// Equals 同じカードかどうかを判定する
func (c *Card) Equals(other *Card) bool {
	if other == nil {
		return false
	}
	return c.id == other.id
}
