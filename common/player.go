package common

import (
	"github.com/google/uuid"
)

// Player プレイヤー
type Player struct {
	// id プレイヤーID
	id uuid.UUID
	// name プレイヤー名
	name string
}

// NewPlayer creates a new Player instance
func NewPlayer(name string) *Player {
	return &Player{
		id:   uuid.New(),
		name: name,
	}
}

// ID returns the player ID
func (p *Player) ID() uuid.UUID {
	return p.id
}

// Name returns the player name
func (p *Player) Name() string {
	return p.name
}

// SetName sets the player name
func (p *Player) SetName(name string) {
	p.name = name
}

// Equals 同じプレイヤーかどうかを判定する
func (p *Player) Equals(other *Player) bool {
	if other == nil {
		return false
	}
	return p.id == other.id
}
