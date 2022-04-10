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
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		//標準エラー出力os.Stderr
		fmt.Fprintf(
			os.Stderr,
			`Usage of %s: %s [OPTIONS] ARGS...Options\n`,
			os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	//Flagでのオプション定義
	var (
		src = flag.String("src", "", "please specify -src flag")
		dst = flag.String("dst", "", "please specify -dst flag")
		dir = flag.String("dir", "", "please specify -dir flag")
	)

	flag.Parse()

	//オプション引数チェック
	if flag.NArg() != 0 && flag.NFlag() >= 2 {
		log.Fatal("main.go: 31 > Invalid Args. Please specify three Flag options.")
	}

	//引数確認
	fmt.Printf("Input image path：%s, Output image path：%s\n", *src, *dst)

	//args := flag.Args()
	//ここ一つの関数にまとめることはできるのか？
	//ファイル形式のチェック
	str, err_status := converter.CheckSrcImageType(*src)
	if err_status != 0 {
		fmt.Printf("main.go 42 > invalid %s ErrorContent:%s\n", str, err_status)
		os.Exit(1)
	}

	//ファイル形式のチェック
	str, err_status = converter.CheckDstImageType(*dst)
	if err_status != 0 {
		fmt.Printf("main.go 48 > invalid %s ErrorContent:%s\n", dst, err_status)
		os.Exit(1)
	}
	//画像の変換処理
	err := converter.ConvertToPNG(*dir, *src, *dst)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(0)
}
