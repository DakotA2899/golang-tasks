package main

import "fmt"

func CaesarCipher(text string, shift int) {
	/*alph := "abcdefghijklmnopqrstuvwxyz"
	alph1 := "ABCDETFGHIJKLMNOPQRSTUVWXYZ"
	alphMass := []rune(alph1)
	fmt.Print(alphMass)
	[65 66 67 68 69 84 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90]
	[97 98 99 100 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116 117 118 119 120 121 122]
	*/
	for _, r := range text {
		n := int(r) + shift
		switch {
		case (n >= 65 && n <= 90) || (n >= 97 && n <= 122):
			fmt.Printf("%c", rune(n))
		case n < 65:
			k := n - 65
			fmt.Printf("%c", rune(91+k))
		case n < 97:
			k := n - 97
			fmt.Printf("%c", rune(123+k))
		case n > 90:
			k := n - 90
			fmt.Printf("%c", rune(64+k))
		case n > 122:
			k := n - 122
			fmt.Printf("%c", rune(96+k))

		}
	}
}

func main() {
	CaesarCipher("AdmEoWgmrGz", -3)
}
