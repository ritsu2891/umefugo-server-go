package cardgroup

import (
	"github.com/rpaka/umefugo-server/common"
)

// DaifugoTrumpCardGroup トランプカードの組み合わせ
type DaifugoTrumpCardGroup struct {
	// cards カードのスライス
	cards []*common.TrumpCard
	// groupType トランプカードの組み合わせの種類
	groupType DaifugoTrumpCardGroupType
}

// NewDaifugoTrumpCardGroup creates a new empty DaifugoTrumpCardGroup
func NewDaifugoTrumpCardGroup() *DaifugoTrumpCardGroup {
	return &DaifugoTrumpCardGroup{
		cards:     make([]*common.TrumpCard, 0),
		groupType: nil,
	}
}

// NewDaifugoTrumpCardGroupWithCards creates a new DaifugoTrumpCardGroup with cards
func NewDaifugoTrumpCardGroupWithCards(cards []*common.TrumpCard) *DaifugoTrumpCardGroup {
	group := &DaifugoTrumpCardGroup{
		cards:     make([]*common.TrumpCard, len(cards)),
		groupType: nil,
	}
	copy(group.cards, cards)
	return group
}

// Cards returns a copy of the cards slice
func (g *DaifugoTrumpCardGroup) Cards() []*common.TrumpCard {
	result := make([]*common.TrumpCard, len(g.cards))
	copy(result, g.cards)
	return result
}

// Add adds a card to the group
func (g *DaifugoTrumpCardGroup) Add(card *common.TrumpCard) {
	g.cards = append(g.cards, card)
}

// AddRange adds multiple cards to the group
func (g *DaifugoTrumpCardGroup) AddRange(cards []*common.TrumpCard) {
	g.cards = append(g.cards, cards...)
}

// Count returns the number of cards in the group
func (g *DaifugoTrumpCardGroup) Count() int {
	return len(g.cards)
}

// Get returns the card at the specified index
func (g *DaifugoTrumpCardGroup) Get(index int) *common.TrumpCard {
	if index < 0 || index >= len(g.cards) {
		return nil
	}
	return g.cards[index]
}

// Type returns the group type
func (g *DaifugoTrumpCardGroup) Type() DaifugoTrumpCardGroupType {
	return g.groupType
}

// SetType sets the group type
func (g *DaifugoTrumpCardGroup) SetType(groupType DaifugoTrumpCardGroupType) {
	g.groupType = groupType
}

// NotJokerCards 非ジョーカーカードを抽出する
func (g *DaifugoTrumpCardGroup) NotJokerCards() []*common.TrumpCard {
	var result []*common.TrumpCard
	for _, card := range g.cards {
		if !card.IsJoker() {
			result = append(result, card)
		}
	}
	return result
}

// JokerCards ジョーカーカードを抽出する
func (g *DaifugoTrumpCardGroup) JokerCards() []*common.TrumpCard {
	var result []*common.TrumpCard
	for _, card := range g.cards {
		if card.IsJoker() {
			result = append(result, card)
		}
	}
	return result
}

// HasAltJokerCards 代役情報があるジョーカーカードを抽出する
func (g *DaifugoTrumpCardGroup) HasAltJokerCards() []*common.TrumpCard {
	var result []*common.TrumpCard
	for _, card := range g.cards {
		if card.IsAltJoker() {
			result = append(result, card)
		}
	}
	return result
}

// NoAltJokerCards 代役情報がないジョーカーカードを抽出する
func (g *DaifugoTrumpCardGroup) NoAltJokerCards() []*common.TrumpCard {
	var result []*common.TrumpCard
	for _, card := range g.cards {
		if card.IsNoAltJoker() {
			result = append(result, card)
		}
	}
	return result
}

// HasAltSuitJokerCards 代役スート情報があるジョーカーカードを抽出する
func (g *DaifugoTrumpCardGroup) HasAltSuitJokerCards() []*common.TrumpCard {
	var result []*common.TrumpCard
	for _, card := range g.cards {
		if card.IsAltSuitJoker() {
			result = append(result, card)
		}
	}
	return result
}

// HasNoAltSuitJokerCards 代役スート情報がないジョーカーカードを抽出する
func (g *DaifugoTrumpCardGroup) HasNoAltSuitJokerCards() []*common.TrumpCard {
	var result []*common.TrumpCard
	for _, card := range g.cards {
		if card.IsNoAltSuitJoker() {
			result = append(result, card)
		}
	}
	return result
}

// HasAltNumberJokerCards 代役数値情報があるジョーカーカードを抽出する
func (g *DaifugoTrumpCardGroup) HasAltNumberJokerCards() []*common.TrumpCard {
	var result []*common.TrumpCard
	for _, card := range g.cards {
		if card.IsAltNumberJoker() {
			result = append(result, card)
		}
	}
	return result
}

// HasNoAltNumberJokerCards 代役数値情報がないジョーカーカードを抽出する
func (g *DaifugoTrumpCardGroup) HasNoAltNumberJokerCards() []*common.TrumpCard {
	var result []*common.TrumpCard
	for _, card := range g.cards {
		if card.IsNoAltNumberJoker() {
			result = append(result, card)
		}
	}
	return result
}

