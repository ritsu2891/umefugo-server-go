package daifugo

// DaifugoPlayerPassStatus 大富豪プレイヤーパス状況
type DaifugoPlayerPassStatus int

const (
	// DaifugoPlayerPassStatusNone 未パス
	DaifugoPlayerPassStatusNone DaifugoPlayerPassStatus = iota
	// DaifugoPlayerPassStatusPassed パス済み
	DaifugoPlayerPassStatusPassed
)

// String returns the string representation of DaifugoPlayerPassStatus
func (s DaifugoPlayerPassStatus) String() string {
	switch s {
	case DaifugoPlayerPassStatusNone:
		return "未パス"
	case DaifugoPlayerPassStatusPassed:
		return "パス済み"
	default:
		return "不明"
	}
}
