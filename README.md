# umefugo-server-go

🃏 umefugoサーバ部分Goで作り直し

## プロジェクト構造

```
umefugo-go/
├── common/                     # 共通データ構造とユーティリティ
│   ├── card.go                # 基本カードクラス
│   ├── card_public_status.go  # カード公開状態enum
│   ├── player.go              # プレイヤークラス
│   ├── trump_card.go          # トランプカード実装
│   ├── trump_card_compare_method.go # トランプカード比較方法enum
│   ├── trump_card_methods.go  # トランプカード比較・ユーティリティメソッド
│   ├── trump_card_suit.go     # トランプカードスートenum
│   ├── trump_card_type.go     # トランプカード種別enum
│   └── common_test.go         # テストファイル
├── daifugo/                   # 大富豪ゲームロジック
│   ├── daifugo_player.go      # 大富豪プレイヤークラス
│   ├── daifugo_player_exchange_status.go # プレイヤーカード交換状況enum
│   ├── daifugo_player_grade.go # プレイヤー階級enum
│   ├── daifugo_player_pass_status.go # プレイヤーパス状況enum
│   ├── daifugo_trump_card_hierarchy.go # トランプカード強さ序列
│   ├── daifugo_test.go        # テストファイル
│   └── cardgroup/             # カードグループ関連
│       ├── daifugo_trump_card_group.go # トランプカードグループ
│       ├── daifugo_trump_card_group_type.go # カードグループタイプinterface
│       ├── daifugo_trump_card_group_type_class.go # カードグループ種別enum
│       └── cardgroup_test.go  # テストファイル
├── go.mod                     # Goモジュール定義
├── go.sum                     # 依存関係チェックサム
└── README.md                  # このファイル
```

## 移植済み機能

### Commonパッケージ

元のC#プロジェクトの`umefugo-server/common`配下のクラスをGoで完全に移植済み：

- **Card**: 基本的なカードクラス（UUID識別子、公開状態管理）
- **CardPublicStatus**: カード公開状態のenum（公開/非公開）
- **Player**: プレイヤークラス（UUID識別子、名前管理）
- **TrumpCard**: トランプカードクラス（継承の代わりに埋め込みを使用）
  - 通常カードとジョーカーの管理
  - 代役スート・数値の設定
  - カード文字列表現（マーク付き）
- **TrumpCardSuit**: トランプスートenum（♣♦♥♠、ジョーカー、任意、未割当）
- **TrumpCardType**: トランプカード種別enum（通常、ジョーカー）
- **TrumpCardCompareMethod**: カード比較方法enum（4つの比較戦略）
- **比較・ユーティリティ関数**: カードの大小比較、数値回転など

### Daifugoパッケージ

元のC#プロジェクトの`umefugo-server/daifugo`配下のクラスを部分的に移植済み：

- **DaifugoPlayer**: 大富豪プレイヤークラス（基本Playerを拡張）
  - パス状況、交換状況、階級、順位の管理
- **DaifugoPlayerExchangeStatus**: カード交換状況enum
- **DaifugoPlayerGrade**: プレイヤー階級enum（超大富豪〜超大貧民）
- **DaifugoPlayerPassStatus**: パス状況enum
- **DaifugoTrumpCardHierarchy**: カード強さ序列管理
  - 革命状態の管理
  - 数値の強さ序列計算
  - カード比較メソッド

#### Cardgroupサブパッケージ（完全移植済み）

- **DaifugoTrumpCardGroup**: カードグループクラス（完全実装）
  - ジョーカー・非ジョーカーの分離
  - スート・数字による分類
  - 代役情報の管理
  - 最弱カード取得（WeakestCard）
  - 階段判定（IsSequence）
  - 強さ序列距離計算（StrongOrderDistance）
- **DaifugoTrumpCardGroupType**: カードグループタイプinterface
- **DaifugoTrumpCardGroupTypeClass**: カードグループ種別enum
- **カードグループ条件系クラス（完全移植済み）**:
  - **DaifugoTrumpCardGroupCondition**: 基底条件interface
  - **DaifugoTrumpCardGroupTypeSingleCondition**: 単体カード条件
  - **DaifugoTrumpCardGroupTypePairCondition**: ペアカード条件
  - **DaifugoTrumpCardGroupStrongerCondition**: 強さ比較条件
  - **DaifugoTrumpCardGroupHierarchyCondition**: 数字縛り条件
  - **DaifugoTrumpCardGroupSuitCondition**: スート縛り条件
- **カードグループタイプ実装（完全移植済み）**:
  - **DaifugoTrumpCardGroupTypeSingle**: 単体カードタイプ
  - **DaifugoTrumpCardGroupTypePair**: ペアカードタイプ
  - **DaifugoTrumpCardGroupTypeSequence**: 階段カードタイプ
- **TrumpCardHierarchy**: 循環import回避用interface
  - **AdvancedTrumpCardHierarchy**: 拡張機能用interface

## 依存関係

- `github.com/google/uuid` - UUID生成用

## ビルドとテスト

```bash
# プロジェクトのルートディレクトリに移動
cd umefugo-go

# 依存関係の解決
go mod tidy

# 全パッケージのビルド
go build ./...

# テストの実行
go test ./... -v

# 特定パッケージのテスト実行
go test ./common -v
```

## 設計上の変更点

### C#からGoへの移植における主な変更：

1. **継承 → 埋め込み**: C#の継承をGoの構造体埋め込みで実装
2. **プロパティ → ゲッター/セッターメソッド**: C#のプロパティをGo慣例に従ったメソッドで実装
3. **例外 → エラー**: C#の例外をGoのerrorインターフェースで実装
4. **static → パッケージレベル関数**: C#のstaticメソッドをGoのパッケージレベル関数で実装
5. **namespace → package**: C#のnamespaceをGoのpackageで実装

## 課題

### 循環Import
- **問題**: `daifugo/cardgroup`が`daifugo`の`DaifugoTrumpCardHierarchy`を参照
- **解決策**: 
  1. `TrumpCardHierarchy` interfaceを`cardgroup`パッケージに定義
  2. `DaifugoTrumpCardHierarchy`がこのinterfaceを実装
  3. `AdvancedTrumpCardHierarchy` interfaceで拡張機能を提供
  4. 型アサーションにより具体実装の詳細機能を利用