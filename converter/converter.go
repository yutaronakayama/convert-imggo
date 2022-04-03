package converter

import (
    "fmt"
    "image"
    "image/jpeg"
    _ "image/png"
    "io"
    "os"
)

func ToJPEG(in io.Reader, out io.Writer) error {
    img, kind, err := image.Decode(in)
    if err != nil {
        return err
    }
    fmt.Fprintln(os.Stderr, "Input format =", kind)
    return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