// CountJoker ジョーカーの枚数をカウントする
func (g *DaifugoTrumpCardGroup) CountJoker() int {
	return len(g.JokerCards())
}

// CountAltSuitJoker 代役スート情報があるジョーカーの枚数をカウントする
func (g *DaifugoTrumpCardGroup) CountAltSuitJoker() int {
	return len(g.HasAltSuitJokerCards())
}

// CountNoAltSuitJoker 代役スート情報がないジョーカーの枚数をカウントする
func (g *DaifugoTrumpCardGroup) CountNoAltSuitJoker() int {
	return len(g.HasNoAltSuitJokerCards())
}

// CountAltNumberJoker 代役数値情報があるジョーカーの枚数をカウントする
func (g *DaifugoTrumpCardGroup) CountAltNumberJoker() int {
	return len(g.HasAltNumberJokerCards())
}

// CountNoAltNumberJoker 代役数値情報がないジョーカーの枚数をカウントする
func (g *DaifugoTrumpCardGroup) CountNoAltNumberJoker() int {
	return len(g.HasNoAltNumberJokerCards())
}

// HasJoker ジョーカーがあるかどうかを判定する
func (g *DaifugoTrumpCardGroup) HasJoker() bool {
	return g.CountJoker() > 0
}

// HasNoAltJoker 代役情報がないジョーカーがあるかどうかを判定する
func (g *DaifugoTrumpCardGroup) HasNoAltJoker() bool {
	return len(g.NoAltJokerCards()) > 0
}

// HasAltSuitJoker 代役スート情報があるジョーカーがあるかどうかを判定する
func (g *DaifugoTrumpCardGroup) HasAltSuitJoker() bool {
	return g.CountAltSuitJoker() > 0
}

// HasNoAltSuitJoker 代役スート情報がないジョーカーがあるかどうかを判定する
func (g *DaifugoTrumpCardGroup) HasNoAltSuitJoker() bool {
	return g.CountNoAltSuitJoker() > 0
}

// HasAltNumberJoker 代役数値情報があるジョーカーがあるかどうかを判定する
func (g *DaifugoTrumpCardGroup) HasAltNumberJoker() bool {
	return g.CountAltNumberJoker() > 0
}

// HasNoAltNumberJoker 代役数値情報がないジョーカーがあるかどうかを判定する
func (g *DaifugoTrumpCardGroup) HasNoAltNumberJoker() bool {
	return g.CountNoAltNumberJoker() > 0
}

// SeparateBySuit スート毎にカードを仕分ける
func (g *DaifugoTrumpCardGroup) SeparateBySuit() map[common.TrumpCardSuit][]*common.TrumpCard {
	result := map[common.TrumpCardSuit][]*common.TrumpCard{
		common.TrumpCardSuitSpade:   {},
		common.TrumpCardSuitHeart:   {},
		common.TrumpCardSuitDiamond: {},
		common.TrumpCardSuitClub:    {},
	}

	for _, card := range g.cards {
		if card.IsJoker() {
			continue
		}
		result[card.Suit()] = append(result[card.Suit()], card)
	}

	return result
}

// SeparateByNumber 数字毎にカードを仕分ける
func (g *DaifugoTrumpCardGroup) SeparateByNumber() map[int][]*common.TrumpCard {
	result := make(map[int][]*common.TrumpCard)
	for i := common.NumberMin; i <= common.NumberMax; i++ {
		result[i] = []*common.TrumpCard{}
	}

	for _, card := range g.cards {
		if card.IsJoker() {
			continue
		}
		result[card.Number()] = append(result[card.Number()], card)
	}

	return result
}

// CountBySuit スート毎にカードをカウントする
func (g *DaifugoTrumpCardGroup) CountBySuit() map[common.TrumpCardSuit]int {
	result := map[common.TrumpCardSuit]int{
		common.TrumpCardSuitSpade:   0,
		common.TrumpCardSuitHeart:   0,
		common.TrumpCardSuitDiamond: 0,
		common.TrumpCardSuitClub:    0,
	}

	for _, card := range g.cards {
		if card.IsJoker() {
			continue
		}
		result[card.Suit()]++
	}

	return result
}

// CountByNumber 数字毎にカードをカウントする
func (g *DaifugoTrumpCardGroup) CountByNumber() map[int]int {
	result := make(map[int]int)
	for i := common.NumberMin; i <= common.NumberMax; i++ {
		result[i] = 0
	}

	for _, card := range g.cards {
		if card.IsJoker() {
			continue
		}
		result[card.Number()]++
	}

	return result
}

// Suits 含まれるスートを取得する
func (g *DaifugoTrumpCardGroup) Suits() []common.TrumpCardSuit {
	suitSet := make(map[common.TrumpCardSuit]bool)
	for _, card := range g.cards {
		if !card.IsJoker() {
			suitSet[card.Suit()] = true
		}
	}

	var result []common.TrumpCardSuit
	for suit := range suitSet {
		result = append(result, suit)
	}

	// ソート（enumの値順）
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}

