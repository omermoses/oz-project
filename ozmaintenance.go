package main

func UpdateMaxPackNumCity(ozCity *Oz, cityA *City, cityB *City) {
	// Checking If we should update the city with max number of packages

	if cityB == nil {
		ozCity.MaxNumPackCity = cityA.name
		return
	}
	if cityA.TotalNumOfPacks < cityB.TotalNumOfPacks {
		ozCity.MaxNumPackCity = cityB.name
		return
	}

}
