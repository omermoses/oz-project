package main

type City struct {
	name                string
	numberOfPostOffices int
	TotalNumOfPacks     int
	postOffices         map[string]*PostOffice
}

func CreateCity() *City {
	cityName := getStringUserInput("please select a name for a city:")
	numOfOffices := getIntUserInput("please select the number of post offices:")

	newCity := &City{
		name:                cityName,
		numberOfPostOffices: numOfOffices,
		TotalNumOfPacks:     0,
		postOffices:         make(map[string]*PostOffice),
	}
	return newCity
}

func repr(city *City) {

}
