package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	person1 := person{name: "John", age: 20}

	fmt.Printf("Pessoa : %v \n", person1)

}
