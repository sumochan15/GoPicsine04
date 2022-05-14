package piscine

import (
	"ft"
	"os"
) 

func slicecount(str []string) (count int) {
	for range str {
		count++
	}
	return
}

func printstr(str string) {
	for _, r := range str {
		ft.PrintRune(r)
	}
}

func Printparams() {
	str := os.Args
	count := slicecount(str)

	for i := 0; i < count; i++ {
		printstr(str[i])
		ft.PrintRune('\n')
	}
}