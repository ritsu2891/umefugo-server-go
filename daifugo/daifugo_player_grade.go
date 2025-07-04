package daifugo

// DaifugoPlayerGrade 大富豪プレイヤー階級
type DaifugoPlayerGrade int

const (
	// DaifugoPlayerGradeUltimateMillionaire 超大富豪
	DaifugoPlayerGradeUltimateMillionaire DaifugoPlayerGrade = iota
	// DaifugoPlayerGradeGreatMillionaire 大富豪
	DaifugoPlayerGradeGreatMillionaire
	// DaifugoPlayerGradeMillionaire 富豪
	DaifugoPlayerGradeMillionaire
	// DaifugoPlayerGradeCommoner 平民
	DaifugoPlayerGradeCommoner
	// DaifugoPlayerGradePoor 貧民
	DaifugoPlayerGradePoor
	// DaifugoPlayerGradeGreatPoor 大貧民
	DaifugoPlayerGradeGreatPoor
	// DaifugoPlayerGradeUltimatePoor 超大貧民
	DaifugoPlayerGradeUltimatePoor
	// DaifugoPlayerGradeNotAssigned 未割当
	DaifugoPlayerGradeNotAssigned
)

// String returns the string representation of DaifugoPlayerGrade
func (g DaifugoPlayerGrade) String() string {
	switch g {
	case DaifugoPlayerGradeUltimateMillionaire:
		return "超大富豪"
	case DaifugoPlayerGradeGreatMillionaire:
		return "大富豪"
	case DaifugoPlayerGradeMillionaire:
		return "富豪"
	case DaifugoPlayerGradeCommoner:
		return "平民"
	case DaifugoPlayerGradePoor:
		return "貧民"
	case DaifugoPlayerGradeGreatPoor:
		return "大貧民"
	case DaifugoPlayerGradeUltimatePoor:
		return "超大貧民"
	case DaifugoPlayerGradeNotAssigned:
		return "未割当"
	default:
		return "不明"
	}
}
