package asciitextart_test

import (
	asciitextart "github.com/LYF123123/ASCIITextArt"
	"testing"
)

func TestToAsciiTextArt(t *testing.T) {
	res1 := asciitextart.ToAsciiTextArt("welcome")
	t.Log(res1)
	res2 := asciitextart.ToAsciiTextArt("Goodbye")
	t.Log(res2)
	res3 := asciitextart.ToAsciiTextArt("1234567890")
	t.Log(res3)
	res4 := asciitextart.ToAsciiTextArt(" !\"#$%&")
	t.Log(res4)
	res5 := asciitextart.ToAsciiTextArt("'()*+,-./")
	t.Log(res5)
}
