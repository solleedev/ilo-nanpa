package main

import (
	"fmt"
	"log"
)

func main() {
	var input string

	fmt.Println("=============================")
	fmt.Println("toki pona number converter")
	fmt.Println("made by jan Soli for learning Go")
	fmt.Println("=============================")

	fmt.Println("sina wile pali e nanpa tan toki Inli la o toki e 'pona'")
	fmt.Println("sina wile pali e nanpa tan toki pona la o toki e 'inli'")
	fmt.Println("=============================")
	fmt.Print("> ")
	fmt.Scanln(&input)

	if input == "pona" {
		tawaPona()
	} else if input == "inli" {
		tawaInli()
	} else {
		log.Fatal("[IKE] ni li anu ala!")
	}
}
