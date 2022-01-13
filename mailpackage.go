package main

type Package struct {
	name   string
	weight int
}

func CreatePack() *Package {
	packname, packweight := getStrIntInput("please select a name and weight for pack:")

	newPack := &Package{
		name:   packname,
		weight: packweight,
	}
	return newPack
}
