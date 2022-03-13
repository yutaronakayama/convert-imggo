package cvtimg

import (
	"fmt"
	"image"
	"os"
)

//画像のオープン
//画像の読み込み
//画像の変換
//画像の書き出し
//画像のクローズ

func convertToPNG(src string, dst string, output_name string) error {
	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()

	fmt.Printf("%d px", bounds.Dx())
	file.Close()

	return err
}
