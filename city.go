package main

type City struct {
	name                string
	numberOfPostOffices int
	postOffices         map[int]*PostOffice
}

func CreateCity() *City {
	cityName := getStringUserInput("please select a name for a city:")
	numOfOffices := getIntUserInput("please select the number of post offices:")

	newCity := &City{
		name:                cityName,
		numberOfPostOffices: numOfOffices,
		postOffices:         make(map[int]*PostOffice),
	}
	return newCity
}

func repr(city *City) {
	
}
