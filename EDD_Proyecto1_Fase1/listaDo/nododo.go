package listaDo

import (
	"fmt"

	"proyecto.com/proyecti/pilaDoble"
)

type NodoAD struct {
	sigue      *NodoAD
	anteior    *NodoAD
	pilaD      *pilaDoble.PilaDob
	Carnet     int
	Nombre     string
	Apellido   string
	Contraseña string
}

func (n *NodoAD) Mostrar() {
	fmt.Println("Carnet:", n.Carnet)
	fmt.Println("Nombre:", n.Nombre)
	fmt.Println("Apellido:", n.Apellido)
	fmt.Println("Contraseña:", n.Contraseña)
}
