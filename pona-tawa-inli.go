package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func count[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return count
}

func check(input string, value string) int {
	return count(strings.Split(input, " "), func(x string) bool { return strings.Contains(x, value) })
}

func tawaInli() {
	var input string

	fmt.Print("o pana e nanpa pi toki pona: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		input = s
		break
	}
	fmt.Println("=============================")

	// count the number of "ale", "mute", "luka", "tu", "wan"
	ales := check(input, "ale") * 100
	mutes := check(input, "mute") * 20
	lukas := check(input, "luka") * 5
	tu := check(input, "tu") * 2
	wan := check(input, "wan")

	fmt.Println(ales + mutes + lukas + tu + wan)
}
