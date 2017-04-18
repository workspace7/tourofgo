package methinf

import (
	"bytes"
	"io"
	"os"
	"strings"
)

var ascii_uppercase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var ascii_lowercase = []byte("abcdefghijklmnopqrstuvwxyz")

const alpa_len = 26

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	pos := bytes.IndexByte(ascii_uppercase, b)
	if pos != -1 {
		return ascii_uppercase[(pos+13)%alpa_len]
	}
	pos = bytes.IndexByte(ascii_lowercase, b)
	if pos != -1 {
		return ascii_lowercase[(pos+13)%alpa_len]
	}
	return b
}

//Read read and convert the byte array based on ROT13 algorothm
func (r rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, err
}

// func main() {
// 	s := strings.NewReader("Lbh penpxrq gur pbqr!")
// 	r := rot13Reader{s}
// 	io.Copy(os.Stdout, &r)
// }
