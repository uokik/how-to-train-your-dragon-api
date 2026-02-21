package main

import (
	"fmt"
	"log"

	httydAPI "github.com/szzok/how-to-train-your-dragon-api"
)

// all

func alldragons() {
	dragons := httydAPI.GetAllDragons()

	fmt.Println("Dragons list:")
	for _, d := range dragons {
		fmt.Printf("%d %s %s \nDescription: %s\n", d.ID, d.Name, d.Species, d.Description)
	}
}

func allvikings() {
	vikings := httydAPI.GetAllVikings()

	fmt.Println("Vikings list:")
	for _, v := range vikings {
		fmt.Printf("%d, %s %s %s %s \n", v.ID, v.FirstName, v.LastName, v.Gender, v.Location)
	}
}

func alllocations() {
	locations := httydAPI.GetAllLocations()

	fmt.Println("Dragons list:")
	for _, l := range locations {
		fmt.Printf("%d %s %s \nDescription: %s\n", l.ID, l.Name, l.Inhabitants, l.Description)
	}
}

// by id

func dragonbyid() {
	dragon := httydAPI.GetDragonByID(4)

	if dragon != nil {
		fmt.Printf("Dragon id %d: %s %s | Class: %s \nDescription: %s\n", dragon.ID, dragon.Name, dragon.Species, dragon.Class, dragon.Description)
	} else {
		log.Println("No dragon with ID 4 found")
	}
}

func vikingbyid() {
	viking := httydAPI.GetVikingByID(6)

	if viking != nil {
		fmt.Printf("Viking id %d: %s %s %s %s\n", viking.ID, viking.FirstName, viking.LastName, viking.Gender, viking.Location)
	} else {
		log.Println("No viking with ID 6 found")
	}
}

func locationbyid() {
	location := httydAPI.GetLocationByID(3)

	fmt.Println("Location list:")
	if location != nil {
		fmt.Printf("%d: %s %s %s \n", location.ID, location.Name, location.Inhabitants, location.Description)
	} else {
		log.Println("No location with ID 3 found")
	}
}

// by class

func getdragonbyclass() {
	dragons := httydAPI.GetDragonsByClass("Mystery")

	fmt.Println("Dragons list:")
	for _, d := range dragons {
		fmt.Printf("%d %s %s \nDescription: %s\n", d.ID, d.Name, d.Species, d.Description)
	}
}

func main() {
	alldragons()
	allvikings()
	alllocations()
	getdragonbyclass()
	vikingbyid()
	dragonbyid()
	locationbyid()
}
