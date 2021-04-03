package utils

import (
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// ISO8859_1toUTF8 convert a ISO8859 string into UTF-8 string
func ISO8859_1toUTF8(in string) (out string, err error) {
	// create a Reader using the input string as Reader and ISO8859 decoder
	decoded := transform.NewReader(strings.NewReader(in), charmap.ISO8859_1.NewDecoder())
	decodedBytes, err := ioutil.ReadAll(decoded)
	if err != nil {
		return
	}
	out = string(decodedBytes)
	return
}
