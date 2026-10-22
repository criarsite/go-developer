package estruturadecontrole

import "fmt"

func Estruturas() {
	EstruturaFor()
	EstruturaIFelse()
}

func EstruturaFor() {
	var i = 1
	for i <= 10 {
		fmt.Println("i =", i)
		i++

	}
}

func EstruturaIFelse() {
	var i = 67

	if i <= 18 {
		fmt.Println("menor de idade")
	} else if i > 60 {
		fmt.Println("velhinho demais")
	} else {
		fmt.Println("adulto")
	}

}
