//画像をjpg - pngへ変換するためのパッケージです
package converter

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

const (
	ExitSuccess int = iota
	ExitError
	ExitFileError
	ExitFileTypeError
)

//type ImageTypeError string

//func (e ImageTypeError) Error() string {
//	return string(e)
//}

//入力画像ファイル形式の確認
func CheckSrcImageType(src string) (string, int) {
	err_message := "converter.go: CheckSrcImageType - File Type Error."
	e := filepath.Ext(src)
	switch e {
	case ".jpg", ".jpeg":
		return ".jpg .jpeg", ExitSuccess
	default:
		fmt.Println(err_message)
		return e, ExitFileTypeError
	}
}

//出力画像ファイル形式の確認
func CheckDstImageType(dst string) (string, int) {
	err_message := "converter.go: CheckDstImageType - File Type Error."
	e := filepath.Ext(dst)
	switch e {
	case ".png":
		return ".png", ExitSuccess
	default:
		fmt.Println(err_message)
		return e, ExitFileTypeError
	}
}

//拡張子を除去して画像名を取得する
func getFileNameWithoutExt(path string) string {
	// Fixed with a nice method given by mattn-san
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

//画像形式を変換する
func ConvertToPNG(dir string, src string, dst string) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		//ディレクトリの場合処理を抜ける
		if info.IsDir() {
			return err
		}
		//変換処理を行う画像パス
		fmt.Println(path)
		if _, status := CheckSrcImageType(filepath.Ext(path)); status == 0 {
			if err != nil {
				fmt.Printf("converter.go > filepath error.\n")
				return err
			}

			//画像を開く
			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("converter.go > File open error.\n")
				return err
			}

			defer file.Close()
			if err := file.Sync(); err != nil {
				fmt.Printf("converter.go > Sync Error.\n")
				return err
			}

			//ファイルオブジェクト→画像オブジェクトに変換
			img_data, err := jpeg.Decode(file)
			if err != nil {
				fmt.Printf("converter.go > Decode Error.\n %s \n", err)
				return err
			}

			//pngファイルを生成する
			var png_file = ""
			if dir == "" {
				png_file = dst
			} else {
				png_file = "images/" + getFileNameWithoutExt(path) + ".png"
			}
			fmt.Println(png_file)
			path = ""
			output, err := os.Create(png_file)
			if err != nil {
				path = ""
				fmt.Printf("converter.go > Error: Create output file.\n %s \n", err)
				return err
			}

			defer output.Close()
			if err := output.Sync(); err != nil {
				fmt.Printf("converter.go 84 > Sync Error.\n %s \n", err)
				return err
			}

			//エンコード
			err = png.Encode(output, img_data)
			if err != nil {
				fmt.Printf("converter.go 90 > Sync Error.\n %s \n", err)
				return err
			}

			//コンバートが完了した旨を通知
			fmt.Printf("Convert completed: %s", path)

		} else {
			fmt.Printf(">> %s", path)
		}
		return err
	})
	if err != nil {
		fmt.Println("converter.go > Convert Error. The following are the details.")
		return err
	}
	return err
}
