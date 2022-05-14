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
	return str[1:]
}

func printstr(str string){
	for _, r := range str{
		ft.PrintRune(r)
		ft.PrintRune('\n')
	}	
}

func SortParams(){
	str := os.Args
	re_str := checkascii(str)
	count := counteleme(re_str)
	for i := 0; i < count; i++ {
		printstr(re_str[i])
	}
}