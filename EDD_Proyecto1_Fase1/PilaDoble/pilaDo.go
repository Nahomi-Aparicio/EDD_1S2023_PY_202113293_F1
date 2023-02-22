package pilaDoble

type PilaDob struct {
	cabeza *NodoPiDo
	sizep  int
}

func (p *PilaDob) AgregarP(secion string, hora string) {

	nuevono := &NodoPiDo{secion, hora, nil}
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
