package main

import (
	"io"
	"os"
	"strings"

	"github.com/workspace7/tourofgo/basics"
	"github.com/workspace7/tourofgo/methinf"
)

func main() {

	pName := "rot13" // will be repalced with commands
	switch pName {
	case "arrays_slice":
		{
			basics.ChantWithArray()
			chants := make([]string, 2, 10)
			chants[0] = "Hare Krishna"
			basics.ChantWithSlice(chants)
		}
	case "rot13":
		{
			s := strings.NewReader("I love go lang")
			r := methinf.Rot13Reader{R: s}
			io.Copy(os.Stdout, &r)
		}
	}
}
