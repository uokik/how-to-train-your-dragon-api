package main

import (
	"fmt"
	"log"

	httydAPI "github.com/szzok/how-to-train-your-dragon-api"
)

func main() {
	location := httydAPI.GetLocationByID(3)

	fmt.Println("Location list:")
	if location != nil {
		fmt.Printf("%d: %s %s %s \n", location.ID, location.Name, location.Inhabitants, location.Description)
	} else {
		log.Println("No location with ID 3 found")
	}
}
