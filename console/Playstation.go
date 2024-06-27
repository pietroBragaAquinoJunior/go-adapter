package console

import "fmt"

type Playstation struct {
}

func (p *Playstation) InicializarSomenteComAdaptacoes() {
	fmt.Println("O playstation está inicializando...")
}
