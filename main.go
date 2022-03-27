package main

import (
	"flag"
	"fmt"
	"github.com/yutaronakayama/convert_imggo/converter"
	"time"
)

func main() {
	flag.Parse()
	fmt.Printf("受け取った引数の確認：%v\n", flag.Args())
	converter.ConvertToPNG("image/たかしゃべ.png")
	fmt.Printf("Hello World Time: %v\n", time.Now())
}
