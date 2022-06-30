package asciitextart

// Uppercase and lowercase letters and numbers

// Lowercase letters
// a~z
var lowercase = [...]string{
	//a
	"_________",
	"_________",
	"__$$$$___",
	"_$$__$$__",
	"_____$$__",
	"__$$$$$__",
	"_$$__$$__",
	"_$$__$$__",
	"__$$$_$$_",
	"_________",
	// b
	"________",
	"________",
	"_$$_____",
	"_$$_____",
	"_$$_____",
	"_$$$$$__",
	"_$$__$$_",
	"_$$__$$_",
	"_$$$$$__",
	"________",
	// c
	"________",
	"________",
	"__$$$$__",
	"_$$__$$_",
	"_$$_____",
	"_$$_____",
	"_$$_____",
	"_$$__$$_",
	"__$$$$__",
	"________",
	// d
	"________",
	"________",
	"_____$$_",
	"_____$$_",
	"_____$$_",
	"__$$$$$_",
	"_$$__$$_",
	"_$$__$$_",
	"__$$$$$_",
	"________",
	// e
	"________",
	"________",
	"__$$$$__",
	"_$$__$$_",
	"_$$__$$_",
	"_$$$$$$_",
	"_$$_____",
	"_$$_____",
	"__$$$$__",
	"________",
	// f
	"________",
	"________",
	"____$$$_",
	"___$$___",
	"_$$$$$$_",
	"___$$___",
	"___$$___",
	"___$$___",
	"___$$___",
	"________",
	// g
	"________",
	"________",
	"________",
	"__$$$$$_",
	"_$$__$$_",
	"_$$__$$_",
	"__$$$$$_",
	"_____$$_",
	"_____$$_",
	"__$$$$__",
	// h
	"________",
	"________",
	"_$$_____",
	"_$$_____",
	"_$$_____",
	"_$$$$$__",
	"_$$__$$_",
	"_$$__$$_",
	"_$$__$$_",
	"________",
	// i
	"________",
	"________",
	"___$$___",
	"________",
	"__$$$___",
	"___$$___",
	"___$$___",
	"___$$___",
	"_$$$$$$_",
	"________",
	// j
	"_______",
	"_______",
	"____$$_",
	"_______",
	"__$$$$_",
	"____$$_",
	"____$$_",
	"____$$_",
	"____$$_",
	"_$$$$__",
	// k
	"________",
	"________",
	"_$$_____",
	"_$$__$$_",
	"_$$_$$__",
	"_$$$____",
	"_$$$$___",
	"_$$_$$__",
	"_$$__$$_",
	"________",
	// l
	"________",
	"________",
	"_$$$$___",
	"___$$___",
	"___$$___",
	"___$$___",
	"___$$___",
	"___$$___",
	"_$$$$$$_",
	"________",
	// m
	"_________",
	"_________",
	"_$$$$$$__",
	"_$$_$_$$_",
	"_$$_$_$$_",
	"_$$_$_$$_",
	"_$$_$_$$_",
	"_$$_$_$$_",
	"_$$_$_$$_",
	"_________",
	// n
	"_________",
	"_________",
	"_$$$$$$__",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_________",
	// o
	"________",
	"________",
	"__$$$$__",
	"_$$__$$_",
	"_$$__$$_",
	"_$$__$$_",
	"_$$__$$_",
	"_$$__$$_",
	"__$$$$__",
	"________",
	// p
	"________",
	"________",
	"________",
	"_$$$$$__",
	"_$$__$$_",
	"_$$__$$_",
	"_$$$$$__",
	"_$$_____",
	"_$$_____",
	"_$$_____",

	// q
	"________",
	"________",
	"________",
	"__$$$$$_",
	"_$$__$$_",
	"_$$__$$_",
	"__$$$$$_",
	"_____$$_",
	"_____$$_",
	"_____$$_",
	// r
	"__________",
	"__________",
	"_$$_$$$$__",
	"_$$$$__$$_",
	"_$$$______",
	"_$$_______",
	"_$$_______",
	"_$$_______",
	"_$$_______",
	"__________",
	// s
	"________",
	"________",
	"__$$$$__",
	"_$$__$$_",
	"_$$_____",
	"__$$$$__",
	"_____$$_",
	"_$$__$$_",
	"__$$$$__",
	"________",
	// t
	"_________",
	"_________",
	"___$$____",
	"___$$____",
	"_$$$$$$__",
	"___$$____",
	"___$$____",
	"___$$____",
	"____$$$$_",
	"_________",
	// u
	"_________",
	"_________",
	"_$$__$$__",
	"_$$__$$__",
	"_$$__$$__",
	"_$$__$$__",
	"_$$__$$__",
	"_$$__$$__",
	"__$$$_$$_",
	"_________",
	// v
	"________",
	"________",
	"_$$__$$_",
	"_$$__$$_",
	"_$$__$$_",
	"_$$__$$_",
	"_$$__$$_",
	"__$__$__",
	"___$$___",
	"________",
	// w
	"___________",
	"___________",
	"_$$_____$$_",
	"_$$_____$$_",
	"_$$__$__$$_",
	"_$$__$__$$_",
	"__$$_$_$$__",
	"___$$_$$___",
	"___$$_$$___",
	"___________",
	// x
	"___________",
	"___________",
	"_$$_____$$_",
	"__$$___$$__",
	"___$$_$$___",
	"____$$$____",
	"___$$_$$___",
	"__$$___$$__",
	"_$$_____$$_",
	"___________",
	// y
	"__________",
	"__________",
	"__________",
	"_$$____$$_",
	"__$$__$$__",
	"___$$$$___",
	"____$$$___",
	"_____$$___",
	"_____$$___",
	"__$$$$____",
	// z
	"________",
	"________",
	"________",
	"_$$$$$$_",
	"_____$$_",
	"____$$__",
	"___$$___",
	"__$$____",
	"_$$$$$$_",
	"________",
}

