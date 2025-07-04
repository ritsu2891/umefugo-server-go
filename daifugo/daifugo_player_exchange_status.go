package daifugo

// DaifugoPlayerExchangeStatus 大富豪プレイヤーカード交換状況
type DaifugoPlayerExchangeStatus int

const (
	// DaifugoPlayerExchangeStatusNoExchange 交換対象外
	DaifugoPlayerExchangeStatusNoExchange DaifugoPlayerExchangeStatus = iota
	// DaifugoPlayerExchangeStatusSelecting 交換カード選択中
	DaifugoPlayerExchangeStatusSelecting
	// DaifugoPlayerExchangeStatusWaitingPair 交換相手待機中
	DaifugoPlayerExchangeStatusWaitingPair
	// DaifugoPlayerExchangeStatusExchanged 交換済み
	DaifugoPlayerExchangeStatusExchanged
)

// String returns the string representation of DaifugoPlayerExchangeStatus
func (s DaifugoPlayerExchangeStatus) String() string {
	switch s {
	case DaifugoPlayerExchangeStatusNoExchange:
		return "交換対象外"
	case DaifugoPlayerExchangeStatusSelecting:
		return "交換カード選択中"
	case DaifugoPlayerExchangeStatusWaitingPair:
		return "交換相手待機中"
	case DaifugoPlayerExchangeStatusExchanged:
		return "交換済み"
	default:
		return "不明"
	}
}
