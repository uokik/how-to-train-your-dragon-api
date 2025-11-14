package model

type Viking struct {
	ID        int
	FirstName string
	LastName  string
	Gender    string
	Location  string
}

var Vikings = []Viking{
	{ID: 1, FirstName: "Hiccup Horrendous", LastName: "Haddock III", Gender: "Male", Location: "Berk"},
	{ID: 2, FirstName: "Astrid", LastName: "Hofferson", Gender: "Female", Location: "Berk"},
	{ID: 3, FirstName: "Fishlegs", LastName: "Ingerman", Gender: "Male", Location: "Berk"},
	{ID: 4, FirstName: "Ruffnut", LastName: "Thorston", Gender: "Female", Location: "Berk"},
	{ID: 5, FirstName: "Tuffnut", LastName: "Thorston", Gender: "Male", Location: "Berk"},
	{ID: 6, FirstName: "Snotloud", LastName: "Jorgenson", Gender: "Male", Location: "Berk"},
	{ID: 7, FirstName: "Stoick The Vast", LastName: "Haddock", Gender: "Male", Location: "Berk"},
	{ID: 8, FirstName: "Valka", LastName: "Haddock", Gender: "Female", Location: "Berk"},
	{ID: 9, FirstName: "Gobber The Belch", LastName: " ", Gender: "Male", Location: "Berk"},
	{ID: 10, FirstName: "Eret Son of Eret", LastName: " ", Gender: "Male", Location: "Unknown"},
	{ID: 11, FirstName: "Dagur", LastName: " ", Gender: "Male", Location: "Berserker Island"},
	{ID: 12, FirstName: "Ryker", LastName: "Grimborn", Gender: "Male", Location: "Unknown"},
	{ID: 13, FirstName: "Drago-Bludvist", LastName: " ", Gender: "Male", Location: "Unknown"},
	{ID: 14, FirstName: "Heather", LastName: " ", Gender: "Female", Location: "Unknown"},
}