// Uppercase letters
// A~Z
var uppercase = [...]string{
	// A
	"____$$____",
	"___$$$$___",
	"__$$__$$__",
	"_$$____$$_",
	"_$$____$$_",
	"_$$$$$$$$_",
	"_$$____$$_",
	"_$$____$$_",
	"_$$____$$_",
	"__________",
	// B
	"_$$$$$$__",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$$$$$__",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$$$$$__",
	"_________",
	// C
	"__$$$$$__",
	"_$$___$$_",
	"_$$___$$_",
	"_$$______",
	"_$$______",
	"_$$______",
	"_$$___$$_",
	"_$$___$$_",
	"__$$$$$__",
	"_________",
	// D
	"_$$$$$___",
	"_$$__$$__",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$__$$__",
	"_$$$$$___",
	"_________",
	// E
	"_$$$$$$$_",
	"_$$______",
	"_$$______",
	"_$$______",
	"_$$$$$$$_",
	"_$$______",
	"_$$______",
	"_$$______",
	"_$$$$$$$_",
	"_________",
	// F
	"_$$$$$$$_",
	"_$$______",
	"_$$______",
	"_$$______",
	"_$$$$$$$_",
	"_$$______",
	"_$$______",
	"_$$______",
	"_$$______",
	"_________",
	// G
	"__$$$$$__",
	"_$$___$$_",
	"_$$___$$_",
	"_$$______",
	"_$$__$$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"__$$$$$__",
	"_________",
	// H
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$$$$$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_________",
	// I
	"_$$$$$$_",
	"___$$___",
	"___$$___",
	"___$$___",
	"___$$___",
	"___$$___",
	"___$$___",
	"___$$___",
	"_$$$$$$_",
	"________",
	// J
	"__$$$$$$_",
	"______$$_",
	"______$$_",
	"______$$_",
	"______$$_",
	"______$$_",
	"_$$___$$_",
	"_$$___$$_",
	"__$$$$$__",
	"_________",
	// K
	"_$$___$$_",
	"_$$__$$__",
	"_$$_$$___",
	"_$$$$____",
	"_$$$_____",
	"_$$$$____",
	"_$$_$$___",
	"_$$__$$__",
	"_$$___$$_",
	"_________",
	// L
	"_$$______",
	"_$$______",
	"_$$______",
	"_$$______",
	"_$$______",
	"_$$______",
	"_$$______",
	"_$$______",
	"_$$$$$$$_",
	"_________",
	// M
	"_$$________$$_",
	"_$$$$____$$$$_",
	"_$$_$$__$$_$$_",
	"_$$__$$$$__$$_",
	"_$$___$$___$$_",
	"_$$________$$_",
	"_$$________$$_",
	"_$$________$$_",
	"_$$________$$_",
	"______________",
	// N
	"_$$_______$$_",
	"_$$$$_____$$_",
	"_$$_$$____$$_",
	"_$$__$$___$$_",
	"_$$___$$__$$_",
	"_$$____$$_$$_",
	"_$$_____$$$$_",
	"_$$______$$$_",
	"_$$_______$$_",
	"_____________",
	// O
	"__$$$$$$__",
	"_$$____$$_",
	"_$$____$$_",
	"_$$____$$_",
	"_$$____$$_",
	"_$$____$$_",
	"_$$____$$_",
	"_$$____$$_",
	"__$$$$$$__",
	"__________",
	// P
	"_$$$$$$__",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"_$$$$$$__",
	"_$$______",
	"_$$______",
	"_$$______",
	"_$$______",
	"_________",
	// Q
	"__$$$$$$___",
	"_$$____$$__",
	"_$$____$$__",
	"_$$____$$__",
	"_$$____$$__",
	"_$$____$$__",
	"_$$____$$__",
	"_$$____$$__",
	"__$$$$$$___",
	"______$$$$_",
	// R
	"_$$$$$$___",
	"_$$___$$__",
	"_$$___$$__",
	"_$$___$$__",
	"_$$$$$$___",
	"_$$_$$____",
	"_$$__$$___",
	"_$$___$$__",
	"_$$____$$_",
	"__________",
	// S
	"__$$$$$$__",
	"_$$____$$_",
	"_$$____$$_",
	"_$$_______",
	"__$$$$$$__",
	"_______$$_",
	"_$$____$$_",
	"_$$____$$_",
	"__$$$$$$__",
	"__________",
	// T
	"_$$$$$$$$_",
	"____$$____",
	"____$$____",
	"____$$____",
	"____$$____",
	"____$$____",
	"____$$____",
	"____$$____",
	"____$$____",
	"__________",
	// U
	"_$$____$$___",
	"_$$____$$___",
	"_$$____$$___",
	"_$$____$$___",
	"_$$____$$___",
	"_$$____$$___",
	"_$$____$$___",
	"_$$____$$___",
	"__$$$$$$_$$_",
	"____________",
	// V
	"_$$______$$_",
	"_$$______$$_",
	"_$$______$$_",
	"_$$______$$_",
	"_$$______$$_",
	"__$$____$$__",
	"___$$__$$___",
	"____$$$$____",
	"_____$$_____",
	"____________",
	// W
	"_$$_________$$_",
	"_$$_________$$_",
	"_$$___$$____$$_",
	"_$$___$$____$$_",
	"_$$___$$____$$_",
	"__$$_$$_$$_$$__",
	"___$$$___$$$___",
	"___$$_____$$___",
	"___$$_____$$___",
	"_______________",
	// X
	"_$$______$$_",
	"__$$____$$__",
	"___$$__$$___",
	"____$$$$____",
	"_____$$_____",
	"____$$$$____",
	"___$$__$$___",
	"__$$____$$__",
	"_$$______$$_",
	"____________",
	// Y
	"_$$______$$_",
	"__$$____$$__",
	"___$$__$$___",
	"____$$$$____",
	"_____$$_____",
	"_____$$_____",
	"_____$$_____",
	"_____$$_____",
	"_____$$_____",
	"____________",
	// Z
	"_$$$$$$$$_",
	"_______$$_",
	"______$$__",
	"_____$$___",
	"____$$____",
	"___$$_____",
	"__$$______",
	"_$$_______",
	"_$$$$$$$$_",
	"__________",
}

