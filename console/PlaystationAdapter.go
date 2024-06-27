package console

import "fmt"

type PlaystationAdapter struct {
	Playstation *Playstation
}

func (pa *PlaystationAdapter) Inicializar() {
	fmt.Println("O adaptador está adaptando para o playstation...")
	pa.Playstation.InicializarSomenteComAdaptacoes()
}
