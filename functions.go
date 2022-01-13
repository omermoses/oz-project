package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func OzActions(ozCity *Oz) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		in := scanner.Text()
		parts := strings.Split(in, " ")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("ERROR:", err)
			fmt.Print("\nPlease enter a valid number: ")
		}
		if x == 1 {

		}
		if x == 2 {

		}
		if x == 3 {
			PrintMaxPackCity(ozCity)
			break
		}
	}
}

func printCity(ozCity *Oz, cityName string) {

}

func PrintMaxPackCity(ozCity *Oz) {
	fmt.Println(ozCity.MaxNumPackCity)
}
