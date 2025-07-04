package common

// CardPublicStatus カード公開状態
type CardPublicStatus int

const (
	// CardPublicStatusPrivate 非公開
	CardPublicStatusPrivate CardPublicStatus = iota
	// CardPublicStatusPublic 公開
	CardPublicStatusPublic
)

// String returns the string representation of CardPublicStatus
func (s CardPublicStatus) String() string {
	switch s {
	case CardPublicStatusPrivate:
		return "非公開"
	case CardPublicStatusPublic:
		return "公開"
	default:
		return "不明"
	}
}
