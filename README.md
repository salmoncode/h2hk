# h2hk

ひらがなを半角カタカナに変換するCLIツール

## Features

h2hk は、ひらがなを半角カタカナに変換するシンプルなコマンドラインツールです。
フォームに名前や住所のカナを入力する際に役立ちます。

## 🚀 Installation

### Goがインストールされている場合

```go
go install github.com/salmoncode/h2hk@latest
```

### GitHubからcloneしてビルドする場合

```go
git clone https://github.com/salmoncode/h2hk.git
cd h2hk
go build -o h2hk
```

## Usage

```go
h2hk <text>
```

例：
```go
$ h2hk こんにちは
ｺﾝﾆﾁﾊ
```

全角スペースにも対応しています。
```go
$ h2hk たなか　たろう
ﾀﾅｶ ﾀﾛｳ
```

## 🔧 開発者向け情報

このプロジェクトは[Go言語](https://go.dev/)で書かれています。

### ビルド手順

```go
go build -o h2hk
```

### コミット
[Conventional Commits v1.0.0](https://www.conventionalcommits.org/ja/v1.0.0/)に従ってください

## License

MIT License

## Author
- name: salmncode
- GitHub: https://github.com/salmoncode