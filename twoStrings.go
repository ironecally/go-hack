package main

import(
	"strings"
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

	//can use several

	//1st way
	// for i:=0; i<len(str1); i++ {

	// 	for j:=0; j<len(str2); j++ {
	// 		if str1[i] == str2[j] {
	// 			return true
	// 		}	
	// 	}
	// }

	// return false

	//2nd way
	// return strings.ContainsAny(str1,str2)

	//3rd way
	//FASTEST!, number of loop is limited to alphabet
	for i:='a'; i<='z'; i++ {
		if strings.Contains(str1, string(i)) && strings.Contains(str2, string(i)) {
			return true
		}
	}
	return false


}