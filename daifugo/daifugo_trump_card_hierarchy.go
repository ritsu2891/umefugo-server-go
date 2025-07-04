package daifugo

import (
	"math"

	"github.com/rpaka/umefugo-server/common"
)

// DaifugoTrumpCardHierarchy constants
const (
	// DefaultWeakestNumber デフォルトの最弱数値
	DefaultWeakestNumber = 3
	// DefaultStrongestNumber デフォルトの最強数値
	DefaultStrongestNumber = 2
	// OrderNA 序列無し
	OrderNA = math.MinInt32
	// OrderInf 序列最強
	OrderInf = math.MaxInt32
)

// DaifugoTrumpCardHierarchy 大富豪のトランプカード強さ序列
type DaifugoTrumpCardHierarchy struct {
	// weakestNumber 最弱数値
	weakestNumber int
	// strongestNumber 最強数値
	strongestNumber int
	// isRevolution 革命状態か
	isRevolution bool
}

// NewDaifugoTrumpCardHierarchy creates a new DaifugoTrumpCardHierarchy instance
func NewDaifugoTrumpCardHierarchy() *DaifugoTrumpCardHierarchy {
	return &DaifugoTrumpCardHierarchy{
		weakestNumber:   DefaultWeakestNumber,
		strongestNumber: DefaultStrongestNumber,
		isRevolution:    false,
	}
}

// WeakestNumber returns the weakest number
func (h *DaifugoTrumpCardHierarchy) WeakestNumber() int {
	return h.weakestNumber
}

// StrongestNumber returns the strongest number
func (h *DaifugoTrumpCardHierarchy) StrongestNumber() int {
	return h.strongestNumber
}

// IsRevolution returns whether it's in revolution state
func (h *DaifugoTrumpCardHierarchy) IsRevolution() bool {
	return h.isRevolution
}

// IsNormal returns whether it's in normal state
func (h *DaifugoTrumpCardHierarchy) IsNormal() bool {
	return !h.isRevolution
}

// Numbers 強さの順に並べた数字のリストを生成する
func (h *DaifugoTrumpCardHierarchy) Numbers() []int {
	var numbers []int
	for n := h.weakestNumber; n != common.NumberNA; n = h.NextStrongNumber(n) {
		numbers = append(numbers, n)
	}
	return numbers
}

// RotateNumberInHierarchy 序列の中で数値を回転させる
// 序列を超えた場合はcommon.NumberNAを返す
func (h *DaifugoTrumpCardHierarchy) RotateNumberInHierarchy(number, amount int) int {
	if h.isRevolution {
		amount *= -1 // 革命中は逆順にする
	}
	rotatedNumber := common.RotateNumber(number, amount)
	if amount > 0 && !h.StrongerThan(rotatedNumber, number) {
		return common.NumberNA
	}
	if amount < 0 && !h.WeakerThan(rotatedNumber, number) {
		return common.NumberNA
	}
	return rotatedNumber
}

// NextWeakNumber 序列の中で次に弱い数値を取得する
// ただし存在しない場合はcommon.NumberNAを返す
func (h *DaifugoTrumpCardHierarchy) NextWeakNumber(number int) int {
	return h.RotateNumberInHierarchy(number, -1)
}

// NextStrongNumber 序列の中で次に強い数値を取得する
// ただし存在しない場合はcommon.NumberNAを返す
func (h *DaifugoTrumpCardHierarchy) NextStrongNumber(number int) int {
	return h.RotateNumberInHierarchy(number, 1)
}

// RotateNumberInHierarchyCard 序列の中で数値を回転させる（カード版）
func (h *DaifugoTrumpCardHierarchy) RotateNumberInHierarchyCard(card *common.TrumpCard, amount int) int {
	return h.RotateNumberInHierarchy(card.Number(), amount)
}

// NextWeakNumberCard 序列の中で次に弱い数値を取得する（カード版）
func (h *DaifugoTrumpCardHierarchy) NextWeakNumberCard(card *common.TrumpCard) int {
	return h.NextWeakNumber(card.Number())
}

