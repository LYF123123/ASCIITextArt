package asciitextart_test

import (
	asciitextart "ASCIITextArt"
	"testing"
)

func TestToAsciiTextArt(t *testing.T) {
	res1:=asciitextart.ToAsciiTextArt("welcome")
	t.Log(res1)
	res2:=asciitextart.ToAsciiTextArt("Goodbye")
	t.Log(res2)
	res3:=asciitextart.ToAsciiTextArt("1234567890")
	t.Log(res3)
	res:=asciitextart.ToAsciiTextArt("hello, 世界")
	if res==""{
		t.Log("OK!")
	}else{
		t.Fatal("Error!")
	}
}
