package model

type Dragon struct {
	ID          int
	Name        string
	Species     string
	Class       string
	Description string
}

var Dragons = []Dragon{
	{ID: 1, Name: "Toothless", Species: "Night Fury", Class: "Strike", Description: "The rarest and most intelligent of the dragon species, the Night Fury is distinguished by its dark color and piercing yellow eyes. Toothless is one of a kind; the only Night Fury ever seen in Berk."},
	{ID: 2, Name: "Stormfly", Species: "Deadly Nadder", Class: "Tracker", Description: "Deadly Nadders can suddenly raise the hundreds of sharp spines that stud their hides and tails and fling them with incredible accuracy. But when the shooting starts, there is no safer place to be than face-to-face with a Nadder - If you stand right in front of its nose, a Nadder won't be able to see you"},
	{ID: 3, Name: "Hookfang", Species: "Monstrous Nightmare", Class: "Stoker", Description: "A stubborn and tenacious Stoker Class Dragon. Look out for its Fire Jacket: It covers itself in flames when attacking! The Monstrous Nightmare attacks its foes with powerful streams of fire."},
	{ID: 4, Name: "Next Dragon", Species: "Night Fury", Class: "Strike", Description: "Description"},
}
