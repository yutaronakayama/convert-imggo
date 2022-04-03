package main

/*
ディレクトリを指定する
--指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
--ディレクトリ以下は再帰的に処理する
変換前と変換後の画像形式を指定できる（オプション）
以下を満たすように開発してください
--mainパッケージと分離する
--自作パッケージと標準パッケージと準標準パッケージのみ使う
--準標準パッケージ：golang.org/x以下のパッケージ
ユーザ定義型を作ってみる
GoDocを生成してみる
Go Modulesを使ってみる
*/

import (
	"flag"
	"fmt"
	"github.com/yutaronakayama/convert_imggo/converter"
	"time"
	"os"
)

func main() {
	flag.Parse()
    if err := converter.toJPEG(os.Stdin, os.Stdout); err != nil {
        fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
        os.Exit(1)
    }
	//fmt.Printf("Hello World Time: %v\n", time.Now())
}
