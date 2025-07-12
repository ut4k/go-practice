package main

import "fmt"

func store() func(int) int {
	var x int
	return func(n int) int {
		x += n
		return x
	}
}

func main() {
	s1 := store()
	s2 := store()
	// クロージャがs1に束縛されている限りs1の中のxは消えない。だから+1するとxは前の値+1になる。
	fmt.Printf("s1: %d, s2: %d\n", s1(1), s2(2))
	fmt.Printf("s1: %d, s2: %d\n", s1(1), s2(2))
	fmt.Printf("s1: %d, s2: %d\n", s1(1), s2(2))
}