// NextStrongNumberCard 序列の中で次に強い数値を取得する（カード版）
func (h *DaifugoTrumpCardHierarchy) NextStrongNumberCard(card *common.TrumpCard) int {
	return h.NextStrongNumber(card.Number())
}

// ResetHierarchy 序列をデフォルトに戻す
func (h *DaifugoTrumpCardHierarchy) ResetHierarchy() {
	h.strongestNumber = DefaultStrongestNumber
	h.weakestNumber = DefaultWeakestNumber
}

// SetHierarchy 指定数値を最強として序列を設定する
func (h *DaifugoTrumpCardHierarchy) SetHierarchy(strongestNumber int) {
	h.strongestNumber = strongestNumber
	h.weakestNumber = common.PreviousNumber(h.strongestNumber)
}

// setRevolution 革命状態を設定する（private）
func (h *DaifugoTrumpCardHierarchy) setRevolution(revolution bool) {
	h.isRevolution = revolution
}

// ActivateRevolution 革命状態にする
func (h *DaifugoTrumpCardHierarchy) ActivateRevolution() {
	h.setRevolution(true)
}

// DeactivateRevolution 革命状態を解除する
func (h *DaifugoTrumpCardHierarchy) DeactivateRevolution() {
	h.setRevolution(false)
}

// ToggleRevolution 革命状態を反転させる
func (h *DaifugoTrumpCardHierarchy) ToggleRevolution() {
	h.setRevolution(!h.isRevolution)
}

// StrongOrder 数値の序列を取得する
func (h *DaifugoTrumpCardHierarchy) StrongOrder(number int) int {
	if number == common.NumberNA {
		return OrderNA
	}
	if number == common.NumberAny || number == common.NumberJoker {
		return OrderInf
	}
	if common.NumberMin <= number && number <= h.strongestNumber {
		return common.NumbersCount - (h.strongestNumber - number)
	}
	if h.weakestNumber <= number && number <= common.NumberMax {
		return number - h.weakestNumber + 1
	}
	return OrderNA
}

// strongOrderForCards トランプカードの数値の序列を取得する（private）
func (h *DaifugoTrumpCardHierarchy) strongOrderForCards(method common.TrumpCardCompareMethod, cardL, cardR *common.TrumpCard) (int, int) {
	numberA, numberB := common.GetNumberForCompare(method, cardL, cardR)
	return h.StrongOrder(numberA), h.StrongOrder(numberB)
}

// StrongerThan 数値が強いかどうかを判定する
func (h *DaifugoTrumpCardHierarchy) StrongerThan(numberL, numberR int) bool {
	return h.StrongOrder(numberL) > h.StrongOrder(numberR)
}

// WeakerThan 数値が弱いかどうかを判定する（private）
func (h *DaifugoTrumpCardHierarchy) WeakerThan(numberL, numberR int) bool {
	return h.StrongOrder(numberL) < h.StrongOrder(numberR)
}

// ListStrongerThan 指定数値より強い数値のリストを取得する
func (h *DaifugoTrumpCardHierarchy) ListStrongerThan(baseNumber int, containBase bool) []int {
	var res []int
	if containBase {
		res = append(res, baseNumber)
	}
	for n := h.NextStrongNumber(baseNumber); n != common.NumberNA; n = h.NextStrongNumber(n) {
		res = append(res, n)
	}
	return res
}

// ListWeakerThan 指定数値より弱い数値のリストを取得する
func (h *DaifugoTrumpCardHierarchy) ListWeakerThan(baseNumber int, containBase bool) []int {
	var res []int
	if containBase {
		res = append(res, baseNumber)
	}
	for n := h.NextWeakNumber(baseNumber); n != common.NumberNA; n = h.NextWeakNumber(n) {
		res = append(res, n)
	}
	return res
}

