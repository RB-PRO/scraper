package scraperfogplay

import (
	fogplay "github.com/RB-PRO/PhotoTemaParser/pkg/fogplaymts"
)

func Parsing() {
	games := fogplay.Servers()
	fogplay.Save(games)
}
