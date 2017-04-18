package methinf // import "github.com/workspace7/tourofgo/methinf"

import (
	"bytes"
	"io"
)

var asciiUppercase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var asciiLowercase = []byte("abcdefghijklmnopqrstuvwxyz")

const lenAlpha = 26

//Rot13Reader reader used to read rot13 bytes
type Rot13Reader struct {
	//R the reader
	R io.Reader
}

//Rot13 return the ROT13 byte
func Rot13(b byte) byte {
	pos := bytes.IndexByte(asciiUppercase, b)
	if pos != -1 {
		return asciiUppercase[(pos+13)%lenAlpha]
	}
	pos = bytes.IndexByte(asciiLowercase, b)
	if pos != -1 {
		return asciiLowercase[(pos+13)%lenAlpha]
	}
	return b
}

//Read read and convert the byte array based on ROT13 algorothm
func (r Rot13Reader) Read(p []byte) (int, error) {
	n, err := r.R.Read(p)
	for i := 0; i < n; i++ {
		p[i] = Rot13(p[i])
	}
	return n, err
}
