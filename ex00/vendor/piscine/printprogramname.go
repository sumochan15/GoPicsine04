package piscine

import (
	"ft"
	"os"
)

func printstr(str string) {
	for _, v := range str {
		ft.PrintRune(v)
	}
}

func strlen(str string) (len int) {
	for range str {
		len++
	}
	return
}

func strrchr(str string, n rune) string {

	s_len := strlen(str)

	for s_len > 0 {
		if rune(str[s_len-1]) == n {
			return str[s_len:]
		}
		s_len--
	}
	return ""
}

func Printprogramname() {
	printstr(strrchr(os.Args[0], '/'))
	ft.PrintRune('\n')
}