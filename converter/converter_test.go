//画像をjpg - pngへ変換するためのパッケージです(テスト)
package converter

import (
	//"errors"
	"testing"
)

//type ImageTypeError string

//func (e ImageTypeError) Error() string {
//	return string(e)
//}

//入力画像ファイル形式の確認
func TestCheckSrcImageType(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "test: .jpg", input: ".jpg", expected: 0},
		{name: "test: .jpeg", input: ".jpeg", expected: 0},
		{name: "test: .", input: ".", expected: 3},
		{name: "test: .png", input: ".png", expected: 3},
		{name: "test: .gif", input: ".gif", expected: 3},
		{name: "test: .csv", input: ".csv", expected: 3},
		{name: "test: ./", input: "./", expected: 3},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if _, actual := CheckSrcImageType(c.input); c.expected != actual {
				t.Errorf(
					"want CheckSrcImageType(%s) = %d, got %d",
					c.input, c.expected, actual)
			}
		})
	}
}

//出力画像ファイル形式の確認
func TestCheckDstImageType(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "test: .png", input: ".png", expected: 0},
		{name: "test: .jpg", input: ".jpg", expected: 3},
		{name: "test: .jpeg", input: ".jpeg", expected: 3},
		{name: "test: .", input: ".", expected: 3},
		{name: "test: .gif", input: ".gif", expected: 3},
		{name: "test: .csv", input: ".csv", expected: 3},
		{name: "test: ./", input: "./", expected: 3},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if _, actual := CheckDstImageType(c.input); c.expected != actual {
				t.Errorf(
					"want CheckSrcImageType(%s) = %d, got %d",
					c.input, c.expected, actual)
			}
		})
	}
}

/*
//画像形式を変換する
func TestConvertToPNG(t *testing.T) {
	cases := []struct {
		name     string
		dir      string
		src      string
		dst      string
		expected error
	}{
		{name: "test: dir images/image01.jpeg", dir: "images/", src: "images/sample01.jpeg", dst: "samplegrgge.png", expected: nil},
		{name: "test: dir im/", dir: "im", src: "sample01.jpeg", dst: "sample01.png", expected: errors.New("")},
		{name: "test: src saojfe.png", dir: "images/", src: "saojfe.png", dst: "sample01.png", expected: errors.New("")},
		{name: "test: src saojfe.png", dir: "images/", src: "saojfe.png", dst: "sample01.png", expected: errors.New("")},
		{name: "test: src saojfe.png", dir: "images/", src: "saojfe.png", dst: "sample01.png", expected: errors.New("")},
		{name: "test: src saojfe.png", dir: "images/", src: "saojfe.png", dst: "sample01.png", expected: errors.New("")},
		{name: "test: src saojfe.png", dir: "images/", src: "saojfe.png", dst: "sample01.png", expected: errors.New("")},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := ConvertToPNG(c.dir, c.src, c.dst); c.expected != actual {
				t.Errorf(
					"want ConvertToPng(%s, %s, %s) = %v, got %v",
					c.dir, c.src, c.dst, c.expected, actual)
			}
		})
	}
}*/
