package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	flag.Parse()
	fmt.Printf("受け取った引数の確認：%v\n", flag.Args())
	converter.convert_img.convertToPNG(flag.Args(0), flag.Args(1))
	fmt.Printf("Hello World Time: %v\n", time.Now())
}
