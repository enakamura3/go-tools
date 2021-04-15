package http

import (
	"io"
	"net/http"
	"os"
)

// downloadFile download a file from an url
func DownloadFile(url, filePath string) (size int64, err error) {
	rsp, err := http.Get(url)
	if err != nil {
		return
	}
	defer rsp.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	size, err = io.Copy(file, rsp.Body)
	if err != nil {
		return
	}
	return
}
