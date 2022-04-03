package main

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
		src = flag.String("src", "jpg", "please specify -src flag")
		dst = flag.String("dst", "png", "please specify -dst flag")
		dir = flag.String("dir", ".", "please specify -dir flag")
	)

	flag.Parse()

	//オプション引数チェック
	if flag.NArg() != 0 && flag.NFlag() != 2 {
		log.Fatal("main.go: 31 > Invalid Args. Please specify only two Flag options.")
	}

	//引数確認
	fmt.Printf("Input image path：%s, Output image path：%s\n", *src, *dst)

	args := flag.Args()

	//ここ一つの関数にまとめることはできるのか？
	//ファイル形式のチェック
	str, err := converter.Check SrcImageType(*src)
	if err != nil {
		fmt.Printf("main.go 42 > invalid %s ErrorContent:%s\n", *str, err)
	}

	//ファイル形式のチェック
	str, err := converter.CheckDstImageType(*dst)
	if err != nil {
		fmt.Printf("main.go 48 > invalid %s ErrorContent:%s\n", *dst, err)

	}
	//画像の変換処理
	err := converter.ConvertToPNG(*dir, *src, *dst)
	if err != 0{
		return err
	}
}
