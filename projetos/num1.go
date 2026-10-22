package projetos

import "fmt"

func projetonumdivididopor3() {
	numero := 100

	for num := 1; num <= numero; num++ {
		if num%3 == 0 {
			fmt.Println("numero =", num)
		}
	}
}
