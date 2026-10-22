package projetos

import "fmt"

func fizzBuzz() {
	for num := 1; num <= 100; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("pinpan")
		} else if num%3 == 0 {
			fmt.Println("pin")
		} else if num%5 == 0 {
			fmt.Println("pan")
		} else {
			fmt.Println(num)
		}
	}
}