// Numbers 含まれる数字を取得する
func (g *DaifugoTrumpCardGroup) Numbers() []int {
	numberSet := make(map[int]bool)
	for _, card := range g.cards {
		if !card.IsJoker() {
			numberSet[card.Number()] = true
		}
	}

	var result []int
	for number := range numberSet {
		result = append(result, number)
	}

	// ソート
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}

// baseCard 数値・スートを比較するためのカードを抽出する（代表カード）
func (g *DaifugoTrumpCardGroup) baseCard() *common.TrumpCard {
	if g.Count() == 0 {
		return common.CardNA
	}

	notJokerCards := g.NotJokerCards()
	if len(notJokerCards) != 0 {
		return notJokerCards[0]
	}

	hasAltJokerCards := g.HasAltJokerCards()
	if len(hasAltJokerCards) != 0 {
		return hasAltJokerCards[0]
	}

	return g.cards[0]
}

// WeakestCard 最弱数値のカードを取得する
func (g *DaifugoTrumpCardGroup) WeakestCard(method common.TrumpCardCompareMethod, hierarchy TrumpCardHierarchy) *common.TrumpCard {
	weakestCard := common.CardNA
	for _, card := range g.cards {
		if weakestCard.IsNA() || hierarchy.WeakerThanCard(method, card, weakestCard) {
			weakestCard = card
		}
	}
	return weakestCard
}

// IsSameSuits カードのスートが全て同じと見なせるかを判定する
func (g *DaifugoTrumpCardGroup) IsSameSuits(method common.TrumpCardCompareMethod) bool {
	baseCard := g.baseCard()
	if baseCard.IsNA() {
		return false
	}

	for _, card := range g.cards {
		if !common.EqualsSuit(method, baseCard, card) {
			return false
		}
	}
	return true
}

// IsSameNumbers カードの数値が全て同じと見なせるかを判定する
func (g *DaifugoTrumpCardGroup) IsSameNumbers(method common.TrumpCardCompareMethod) bool {
	baseCard := g.baseCard()
	if baseCard.IsNA() {
		return false
	}

	for _, card := range g.cards {
		if !common.EqualsNumber(method, baseCard, card) {
			return false
		}
	}
	return true
}

// IsSequence 階段かどうかを判定する
func (g *DaifugoTrumpCardGroup) IsSequence(hierarchy TrumpCardHierarchy, checkMinCount bool) bool {
	// 最低3枚以上なければ階段でない
	if checkMinCount && g.Count() < 3 {
		return false
	}

	// 数字の数を超える枚数は階段でない
	if common.NumbersCount < g.Count() {
		return false
	}

	// 判定処理用の情報を用意
	remJokers := g.CountJoker()             // 残りジョーカー枚数
	remNormalCards := g.Count() - remJokers // 残り非ジョーカー枚数
	numberDic := g.SeparateByNumber()       // 数字毎にカードを仕分ける
	inStairRange := false                   // 階段の範囲内

	for _, n := range hierarchy.Numbers() {
		// 同じ数字が2枚以上ある場合は階段ではない
		if len(numberDic[n]) > 1 {
			return false
		}

		if !inStairRange {
			// 階段範囲外継続
			if len(numberDic[n]) == 0 {
				continue
			}

			// 階段開始
			inStairRange = true
		} else {
			// 階段範囲内で途中の数字がない
			if len(numberDic[n]) == 0 {
				// ジョーカーでこの数字の補填を試し、不可能であれば階段でない
				remJokers--
				if remJokers < 0 {
					return false
				}
			}
		}

		// 残りのカードがない場合は終了
		remNormalCards -= len(numberDic[n])
		if remNormalCards == 0 {
			break
		}
	}

	return true
}

// ToSlice 通常のスライスに変換する
func (g *DaifugoTrumpCardGroup) ToSlice() []*common.TrumpCard {
	result := make([]*common.TrumpCard, len(g.cards))
	copy(result, g.cards)
	return result
}

// StrongOrderDistance 別のカードグループとの強さ序列距離を取得する
func (g *DaifugoTrumpCardGroup) StrongOrderDistance(method common.TrumpCardCompareMethod, hierarchy TrumpCardHierarchy, other *DaifugoTrumpCardGroup) int {
	weakestA := g.WeakestCard(method, hierarchy)
	weakestB := other.WeakestCard(method, hierarchy)

	// AdvancedTrumpCardHierarchy を実装している場合は詳細計算を行う
	if advHierarchy, ok := hierarchy.(AdvancedTrumpCardHierarchy); ok {
		return advHierarchy.StrongOrderDistanceFromCards(method, weakestA, weakestB)
	}

	// 基本的な比較のみ提供
	if hierarchy.StrongerThanCard(method, weakestB, weakestA) {
		return 1 // B > A
	}
	if hierarchy.StrongerThanCard(method, weakestA, weakestB) {
		return -1 // A > B
	}
	return 0 // A == B
}
