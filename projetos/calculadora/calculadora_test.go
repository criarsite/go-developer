package main

import (
	"errors"
	"testing"
)

func TestSomar(t *testing.T) {
	if got := Somar(10, 2); got != 12 {
		t.Errorf("Somar(10, 2) = %v; esperado 12", got)
	}
}

func TestSubtrair(t *testing.T) {
	if got := Subtrair(10, 2); got != 8 {
		t.Errorf("Subtrair(10, 2) = %v; esperado 8", got)
	}
}

func TestMultiplicar(t *testing.T) {
	if got := Multiplicar(10, 2); got != 20 {
		t.Errorf("Multiplicar(10, 2) = %v; esperado 20", got)
	}
}

func TestDividir(t *testing.T) {
	got, err := Dividir(10, 2)
	if err != nil {
		t.Fatalf("Dividir(10, 2) retornou erro inesperado: %v", err)
	}
	if got != 5 {
		t.Errorf("Dividir(10, 2) = %v; esperado 5", got)
	}
}

func TestDividirPorZero(t *testing.T) {
	_, err := Dividir(10, 0)
	if !errors.Is(err, ErrDivisaoPorZero) {
		t.Errorf("Dividir(10, 0) erro = %v; esperado %v", err, ErrDivisaoPorZero)
	}
}
