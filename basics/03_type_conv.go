package main

import "fmt"

var a int = 10
var b string = "golang"

func main() {
	var c float64 = float64(a)
	fmt.Printf("O Valor de C é: %v %T \n", c, c)
}
