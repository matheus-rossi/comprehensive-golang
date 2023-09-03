package main

import "fmt"

const pontoEbulicaoAguaKelvin float64 = 373.0

func main() {
	pontoEbulicaoAguaCelsius := pontoEbulicaoAguaKelvin - 273.0
	fmt.Printf("O resultado é: %v \n", pontoEbulicaoAguaCelsius)
}
