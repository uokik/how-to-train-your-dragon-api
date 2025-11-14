package service

import "github.com/szzok/how-to-train-your-dragon-api/model"

func GetAllVikings() []model.Viking {
	return model.Vikings
}

func GetVikingByID(id int) *model.Viking {
	for _, v := range model.Vikings {
		if v.ID == id {
			return &v
		}
	}
	return nil
}
