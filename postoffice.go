package main

type PostOffice struct {
	numOfPack   int
	max_weight  int
	min_weight  int
	mailPackage []*Package
}

func CreatePostOffice(hostCity *City) *PostOffice {
	packNum, max, min := getSplittedInput("please enter number of packs, max weight and min weight")

	newPost := &PostOffice{
		numOfPack:   packNum,
		max_weight:  max,
		min_weight:  min,
		mailPackage: make([]*Package, packNum),
	}

	hostCity.TotalNumOfPacks += packNum
	return newPost
}
