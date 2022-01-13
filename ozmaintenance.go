package main

import "fmt"

func UpdateMaxPackNumCity(ozCity *Oz, cityA *City, cityB *City) {
	// Checking If we should update the city with max number of packages

	if cityB == nil {
		ozCity.MaxNumPackCity = cityA.name
		return
	}
	if cityA.TotalNumOfPacks > cityB.TotalNumOfPacks {
		ozCity.MaxNumPackCity = cityA.name
		return
	}

}

func transferPacks(
	sourceCity *City,
	sourcePostOffice *PostOffice,
	destCity *City,
	destPostOffice *PostOffice) {

	tempFilteredReceiveList := make([]*Package, sourcePostOffice.numOfPack)
	cantReceiveList := make([]*Package, sourcePostOffice.numOfPack)

	for key, element := range sourcePostOffice.mailPackage {
		fmt.Println(key)
		if element == nil {
			continue
		}
		if element.weight < destPostOffice.max_weight && element.weight > destPostOffice.min_weight {
			tempFilteredReceiveList = append(tempFilteredReceiveList, element)
		} else {
			cantReceiveList = append(cantReceiveList, element)
		}
	}

}
