package asciitextart_test

import (
	asciitextart "ASCIITextArt"
	"testing"
)

func TestToAsciiTextArt(t *testing.T) {
	res:=asciitextart.ToAsciiTextArt("welcome")
	t.Log(res)
}
