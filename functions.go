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
		fmt.Println("Enter the action you want to perform")
		scanner.Scan()
		in := scanner.Text()
		userInput := strings.Split(in, " ")
		x, err := strconv.Atoi(userInput[0])
		if err != nil {
			fmt.Println("ERROR:", err)
			fmt.Print("\nPlease enter a valid number: ")
		}
		if x == 1 {

		}
		if x == 2 {
			MakeCityTransfer(ozCity, userInput)
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

func MakeCityTransfer(ozCity *Oz, userInput []string) {
	sourceCity, sourcePostOffice, destCity, destPostOffice := userInput[1], userInput[2], userInput[3], userInput[4]

	transferPacks(
		ozCity.Cities[sourceCity],
		ozCity.Cities[sourceCity].postOffices[sourcePostOffice],
		ozCity.Cities[destCity],
		ozCity.Cities[destCity].postOffices[destPostOffice],
	)
}
