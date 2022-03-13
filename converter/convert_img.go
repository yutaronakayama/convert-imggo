package cvtimg

import (
	"image"
	"image/png"
	"os"
)

//画像のオープン
//画像の読み込み
//画像の変換
//画像の書き出し
//画像のクローズ

func convertToPNG(src string, dst string) error {
	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}

	output, err := os.Open(dst)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	err = png.Encode(output, img)
	if err != nil {
		panic(err)
	}

	file.Close()

	return err
}
