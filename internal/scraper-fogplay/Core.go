package scraperfogplay

import (
	"fmt"

	fogplay "github.com/RB-PRO/PhotoTemaParser/pkg/fogplaymts"
)

func Parsing() {
	ss := fogplay.Serverss()
	GGG := make([]fogplay.Game, 0, len(ss))
	for _, s := range ss {
		GGG = append(GGG, fogplay.Gamess(s))
	}
	fmt.Println(len(GGG))
	fogplay.Save(GGG)
}
