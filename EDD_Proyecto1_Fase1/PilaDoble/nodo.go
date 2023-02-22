package pilaDoble

import "fmt"

type NodoPiDo struct {
	secion string
	hora   string
	sig    *NodoPiDo
}

func (dd *NodoPiDo) MostarPilaDo() {

	fmt.Println(dd.secion)
	fmt.Println(dd.hora)

}
