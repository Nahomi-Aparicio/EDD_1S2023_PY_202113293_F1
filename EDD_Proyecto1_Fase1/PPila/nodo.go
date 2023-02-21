package PPila

import "fmt"

type NodoPi struct {
	Estado string
	hora   string
	sig    *NodoPi
}

func (p *NodoPi) MostrarPila() {

	fmt.Println(p.Estado)
	fmt.Println(p.hora)

}
