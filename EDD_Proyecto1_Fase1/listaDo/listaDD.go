package listaDo

import (
	"fmt"
	"sort"
	"strconv"
)

type DoublyList struct {
	head *NodoAD
	tail *NodoAD
}

func (l *DoublyList) Añadir(carnet int, nombre string, Apellido string, contraseña string) {

	nuevonodo := &NodoAD{Carnet: carnet, Nombre: nombre, Apellido: Apellido, Contraseña: contraseña, pilaD: nil, sigue: nil, anteior: nil}
	if l.head == nil {
		l.head = nuevonodo
		l.tail = nuevonodo
	} else {
		l.tail.sigue = nuevonodo
		nuevonodo.anteior = l.tail
		l.tail = nuevonodo
	}
	fmt.Println("Se agrego el estudiante a la lista")
}

func (l *DoublyList) Imprimir() {
	aux := l.head
	for aux != nil {
		aux.Mostrar()
		fmt.Println(" ")
		aux = aux.sigue
	}
}

func (l *DoublyList) BuscaryagregarHora(carnet string, hora string) {
	num, err := strconv.Atoi(carnet)
	if err != nil {
		fmt.Println("")

	}
	aux := l.head
	for aux != nil {
		if aux.Carnet == num {
			aux.pilaD.AgregarP(carnet, hora)

			fmt.Println("                      ❤ BIEVENIDO AL SISTEMA ❤                          ")
			fmt.Println("")

		}
		aux = aux.sigue

	}
}

/*func (l *DoublyList) Ordenar() {
	aux := l.head
	for aux != nil {
		for aux.sigue != nil {
			if aux.Estudiante.Carnet < aux.sigue.Estudiante.Carnet {
				fmt.Println("este es el primer carnet: ", aux.Estudiante.Carnet)
				fmt.Println("este es el segundo carnet: ", aux.sigue.Estudiante.Carnet)
				aux.Estudiante, aux.sigue.Estudiante = aux.sigue.Estudiante, aux.Estudiante
			}

			aux = aux.sigue
		}
	}
}*/

func (l *DoublyList) Gurdarcarnet() {
	lista := []interface{}{}
	lista2 := [][]interface{}{}
	//lista := make([][][]string, 0)
	aux := l.head
	for aux != nil {

		lista = append(lista, aux.Carnet)
		lista = append(lista, aux.Nombre)
		lista = append(lista, aux.Apellido)

		lista2 = append(lista2, lista)

		aux = aux.sigue
		lista = []interface{}{}
	}

	sort.Slice(lista2, func(i, j int) bool {
		return lista2[i][0].(int) < lista2[j][0].(int)
	})

	// Imprimir la lista resultante
	//fmt.Println(lista2)

	for _, fila := range lista2 {
		// Imprimir cada elemento de la fila separado por un espacio
		fmt.Println("═════════════════════════════════════════")
		for i, elemento := range fila {

			if i == 0 {

				fmt.Print(elemento)

			} else {

				fmt.Print(" " + elemento.(string))

			}

		}

		fmt.Println()
	}

}

func (regEst *DoublyList) OrdenarPorCarnet() {
	if regEst.head == nil {
		return
	}
	cambiado := true
	for cambiado {
		cambiado = false
		actual := regEst.head
		for actual.sigue != nil {
			if actual.Carnet < actual.sigue.Carnet {
				fmt.Println("este es el primer carnet: ", actual.Carnet)
				fmt.Println("este es el segundo carnet: ", actual.sigue.Carnet)
				fmt.Println("═════════════════════════════════════════")
				actual.Carnet, actual.sigue.Carnet = actual.sigue.Carnet, actual.Carnet
				fmt.Println("este es el primer carsnet: ", actual.Carnet)
				fmt.Println("este es el segundo casrnet: ", actual.sigue.Carnet)
				fmt.Println("═════════════════════════════════════════")
				cambiado = true
			}
			actual = actual.sigue
		}

	}
}