// Number
// 0~9
var numbercase = [...]string{
	// 0
	"__________",
	"__$$$$$$__",
	"_$$___$$$_",
	"_$$__$$$$_",
	"_$$_$$_$$_",
	"_$$$$__$$_",
	"_$$$___$$_",
	"_$$____$$_",
	"__$$$$$$__",
	"__________",
	// 1
	"__________",
	"___$$$____",
	"__$$$$____",
	"_$$$$$____",
	"____$$____",
	"____$$____",
	"____$$____",
	"____$$____",
	"_$$$$$$$$_",
	"__________",
	// 2
	"__________",
	"__$$$$$$__",
	"_$$____$$_",
	"_$$___$$__",
	"_____$$___",
	"____$$____",
	"___$$_____",
	"__$$______",
	"_$$$$$$$$_",
	"__________",
	// 3
	"__________",
	"__$$$$$$__",
	"_$$____$$_",
	"_______$$_",
	"___$$$$$$_",
	"_______$$_",
	"_______$$_",
	"_$$____$$_",
	"__$$$$$$__",
	"__________",
	// 4
	"____________",
	"_____$$$$___",
	"____$$_$$___",
	"___$$__$$___",
	"__$$___$$___",
	"_$$$$$$$$$$_",
	"_______$$___",
	"_______$$___",
	"_______$$___",
	"____________",
	// 5
	"_________",
	"_$$$$$$$_",
	"_$$______",
	"_$$______",
	"_$$$$$$__",
	"______$$_",
	"______$$_",
	"_$$___$$_",
	"__$$$$$__",
	"_________",
	// 6
	"_________",
	"__$$$$$__",
	"_$$___$$_",
	"_$$______",
	"_$$$$$$__",
	"_$$___$$_",
	"_$$___$$_",
	"_$$___$$_",
	"__$$$$$__",
	"_________",
	// 7
	"_________",
	"_$$$$$$$_",
	"______$$_",
	"______$$_",
	"______$$_",
	"______$$_",
	"______$$_",
	"______$$_",
	"______$$_",
	"_________",
	// 8
	"__________",
	"__$$$$$$__",
	"_$$____$$_",
	"_$$____$$_",
	"_$$$$$$$$_",
	"_$$____$$_",
	"_$$____$$_",
	"_$$____$$_",
	"__$$$$$$__",
	"__________",
	// 9
	"__________",
	"__$$$$$$__",
	"_$$____$$_",
	"_$$____$$_",
	"__$$$$$$$_",
	"_______$$_",
	"_______$$_",
	"_______$$_",
	"__$$$$$$__",
	"__________",
}
