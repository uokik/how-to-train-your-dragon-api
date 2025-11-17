package main

import (
	"fmt"

	httydAPI "github.com/szzok/how-to-train-your-dragon-api"
)

func main() {
	dragons := httydAPI.GetDragonsByClass("Mystery")

	fmt.Println("Dragons list:")
	for _, d := range dragons {
		fmt.Printf("%d %s %s \nDescription: %s\n", d.ID, d.Name, d.Species, d.Description)
	}
}
