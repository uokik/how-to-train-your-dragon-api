package model

type Dragon struct {
	ID          int
	Name        string
	Species     string
	Class       string
	Description string
}

var Dragons = []Dragon{
	{ID: 1, Name: "Toothless", Species: "Night Fury", Class: "Strike", Description: "Loyal and empathic with boundless, puppy-like energy"},
	{ID: 2, Name: "Stormfly", Species: "Deadly Nadder", Class: "Tracker", Description: "Precise and cunning in battle, yet warm and affectionate with friends old and new"},
	{ID: 3, Name: "Hookfang", Species: "Monstrous Nightmare", Class: "Stoker", Description: "Incendiary in temperament and short on patience, Hookfang would rather ignite 1st and think 2nd"},
	{ID: 4, Name: "Barf & Belch", Species: "Hideous Zippleback", Class: "Mystery", Description: "Truly a split personality! Barf & Belch are each fiercely independent, yet inextricably linked"},
	{ID: 5, Name: "Meatlug", Species: "Gronckle", Class: "Boulder", Description: "Exceedingly demonstrative and sweet, yet quite self-conscious despite her thick hide"},
	{ID: 6, Name: " ", Species: "Light Fury", Class: "Strike", Description: "The feline-like Light Fury is gentle and playful by nature, like Toothless, but fiercely protective when danger is present"},
	{ID: 7, Name: "Grump", Species: "Hotburple", Class: "Boulder", Description: "A loyal, lava-launching helper -- when he isn't napping"},
	{ID: 8, Name: "Cloud Jumper", Species: "Stormcutter", Class: "Sharp", Description: "Proud and confidence"},
	{ID: 9, Name: "Skull Crusher", Species: "Rumblehorn", Class: "Tracker", Description: "Doggedly determined when he's caught a scent"},
	{ID: 10, Name: " ", Species: "Armor Wing", Class: "Mystery", Description: "Metal-Plated Plunderer"},
	{ID: 11, Name: " ", Species: "Eruptodon", Class: "Boulder", Description: "Volcanic & valiant"},
	{ID: 12, Name: " ", Species: "Night terror", Class: "Stoker", Description: "Shy at first, but great as allies once trust is built"},
	{ID: 13, Name: " ", Species: "Seashocker", Class: "Tidal", Description: "Masters of the sneak attack"},
	{ID: 14, Name: " ", Species: "Skrill", Class: "Strike", Description: "Belligerent and as unpredictable as lightning strikes"},
	{ID: 15, Name: " ", Species: "Snaptrapper", Class: "Mystery", Description: "Temptation and terror to the fourth power"},
	{ID: 16, Name: " ", Species: "Teribble Terror", Class: "Stoker", Description: "Inquisitive and curious animals, their only downfall is their territorial in-fighting"},
	{ID: 17, Name: " ", Species: "Thunderdrum", Class: "Tidal", Description: "Strident and assertive, the loud Thunderdrum always makes its feelings known"},
	{ID: 18, Name: " ", Species: "Timberjack", Class: "Sharp", Description: "Modest serenity balanced with fierce protection of their riders"},
	{ID: 19, Name: " ", Species: "Whispering Death", Class: "Boulder", Description: "Blindly aggressive and capable of holding a grudge like no other dragon"},
	{ID: 20, Name: "Drago's Bewilderbeast", Species: "Bewilderbeast", Class: "Tidal", Description: "Scarred and bred for battle"},
	{ID: 21, Name: "Valka's Bewilderbeast", Species: "Bewilderbeast", Class: "Tidal", Description: "The benevolent and leonine king of all dragons"},
	{ID: 22, Name: " ", Species: "Death Song", Class: "Mystery", Description: "Enticing and siren-like, this species of dragon is decidedly solitary"},
	{ID: 23, Name: " ", Species: "Dramillion", Class: "Mystery", Description: "Imitative"},
	{ID: 24, Name: " ", Species: "", Class: "Strike", Description: ""},
	{ID: 25, Name: " ", Species: "", Class: "Strike", Description: ""},
	{ID: 26, Name: " ", Species: "", Class: "Strike", Description: ""},
	{ID: 27, Name: " ", Species: "", Class: "Strike", Description: ""},
	{ID: 28, Name: " ", Species: "", Class: "Strike", Description: ""},
}
