package main

import (
	"fmt"
	"os"
)

func main() {
	home := os.Getenv("HOME")
	if home == "" {
		fmt.Println("HOME environment variable is not set")
	} else {
		fmt.Println("HOME:", home)
	}

	home, exists := os.LookupEnv("DBPORT")
	if !exists {
		fmt.Println("DBPORT environment variable is not set")
	} else {
		fmt.Println("DBPORT:", home)
	}
}
