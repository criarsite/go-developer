package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	a, b := 10.0, 2.0
	resultadoDivisao, err := Dividir(a, b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Calculadora básica")
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, Somar(a, b))
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, Subtrair(a, b))
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, Multiplicar(a, b))
	fmt.Printf("%.2f / %.2f = %.2f\n", a, b, resultadoDivisao)

	if _, err := Dividir(a, 0); errors.Is(err, ErrDivisaoPorZero) {
		fmt.Println("Divisão por zero: erro tratado corretamente")
	}
}
