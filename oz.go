package main

type Oz struct {
	NumberOfCities int
	Cities         map[string]*City
}

func CreateOz() *Oz {
	userInput := getIntUserInput("please enter the number cities:")

	newOz := &Oz{
		NumberOfCities: userInput,
		Cities:         make(map[string]*City),
	}
	return newOz
}

func CityRepr(cityName string) {

}
