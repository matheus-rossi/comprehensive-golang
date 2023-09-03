package main

import "fmt"

const ebulicaoF float64 = 212.0

func main() {
	tempC := (ebulicaoF - 32) * 5 / 9
	fmt.Println("O ponto de ebulição da água é de", ebulicaoF, "ºF ou", tempC, "ºC")
}
