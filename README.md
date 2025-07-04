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