package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Japanese struct {
	Person
	MyNumber int
}

func Hello(p Person) {
	fmt.Println("Hello " + p.Name)
}

func main() {
	p := Person{Name: "Taro", Age: 30}
	Hello(p)

	j := Japanese{
		Person:   Person{Name: "Taro", Age: 30},
		MyNumber: 123456,
	}
	// cannot use j (variable of struct type Japanese) as Person value in argument to Hello
	Hello(j)
}
