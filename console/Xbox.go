package console

import "fmt"

type Xbox struct {
}

func (x *Xbox) Inicializar() {
	fmt.Println("O xbox está inicializando...")
}
