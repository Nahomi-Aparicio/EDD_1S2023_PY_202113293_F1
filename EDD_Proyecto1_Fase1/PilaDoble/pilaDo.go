package PilaDoble

import (
	"strconv"
)

type PilaDob struct {
	cabeza *NodoPiDo
	sizep  int
}

func (p *PilaDob) AgregarP(secion string, hora string, carne int) {

	nuevono := &NodoPiDo{secion, hora, carne, nil}
	//fmt.Println(secion, hora)
	if p.cabeza == nil {
		p.cabeza = nuevono
	} else {
		//insertar al inicio
		temp := p.cabeza
		p.cabeza = nuevono
		p.cabeza.sig = temp
		p.sizep += 1

	}

}

func (p *PilaDob) GraficarPila(num int) string {
	var grap string
	aux := p.cabeza
	a := 1
	//e := 1
	a2 := 0
	for aux != nil {

		if aux.carne == num {

			grap += "{ rank=same \n " //si le pongo el rank no jala el primero
			grap += strconv.Itoa(aux.carne+a) + "[label=\"{" + aux.hora + " " + aux.secion + "}\"];\n"
			grap += strconv.Itoa(aux.carne+a2) + " -> " + strconv.Itoa(aux.carne+a) + ";\n"
			grap += "}"
		}

		a += 1
		a2 += 1
		aux = aux.sig

	}

	return grap
}

func (P *PilaDob) ImprimirPila() {
	aux := P.cabeza
	for aux != nil {
		aux.MostarPilaDo()
		aux = aux.sig
	}
}
