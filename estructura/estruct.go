package estructura

import (
	"fmt"
)

type Estudiante struct {
	Carnet     string
	Nombre     string
	Apellido   string
	Contraseña string
}

func (e *Estudiante) Mostrar() {

	fmt.Println("Carnet:", e.Carnet)
	fmt.Println("Nombre:", e.Nombre)
	fmt.Println("Apellido:", e.Apellido)
	fmt.Println("Contraseña:", e.Contraseña)

}