// EqualsStrong 数値の強さが同じか
func (h *DaifugoTrumpCardHierarchy) EqualsStrong(method common.TrumpCardCompareMethod, cardL, cardR *common.TrumpCard) bool {
	orderL, orderR := h.strongOrderForCards(method, cardL, cardR)
	if orderL == OrderNA || orderR == OrderNA {
		return false // 序列が無い場合は常に不成立
	}
	if orderL == OrderInf && orderR == OrderInf {
		return true // Infinity == Infinity は常に成立
	}
	if orderL == OrderInf || orderR == OrderInf {
		return false // Infinity == (Any) は常に不成立
	}
	return orderL == orderR
}

// StrongerThanCard カードの強さが大なりか
// cardL > cardR (L is Stronger Than R) をtrueとする
func (h *DaifugoTrumpCardHierarchy) StrongerThanCard(method common.TrumpCardCompareMethod, cardL, cardR *common.TrumpCard) bool {
	// L > R (L "is Stronger Than" R)
	orderL, orderR := h.strongOrderForCards(method, cardL, cardR)
	if orderL == OrderNA || orderR == OrderNA {
		return false // 序列が無い場合は常に不成立
	}
	if orderL == OrderInf && orderR == OrderInf {
		return false // Infinity > Infinity は不成立とする
	}
	if orderL == OrderInf {
		return h.IsNormal() // Infinity > (Any) は常に成立（通常時） Infinity < (Any) は常に不成立（革命時）
	}
	if orderR == OrderInf {
		return !h.IsNormal() // (Any) > Infinity は常に不成立（通常時）　(Any) < Infinity は常に成立（革命時）
	}
	if h.IsNormal() {
		return orderL > orderR
	}
	return orderL < orderR
}

// StrongerThanOrEqualsCard 数値の強さが大なり以上か
// cardL >= cardR (L is Stronger Than Or Equals R) をtrueとする
func (h *DaifugoTrumpCardHierarchy) StrongerThanOrEqualsCard(method common.TrumpCardCompareMethod, cardL, cardR *common.TrumpCard) bool {
	return h.EqualsStrong(method, cardL, cardR) || h.StrongerThanCard(method, cardL, cardR)
}

// WeakerThanCard カードの強さが小なりか
// cardL < cardR (L is Weaker Than R) をtrueとする
func (h *DaifugoTrumpCardHierarchy) WeakerThanCard(method common.TrumpCardCompareMethod, cardL, cardR *common.TrumpCard) bool {
	// L < R (L "is Weaker Than" R)
	orderL, orderR := h.strongOrderForCards(method, cardL, cardR)
	if orderL == OrderNA || orderR == OrderNA {
		return false // 序列が無い場合は常に不成立
	}
	if orderL == OrderInf && orderR == OrderInf {
		return false // Infinity < Infinity は不成立とする
	}
	if orderL == OrderInf {
		return !h.IsNormal() // Infinity < (Any) は常に不成立（通常時）　Infinity > (Any) は常に成立（革命時）
	}
	if orderR == OrderInf {
		return h.IsNormal() // (Any) < Infinity は常に成立（通常時）　(Any) > Infinity は常に不成立（革命時）
	}
	if h.IsNormal() {
		return orderL < orderR
	}
	return orderL > orderR
}

// WeakerThanOrEqualsCard 数値の強さが小なり以下か
// cardL <= cardR (L is Weaker Than Or Equals R) をtrueとする
func (h *DaifugoTrumpCardHierarchy) WeakerThanOrEqualsCard(method common.TrumpCardCompareMethod, cardL, cardR *common.TrumpCard) bool {
	return h.EqualsStrong(method, cardL, cardR) || h.WeakerThanCard(method, cardL, cardR)
}

// StrongOrderDistanceFromCards returns the strength order distance between two cards
func (h *DaifugoTrumpCardHierarchy) StrongOrderDistanceFromCards(method common.TrumpCardCompareMethod, cardA, cardB *common.TrumpCard) int {
	orderA, orderB := h.strongOrderForCards(method, cardA, cardB)
	if orderA == OrderNA || orderB == OrderNA {
		return OrderNA
	}
	if orderA == OrderInf && orderB == OrderInf {
		return 0
	}
	if orderA == OrderInf || orderB == OrderInf {
		return OrderInf
	}
	return orderB - orderA
}
