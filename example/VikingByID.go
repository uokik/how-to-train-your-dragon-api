package example4

import (
	"fmt"
	"log"

	httydAPI "github.com/szzok/how-to-train-your-dragon-api"
)

func example4() {
	viking := httydAPI.GetVikingByID(6)

	if viking != nil {
		fmt.Printf("Viking id %d: %s %s %s %s\n", viking.ID, viking.FirstName, viking.LastName, viking.Gender, viking.Location)
	} else {
		log.Println("No viking with ID 6 found")
	}
}
