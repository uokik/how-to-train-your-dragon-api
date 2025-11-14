package main

import (
	"fmt"

	httydAPI "github.com/szzok/how-to-train-your-dragon-api"
)

func main() {
	vikings := httydAPI.GetAllVikings()

	fmt.Println("Vikings list:")
	for _, v := range vikings {
		fmt.Printf("%d, %s %s %s %s \n", v.ID, v.FirstName, v.LastName, v.Gender, v.Location)
	}
}
