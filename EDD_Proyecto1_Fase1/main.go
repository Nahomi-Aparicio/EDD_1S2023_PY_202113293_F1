package main

import (
	"fmt"

	"proyecto.com/proyecti/PPila"
	"proyecto.com/proyecti/administrador"
	"proyecto.com/proyecti/colita"
	"proyecto.com/proyecti/listaDo"
)

func menusecion(espera *colita.Cola, pila *PPila.PilaA, listado *listaDo.DoublyList) {
	//ahora := time.Now()❤

	var usuario string
	var contrasena string
	fmt.Println("❤ INGRESE USUARIO:")
	fmt.Scanln(&usuario)

	fmt.Println("❤ INGRESE CONTRASEÑA:")
	fmt.Scanln(&contrasena)

	listado.BuscaryagregarHora(usuario, contrasena)
	//

	if usuario == "admin" && contrasena == "admin" {

		fmt.Println("")

		fmt.Println("                      ❤ BIEVENIDO AL SISTEMA ❤                          ")
		fmt.Println("")

		administrador.Menuadmi(espera, pila, listado)

		//menuadmi()
		//fmt.Printf("Fecha y hora de ingreso: %s\n", ahora.Format("2006-01-02 15:04:05"))  ❤

	} else {

		fmt.Println("")

	}

}

func Menu1() {
	var opcion int
	espera := &colita.Cola{}
	pila := &PPila.PilaA{}
	listado := &listaDo.DoublyList{}
	opcion = 0
	for opcion != 2 {

		fmt.Println("═══════════════════════ EDD GODRIVER ═══════════════════════ ")
		fmt.Println("❤                   1. INICIAR SESION                 ❤")
		fmt.Println("❤                   2. SALIR DEL SISTEMA              ❤")
		fmt.Println("═════════════════════════════════════════════════════════════")
		fmt.Println("Ingrese una opcion:   ")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			fmt.Println("Iniciando sesion")
			fmt.Println("")
			fmt.Println("")
			menusecion(espera, pila, listado)

			// falta aqui para iniciar el sistema   go run main.go

		case 2:
			fmt.Println("Saliendo del sistema")

		}

	}

}

func main() {
	Menu1()

}
