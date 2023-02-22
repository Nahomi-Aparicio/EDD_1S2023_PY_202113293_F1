package PPila

import "fmt"

type NodoPi struct {
	Estado string
	hora   string
	carnet int
	nombre string
	sig    *NodoPi
}

func (p *NodoPi) MostrarPila() {

	fmt.Println(p.Estado)
	fmt.Println(p.hora)
	fmt.Println(p.carnet)
	fmt.Println(p.nombre)

}
