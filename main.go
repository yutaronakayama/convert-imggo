package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	flag.Parse()
	fmt.Printf("受け取った引数の確認：%v\n", flag.Args())
	converter.convert_img.convertToPNG()
	fmt.Printf("Hello World Time: %v\n", time.Now())
}
