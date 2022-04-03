package converter

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

const (
	ExitSuccess int = iota
	ExitError
	ExitFileError
	ExitFileTypeError
)

type ImageTypeError string

func (e ImageTypeError) Error() string {
	return string(e)
}

func CheckSrcImageType(src string) (string, int) {
	err_message := "converter.go: CheckSrcImageType - File Type Error.\n"
	switch src {
	case ".jpg", ".jpeg":
		return src, nil
	default:
		fmt.Fprintln(os.Stderr, err_message)
		return nil, ExitFileTypeError
	}
}

func CheckDstImageType(dst string) (string, int) {
	err_message := "converter.go: CheckDstImageType - File Type Error.\n"
	switch dst {
	case ".png":
		return dst, nil
	default:
		fmt.Fprintln(os.Stderr, err_message)
		return nil, ExitFileTypeError
	}
}

func ConvertToPNG(dir string, src string, dst string) int {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Printf("converter.go 48 > filepath error.\n")
			return err
		}

		//画像を開く
		file, err := os.Open(src)
		if err != nil {
			fmt.Printf("converter.go 56 > File open error.\n")
			return err
		}

		defer file.Close()
		if err := file.Sync(); err != nil {
			fmt.Printf("converter.go 63 > Sync Error.\n")
			return err
		}

		//ファイルオブジェクト→画像オブジェクトに変換
		img_data, _, err := image.Decode(file)
		if err != nil {
			fmt.Printf("converter.go 69 > Decode Error.\n")
			return err
		}

		//pngファイルを生成する
		png_file := dst
		output, err := os.Create(png_file)
		if err != nil {
			fmt.Printf("converter.go 77 > Error: Create output file.\n")
			return err
		}

		defer output.Close()
		if err := output.Sync(); err != nil {
			fmt.Printf("converter.go 84 > Sync Error.\n")
			return err
		}

		//エンコード
		err = png.Encode(output, img_data)
		if err != nil {
			fmt.Printf("converter.go 90 > Sync Error.\n")
			return err
		}
		return nil
	})

	if err != nil {
		fmt.Println("converter.go 47 > Convert Error. The following are the details.")
		fmt.Fprintln(os.Stderr, err)
		return ExitError
	}

	return ExitSuccess
}
