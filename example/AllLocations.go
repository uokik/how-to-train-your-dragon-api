package main

import (
	"fmt"

	httydAPI "github.com/szzok/how-to-train-your-dragon-api"
)

func main() {
	locations := httydAPI.GetAllLocations()

	fmt.Println("Dragons list:")
	for _, l := range locations {
		fmt.Printf("%d %s %s \nDescription: %s\n", l.ID, l.Name, l.Inhabitants, l.Description)
	}
}
