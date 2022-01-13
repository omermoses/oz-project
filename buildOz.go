package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func BuildNewOz() *Oz {
	var cities = 1
	newOz := CreateOz()

	for cities <= newOz.NumberOfCities {
		newCity := BuildCity()
		newOz.Cities[newCity.name] = newCity
		UpdateMaxPackNumCity(newOz, newCity, newOz.Cities[newOz.MaxNumPackCity])
		cities++
	}

	return newOz
}

func BuildCity() *City {
	var posts = 1
	newCity := CreateCity()

	for posts <= newCity.numberOfPostOffices {
		newPost := BuildPostOffice(newCity)
		newCity.postOffices[strconv.Itoa(posts-1)] = newPost
		posts++
	}

	return newCity
}

func BuildPostOffice(cityHost *City) *PostOffice {
	var packages = 1
	newOffice := CreatePostOffice(cityHost)

	for packages <= newOffice.numOfPack {
		newPack := CreatePack()
		newOffice.mailPackage = append(newOffice.mailPackage, newPack)
		fmt.Println(newPack.weight)
		packages++
	}

	return newOffice
}

/*func BuildNewOz() *structures.Oz {
	var newOz *structures.Oz
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Number Of cities in Oz: ")
	scanner.Scan()
	userInput := scanner.Text()
	citiesNumber, err := strconv.Atoi(userInput)

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return nil
	}

	if (citiesNumber) != 0 {
		newOz = structures.CreateOz(citiesNumber)
		CreateCities(newOz, citiesNumber)
	} else {
		fmt.Println(err)
		return nil
	}

	return nil
}

func CreateCities(oz *structures.Oz, numberOfCities int) {
	var createdCitis = 0
	var numberOfPosts int
	scanner := bufio.NewScanner(os.Stdin)

	for {
		var newCity *structures.City
		fmt.Print("Enter new city name: ")
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 {
			newCity = structures.CreateNewCity(text)
			fmt.Println("Enter the required number of post offices in this city")
			numberOfPosts = getUserIntInput()
			oz.Cities[text] = newCity
			CreatePostOffices(newCity, numberOfPosts)
			createdCitis++
		} else {
			// exit if user entered an empty string
			break
		}
		if createdCitis == numberOfCities {
			break
		}
	}

	fmt.Println("Finished all cities")
}

func CreatePostOffices(city *structures.City, numberOfOffices int) {
	var createdOffices = 0
	scanner := bufio.NewScanner(os.Stdin)

	for {
		var newPost *structures.PostOffice
		fmt.Print("Enter new post office name: ")
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 {
			newPost = structures.CreateNewPostOffice(text)

			city.PostOffices[text] = newPost
			createdOffices++
		} else {
			// exit if user entered an empty string
			break
		}
		if createdOffices == numberOfOffices {
			break
		}
	}

	fmt.Println("Finished all posts")
}

func getUserIntInput() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()
	number, err := strconv.Atoi(userInput)

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return -1
	}

	return number
*/

func iterateFields(v interface{}) {
	valueOf := reflect.ValueOf(v)
	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Field(i)
		if field.Kind() == reflect.Struct {
			iterateFields(field.Interface())
			continue
		}
		fmt.Println(field.String())
	}
}
