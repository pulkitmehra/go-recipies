package simple

import (
	"fmt"
	"strings"
	"unicode"
)

//StringMain func
func StringMain() {
	stringLiteral()
	stringRune()
	stringstrimnFn()
	checkUnicode()
}

func stringLiteral() {
	const utf8 = "ab12AB"
	fmt.Printf("% x", utf8)

	unicode := "€£³B"
	fmt.Printf("\n % x \n", unicode)
}

func stringRune() {

	const runeUnicode = '€'
	fmt.Printf("(decimal-hex) %d (bas10): \n", runeUnicode)
	fmt.Printf("(HEX) %x (bas16)\n", runeUnicode)
	fmt.Printf("(binary) %b (bas2)\n", runeUnicode)
	fmt.Printf("(Rune) %q \n", runeUnicode)
	fmt.Printf("(unicode) %U \n", runeUnicode)
	fmt.Printf("(char) %c \n", runeUnicode)

	const utf8 = 'B'
	fmt.Printf("(decimal-hex) %d (bas10): \n", utf8)
	fmt.Printf("(HEX) %x (bas16)\n", utf8)
	fmt.Printf("(binary) %b (bas2)\n", utf8)
	fmt.Printf("(Rune) %q \n", utf8)
	fmt.Printf("(unicode) %U \n", utf8)
	fmt.Printf("(char) %c \n", utf8)

	aString := []byte("pulkit")
	for x, y := range aString {
		fmt.Printf("Char:%c byte:%b string:%v \n", aString[x], aString[x], y)
	}
}
func checkUnicode() {
	const sL = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	for i, x := range sL {
		if unicode.IsPrint(x) {
			fmt.Printf("%c\n", sL[i])
		} else {
			fmt.Println("Not printable!")
		}
	}
}

func stringstrimnFn() {
	trimFn := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	fmt.Println(strings.TrimFunc("123 abc \t \n def", trimFn))
}
