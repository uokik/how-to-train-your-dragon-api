package service

import "github.com/szzok/how-to-train-your-dragon-api/model"

func GetAllDragons() []model.Dragon {
	return model.Dragons
}

func GetDragonByID(id int) *model.Dragon {
	for _, d := range model.Dragons {
		if d.ID == id {
			return &d
		}
	}
	return nil
}
