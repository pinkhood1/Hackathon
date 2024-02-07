package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for k := '9'; k >= '0'; k-- {
				for l := '9'; l >= '0'; l-- {
					if i > k || (i == k && j > l) {
						Printpair(i, j)
						z01.PrintRune(' ')
						Printpair(k, l)
						if !(i == '0' && j == '1' && k == '0' && l == '0') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}

func Printpair(a, b rune) {
	z01.PrintRune(a)
	z01.PrintRune(b)
}
