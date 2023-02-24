package PilaDoble

import "fmt"

type NodoPiDo struct {
	secion string
	hora   string
	carne  int
	sig    *NodoPiDo
}

func (dd *NodoPiDo) MostarPilaDo() {

	fmt.Println(dd.secion)
	fmt.Println(dd.hora)
	fmt.Println(dd.carne)

}

func (dd *NodoPiDo) MostarPilaDo2() string {
	return dd.secion
}
