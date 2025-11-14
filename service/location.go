package service

import "github.com/szzok/how-to-train-your-dragon-api/model"

func GetAllLocations() []model.Location {
	return model.Locations
}

func GetLocationByID(id int) *model.Location {
	for _, l := range model.Locations {
		if l.ID == id {
			return &l
		}
	}
	return nil
}
