package asciitextart_test

import (
	asciitextart "ASCIITextArt"
	"testing"
)

func TestToAsciiTextArt(t *testing.T) {
	res1:=asciitextart.ToAsciiTextArt("welcome")
	t.Log(res1)
	res2:=asciitextart.ToAsciiTextArt("hello, 世界")
	if res2==""{
		t.Log("OK!")
	}else{
		t.Fatal("Error!")
	}
}
