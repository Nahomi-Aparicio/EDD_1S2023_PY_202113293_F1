package colita

import "proyecto.com/proyecti/estructura"

type NODOcola struct {
	Estudiante *estructura.Estudiante
	Siguiente  *NODOcola
}

func (n *NODOcola) Mostrar() {
	n.Estudiante.Mostrar()
}
