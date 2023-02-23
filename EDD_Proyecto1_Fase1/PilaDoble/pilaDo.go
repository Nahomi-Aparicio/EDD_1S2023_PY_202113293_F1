package PilaDoble

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

func (p *PilaDob) graficar(num int, grap string) {
	aux := p.cabeza
	for aux != nil {
		if aux.carne == num {
			grap += aux.hora

		}
		aux = aux.sig
	}
}

func (p *PilaDob) MostrarPilaDo() {
	aux := p.cabeza
	for aux != nil {
		aux.MostarPilaDo()
		aux = aux.sig
	}
}
