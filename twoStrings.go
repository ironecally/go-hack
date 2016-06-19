package main

import(
	"fmt"
)

func main() {
	
	var t int
	fmt.Scanf("%d", &t)

	for i:=0; i<t; i++ {
		yesNo := processString()
		
		if yesNo {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func processString() bool {
	var str1,str2 string

	fmt.Scanf("%s", &str1)
	fmt.Scanf("%s", &str2)

	for i:=0; i<len(str1); i++ {

		for j:=0; j<len(str2); j++ {
			if str1[i] == str2[j] {
				return true
			}	
		}
	}

	return false
}