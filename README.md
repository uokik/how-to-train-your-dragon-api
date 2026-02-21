# How To Train Your Dragon API

![Go](https://img.shields.io/badge/Go-1.25-blue?logo=go&style=for-the-badge) ![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)

Simple Go library to fetch information about vikings and dragons from the "How To Train Your Dragon"

## Goals

- Add more dragons (curently 61)
- Add more vikings (curently 26)
- Add more locations (curently 10)
- Add more info for all

## Status

⚠️ Work in progress — some features are missing or not fully polished

## Functions

- Fetching Dragons by ID, Class or all
- Fetching Vikings by ID or all
- Fetching Locations by ID or all

## Installation

Make sure you have Go 1.25 or newer installed

go get github.com/szzok/how-to-train-your-dragon-api@master

## Usage

`go get github.com/uokik/how-to-train-your-dragon-api@master`

```
go
package main

import (
	"fmt"
	httydAPI "github.com/uokik/how-to-train-your-dragon-api"
)

func main() {
	dragon := httydAPI.GetDragonByID(1)
	if dragon != nil {
		fmt.Println("Found dragon:", dragon.Name)
	}
}
```

## Contribution

Feel free to open issues or submit pull requests. Any help or feedback is welcome!

## Credits

This project uses some dragons, vikings and locations descriptions sourced from

[RomainChamb](https://github.com/RomainChamb/httydApi)
and
[howtotrainyourdragon.fandom.com](https://howtotrainyourdragon.fandom.com/wiki/Category:Dragon_Species_from_the_Franchise)
