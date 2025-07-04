package daifugo

import (
	"github.com/rpaka/umefugo-server/common"
)

// DaifugoPlayer 大富豪プレイヤー
type DaifugoPlayer struct {
	*common.Player
	// passStatus パス状況
	passStatus DaifugoPlayerPassStatus
	// exchangeStatus カード交換状況
	exchangeStatus DaifugoPlayerExchangeStatus
	// currentGrade 現ラウンド階級
	currentGrade DaifugoPlayerGrade
	// prevGrade 前ラウンド階級
	prevGrade DaifugoPlayerGrade
	// currentRank 現ラウンド順位
	currentRank int
	// prevRank 前ラウンド順位
	prevRank int
}

// NewDaifugoPlayer creates a new DaifugoPlayer instance
func NewDaifugoPlayer(name string) *DaifugoPlayer {
	return &DaifugoPlayer{
		Player:         common.NewPlayer(name),
		passStatus:     DaifugoPlayerPassStatusNone,
		exchangeStatus: DaifugoPlayerExchangeStatusNoExchange,
		currentGrade:   DaifugoPlayerGradeNotAssigned,
		prevGrade:      DaifugoPlayerGradeNotAssigned,
		currentRank:    0,
		prevRank:       0,
	}
}

// PassStatus returns the pass status
func (p *DaifugoPlayer) PassStatus() DaifugoPlayerPassStatus {
	return p.passStatus
}

// SetPassStatus sets the pass status
func (p *DaifugoPlayer) SetPassStatus(status DaifugoPlayerPassStatus) {
	p.passStatus = status
}

// ExchangeStatus returns the exchange status
func (p *DaifugoPlayer) ExchangeStatus() DaifugoPlayerExchangeStatus {
	return p.exchangeStatus
}

// SetExchangeStatus sets the exchange status
func (p *DaifugoPlayer) SetExchangeStatus(status DaifugoPlayerExchangeStatus) {
	p.exchangeStatus = status
}

// CurrentGrade returns the current grade
func (p *DaifugoPlayer) CurrentGrade() DaifugoPlayerGrade {
	return p.currentGrade
}

// SetCurrentGrade sets the current grade
func (p *DaifugoPlayer) SetCurrentGrade(grade DaifugoPlayerGrade) {
	p.currentGrade = grade
}

// PrevGrade returns the previous grade
func (p *DaifugoPlayer) PrevGrade() DaifugoPlayerGrade {
	return p.prevGrade
}

// SetPrevGrade sets the previous grade
func (p *DaifugoPlayer) SetPrevGrade(grade DaifugoPlayerGrade) {
	p.prevGrade = grade
}

// CurrentRank returns the current rank
func (p *DaifugoPlayer) CurrentRank() int {
	return p.currentRank
}

// SetCurrentRank sets the current rank
func (p *DaifugoPlayer) SetCurrentRank(rank int) {
	p.currentRank = rank
}

// PrevRank returns the previous rank
func (p *DaifugoPlayer) PrevRank() int {
	return p.prevRank
}

// SetPrevRank sets the previous rank
func (p *DaifugoPlayer) SetPrevRank(rank int) {
	p.prevRank = rank
}
