package converter

import (
	"image"
	"image/png"
	"image/jpeg"
	"bytes"
	"os"
	"net/http"
)

func ConvertToPNG(image []byte) ([]byte, error) {
	contentType = http.DetectContentType(image)

	switch contentType{

	case "image/png":
		image, err := jpeg.Decode(bytes.NewReader(image))
		if err != nil{
			return nil, errors.Wrap(err, "cannot decode this image type.")
		}
		buf := new(bytes.Buffer)
		if err := png.Encode(buf, img); err != nil {
			return nil, errors.Wrap(err, "cannot encode this image png")
		}
		return buf.Bytes(), nil
	default:
		fmt.Println("cannot convert this file type.")

	}
	return nil, fmt.Errorf("unable to convert this image.")
}
