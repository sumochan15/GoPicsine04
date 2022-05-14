package piscine

import (
	"ft"
	"os"
)

func countslice(str []string) (count int) {
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

func RevParams() {
	args := os.Args
	count := countslice(args)
	for count > 1 {
		printstr(args[count-1])
		ft.PrintRune('\n')
		count--
	}
}
