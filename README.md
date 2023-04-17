# go-package-test

## go.mod を作成

```go
// 実行したところをルートとして、モジュールとする
$ go mod init kyoruni/go-package-test

// 必要なモジュールのダウンロードを行う
// 今回はモジュールのパスと、main.goのインポートに書いたパスが同じなので、ダウンロードの必要がない -> 特に何も行われない
$ go mod tidy
```

## 実行してみる

```go
$ go run main.go
6のダメージを与えた！
```