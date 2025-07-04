package cardgroup

import (
	"github.com/rpaka/umefugo-server/common"
)

// TrumpCardHierarchy interface for breaking circular dependency
type TrumpCardHierarchy interface {
	// Numbers returns the list of numbers in strength order
	Numbers() []int
	// StrongerThanCard compares if cardL is stronger than cardR
	StrongerThanCard(method common.TrumpCardCompareMethod, cardL, cardR *common.TrumpCard) bool
	// WeakerThanCard compares if cardL is weaker than cardR
	WeakerThanCard(method common.TrumpCardCompareMethod, cardL, cardR *common.TrumpCard) bool
}

// AdvancedTrumpCardHierarchy provides additional functionality for concrete implementations
type AdvancedTrumpCardHierarchy interface {
	TrumpCardHierarchy
	// StrongOrderDistanceFromCards returns the strength order distance between two cards
	StrongOrderDistanceFromCards(method common.TrumpCardCompareMethod, cardA, cardB *common.TrumpCard) int
}
