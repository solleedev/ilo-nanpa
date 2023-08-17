package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var input string

	fmt.Println("=============================")
	fmt.Println("toki pona number converter")
	fmt.Println("made by jan Soli for learning Go")
	fmt.Println("=============================")
	fmt.Print("o pana e nanpa sina: ")
	fmt.Scanln(&input)
	fmt.Println("=============================")

	number, _e := strconv.Atoi(input)
	if _e != nil {
		log.Fatal("[IKE] ni li nanpa ala!")
	}

	// break it down into its constituent parts
	hundreds := number / 100
	moduloH := number % 100

	twenties := moduloH / 20
	moduloT := moduloH % 20

	fives := moduloT / 5
	moduloF := moduloT % 5

	twos := moduloF / 2
	ones := moduloF % 2 // guaranteed to be either ala or wan

	fmt.Println(fmt.Sprint(hundreds, "x100 + ", twenties, "x20 + ", fives, "x5 + ", twos, "x2 + ", ones))

	ales := strings.Repeat("ale ", hundreds)
	mutes := strings.Repeat("mute ", twenties)
	lukas := strings.Repeat("luka ", fives)
	tus := strings.Repeat("tu ", twos)
	wan := strings.Repeat("wan", ones)

	nanpa := ales + mutes + lukas + tus + wan
	fmt.Println(nanpa)
}
