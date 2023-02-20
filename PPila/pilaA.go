package PPila

type PilaA struct {
	head  *NodoPi
	size2 int
}

func (p *PilaA) Push(estado string, hora string) {

	nuevono := &NodoPi{estado, hora, nil}
	if p.head == nil {
		p.head = nuevono
	} else {
		//insertar al inicio
		temp := p.head
		p.head = nuevono
		p.head.sig = temp
		p.size2 += 1
	}

}

func (p *PilaA) Print() {
	temp := p.head
	for temp != nil {
		temp.MostrarPila()
		temp = temp.sig

	}
}
