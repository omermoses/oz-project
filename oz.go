package main

type Oz struct {
	NumberOfCities int
	MaxNumPackCity string
	Cities         map[string]*City
}

func CreateOz() *Oz {
	userInput := getIntUserInput("please enter the number cities:")

	newOz := &Oz{
		NumberOfCities: userInput,
		MaxNumPackCity: "",
		Cities:         make(map[string]*City),
	}
	return newOz
}
