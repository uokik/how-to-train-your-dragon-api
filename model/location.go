package model

type Location struct {
	ID          int
	Name        string
	Inhabitants string
	Description string
}

var Locations = []Location{
	{ID: 1, Name: "Hidden World", Inhabitants: "Dragons", Description: "A massive, subterranean, bioluminescent utopia beneath the sea, serving as the ancestral home and sanctuary for all dragon species, hidden from humans."},
	{ID: 2, Name: "Isle of Berk", Inhabitants: "Hooligans", Description: "The original home of the Hooligan tribe, famous for its harsh climate and long, often hostile, co-existence between Vikings and wild dragons. Later became a thriving Viking/dragon society."},
	{ID: 3, Name: "New Berk", Inhabitants: "Hooligans", Description: "The secluded, final settlement founded by Hiccup and the Hooligans after the original Berk became too dangerous. It is a peaceful cliffside refuge where humans and dragons lived together before the dragons departed."},
	{ID: 4, Name: "Berserker Island", Inhabitants: "Berserker", Description: "The rugged and hostile home island of the volatile Berserker tribe, led by Dagur the Deranged. Known for its aggressive inhabitants and constant tribal conflict."},
	{ID: 5, Name: "Vanaheim", Inhabitants: "Sentinels", Description: "A mythical, jungle-covered island that serves as the sacred cemetery and final resting place for all dragons. It is permanently guarded by the elder, blind Sentinel dragons."},
	{ID: 6, Name: "Valka Mountain", Inhabitants: "Valka with dragons", Description: "A secluded ice sanctuary and massive ice cave, carved out by a massive Alpha dragon, the Bewilderbeast. It was the hidden home of Valka and hundreds of rescued dragons for two decades."},
	{ID: 7, Name: "Outcast Island", Inhabitants: "Outcast Tribe", Description: "The Outcasts are a tribe of Vikings that have been banished from various tribes, particularly the Hooligan Tribe. Its leader is Alvin the Treacherous, a former Hooligan. While originally enemies of Berk, the two tribes eventually made peace and became allies."},
	{ID: 8, Name: "Valhalla", Inhabitants: "Dead Vikings", Description: "Valhalla is the Norse location where worthy Vikings are believed to go when they die. It is mentioned numerous times in the franchise, commonly alongside a mention of dying, or just simply used for a random Viking phrase. "},
}
