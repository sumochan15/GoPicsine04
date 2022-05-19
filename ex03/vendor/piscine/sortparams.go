package piscine

import (
	"ft"
	"os"
)

func counteleme(str []string)(count int){
	for range str{
		count++
	}
	return
}
func checkascii(str []string)([]string){
	for i,count := 0, counteleme(str); i < count -1; i++{
		for j := i + 1; j < count; j++{
			if str[i] > str[j]{
				str[i],str[j] = str[j],str[i]
			}
		}
	}
	return str
}

func printstr(str string){
	for _, r := range str{
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func SortParams(){
	str := os.Args
	count := counteleme(str)
	if count <= 1 {
		return
	}
	re_str := checkascii(str[1:])
	for _, v := range re_str{
		printstr(v)
	}
}