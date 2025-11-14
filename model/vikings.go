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
}
