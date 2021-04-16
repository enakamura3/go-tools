package file

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Unzip decompress a zip file
func Unzip(src, dst string) (err error) {
	z, err := zip.OpenReader(src)
	if err != nil {
		return
	}
	defer z.Close()

	os.MkdirAll(dst, os.ModePerm)
	for _, zf := range z.File {
		zfo, err := zf.Open()
		if err != nil {
			return err
		}
		defer zfo.Close()

		filePath := filepath.Join(dst, zf.Name)
		df, err := os.Create(filePath)
		if err != nil {
			return err
		}
		_, err = io.Copy(df, zfo)
		if err != nil {
			return err
		}
	}
	return
}
