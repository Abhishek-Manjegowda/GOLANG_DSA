package pattern

import "fmt"

func NumPattern(num int) {
	for i := 0; i < num; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j+1, " ")
		}
		fmt.Println()
	}
}
