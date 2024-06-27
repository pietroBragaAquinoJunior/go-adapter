package main

import (
	"adapter/console"
	"adapter/gamer"
)

func main() {
	g := gamer.Gamer{}
	p := console.Playstation{}
	pa := console.PlaystationAdapter{Playstation: &p}
	x := console.Xbox{}
	g.LigarConsole(&pa)
	g.LigarConsole(&x)
}
