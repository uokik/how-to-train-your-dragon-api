package httydAPI

import "github.com/szzok/how-to-train-your-dragon-api/service"

var (
	GetAllVikings     = service.GetAllVikings
	GetVikingByID     = service.GetVikingByID
	GetAllDragons     = service.GetAllDragons
	GetDragonByID     = service.GetDragonByID
	GetAllLocations   = service.GetAllLocations
	GetLocationByID   = service.GetLocationByID
	GetDragonsByClass = service.GetDragonsByClass
)
