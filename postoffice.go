package main

type PostOffice struct {
	numOfPack   int
	max_weight  int
	min_weight  int
	mailPackage map[string]*int
}

func CreatePostOffice(hostCity *City) *PostOffice {
	packNum, max, min := getSplittedInput("please enter number of packs, max weight and min weight")

	newPost := &PostOffice{
		numOfPack:   packNum,
		max_weight:  max,
		min_weight:  min,
		mailPackage: make(map[string]*int),
	}

	hostCity.TotalNumOfPacks += packNum
	return newPost
}
