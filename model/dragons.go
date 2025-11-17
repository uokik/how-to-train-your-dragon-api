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
	{ID: 24, Name: " ", Species: "Red Death", Class: "Stoker", Description: "The Red Death was a Titan Wing dragon secretly responsible for centuries of human-dragon conflict, using its commanding presence to force other dragons to raid villages and bring it food. It reigned supreme from its volcanic lair on Dragon Island until the Berk Vikings and Toothless found it. The enormous dragon had to shatter its mountain home to emerge. Ultimately, Hiccup and Toothless discovered and exploited its only vulnerable spot—its insides—leading to its defeat"},
	{ID: 25, Name: " ", Species: "Boneknapper", Class: "Mystery", Description: "This Mystery Class scavenger covers itself in bones to shield its delicate body. Its choice of armor makes it unpopular with other Dragons! The Boneknapper is well-protected in its armor while it gets close and delivers a potent stream of fire to its fleeing opponents"},
	{ID: 26, Name: " ", Species: "Buffalord", Class: "Mystery", Description: "Don't let its docile demeanor fool you! Thought to be extinct, this already massive dragon will double in size and will become extremely aggressive if removed from its food source. When in battle (or away from its food) this dragon can puff up to reveal spikes covering it from head to tail. If it's not blasting these spikes, it's snorting out flames"},
	{ID: 27, Name: " ", Species: "Catastrophic Quaken", Class: "Boulder", Description: "Catastrophic Quakens are known for their cruel attitudes. They don't let anyone get in their way, and they'll go out of their way to wreak havoc on their target. Catastrophic Quakens eat more rocks than any other in their class! They also have a nasty habit of spitting up digested rocks at attackers in the form of hot lava"},
	{ID: 28, Name: " ", Species: "Cavern Crasher", Class: "Mystery", Description: "The Cavern Crasher flings flammable mucus and scurries across underground catacombs in search of its favorite food: Firecomb"},
	{ID: 29, Name: " ", Species: "Changewing", Class: "Mystery", Description: " Unlike other dragons the Changewing spits acid instead of fire. Even newly-hatched Changewings have the ability to spit acid onto any potential threat. They got their name from the changing color of their scales which they use as camouflage"},
	{ID: 30, Name: " ", Species: "Crimson Goregutter", Class: "Boulder", Description: "Although Goregutters have impressive sets of antlers and axe-shaped tails, they are gentler than they appear and like to be left in peace. In battle, however, they have the ability to spew molten lava onto their antlers before ramming opponents with fiery force"},
	{ID: 31, Name: " ", Species: "Deathgripper", Class: "Strike", Description: "Deathgrippers are a one-of-a-kind dragon, born with venomous pincer tails which are eerily to those of a venomous scorpion. These dragons love spending time up in the air, helping Vikings reinforce dragon riding skills as they go"},
	{ID: 32, Name: " ", Species: "Fireworm", Class: "Stoker", Description: "Each Fireworm colony holds only one female, the Queen! Although rare, when a new female is born, she sets out to establish her own colony"},
	{ID: 33, Name: " ", Species: "Flightmare", Class: "Mystery", Description: "Instead of shooting fire, this Dragon from the Mystery Class emits a toxic mist that paralyzes its foes"},
	{ID: 34, Name: " ", Species: "Hobgobbler", Class: "Mystery", Description: "Believed by many to bring bad luck, Hobgobblers had a nasty reputation due to superstitious assumptions"},
	{ID: 35, Name: " ", Species: "Razorwhip", Class: "Sharp", Description: "Try to get on this Dragon's good side... if you can find one. Their toxic breath and razor-sharp tail make them dangerous from top to bottom"},
	{ID: 36, Name: " ", Species: "Sand Wraith", Class: "Tidal", Description: " A Tidal Class expert in camouflage that buries itself in the sand. It's a good way to surprise their prey"},
	{ID: 37, Name: " ", Species: "Scauldron", Class: "Tidal", Description: "This huge underwater dragon is a speedy and silent swimmer. Instead of breathing fire, Scauldron spits scalding hot water at its prey"},
	{ID: 38, Name: " ", Species: "Screaming Death", Class: "Boulder", Description: " Born every one hundred or so years, the Screaming Death has all the strengths of its subordinate cousin, the Whispering Death, with none of its weaknesses. It is the most powerful dragon Berk has faced since the Red Death"},
}
