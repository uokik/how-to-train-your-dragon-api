package service

import "github.com/uokik/how-to-train-your-dragon-api/model"

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

func GetDragonsByClass(class string) []model.Dragon {
	var dragons []model.Dragon
	for _, d := range model.Dragons {
		if d.Class == class {
			dragons = append(dragons, d)
		}
	}
	return dragons
}
