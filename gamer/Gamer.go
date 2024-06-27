package gamer

import (
	"adapter/console"
	"fmt"
)

type Gamer struct {
}

func (g *Gamer) LigarConsole(c console.IConsole) {
	fmt.Println("O player ligou o console...")
	c.Inicializar()
}
