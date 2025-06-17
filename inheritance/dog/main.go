package main

import "fmt"

// EmbeddingによるOOPの継承っぽいコード
type Dog struct{}

func (d *Dog) Bark() string { return "Bow" }

type Bulldog struct{ Dog }

type ShibaInu struct{ Dog }

func (s *ShibaInu) Bark() string { return "ワン" }

func DogVoice(d *Dog) string { return d.Bark() }

func main() {
	bd := &Bulldog{}
	fmt.Println(bd.Bark())
	si := &ShibaInu{}
	fmt.Println(si.Bark())

	// cannot use si (variable of type *ShibaInu) as *Dog value in argument to DogVoice
	// fmt.Println(DogVoice(si))
}
