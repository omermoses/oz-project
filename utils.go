package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getStringUserInput(presentedString string) string {
	var userInput string

	fmt.Println(presentedString)

	n, err := fmt.Scanln(&userInput)

	if n < 0 {
		return ""
	}
	if err != nil {
		return ""
	}
	return userInput
}

func getIntUserInput(presentedString string) int {
	var userInput string
	fmt.Println(presentedString)
	fmt.Scanln(&userInput)

	// string to int
	i, err := strconv.Atoi(userInput)
	if err != nil {
		// handle error
		fmt.Println(err)
		return -1
	}

	return i
}

func getSplittedInput(presentedString string) (int, int, int) {
	fmt.Println(presentedString)
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	in := scanner.Text()
	parts := strings.Split(in, " ")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Print("\nPlease enter a valid number: ")
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Print("\nPlease enter a valid number: ")
	}
	z, err := strconv.Atoi(parts[2])
	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Print("\nPlease enter a valid number: ")
	}

	return x, y, z
}

func getStrIntInput(presentedString string) (string, int) {
	fmt.Println(presentedString)
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	in := scanner.Text()
	parts := strings.Split(in, " ")

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Print("\nPlease enter a valid number: ")
	}

	return parts[0], y

}
