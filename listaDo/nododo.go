package listaDo

import "proyecto.com/proyecti/estructura"

type Nodo struct {
	Next *Nodo
	Prev *Nodo

	Estudiante *estructura.Estudiante
}
