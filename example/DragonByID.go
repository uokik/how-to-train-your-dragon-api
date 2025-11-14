package main

import (
	"fmt"
	"log"

	httydAPI "github.com/szzok/how-to-train-your-dragon-api"
)

func example() {
	dragon := httydAPI.GetDragonByID(4)

	if dragon != nil {
		fmt.Printf("Dragon id %d: %s %s | Class: %s \nDescription: %s\n", dragon.ID, dragon.Name, dragon.Species, dragon.Class, dragon.Description)
	} else {
		log.Println("No dragon with ID 4 found")
	}
}
