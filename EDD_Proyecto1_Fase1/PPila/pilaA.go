package PPila

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type PilaA struct {
	head  *NodoPi
	size2 int
}

func (p *PilaA) Push(estado string, hora string, carnet int, nombre string) {

	nuevono := &NodoPi{estado, hora, carnet, nombre, nil}
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

func (p *PilaA) Graph() {

	p.size2 += 1
	var graphipila1 string

	temp := p.head

	e := 1
	graphipila1 = "digraph G {\n"
	graphipila1 += "rankdir=BT;\n"
	graphipila1 += "node [shape=box];\n"
	graphipila1 += "node [shape=record fontname=Arial]\n"
	graphipila1 += "label = \"PILA DE ACCIONES DEL ADMIN\";\n"
	for temp != nil && temp.sig != nil {

		graphipila1 += "N" + strconv.Itoa(temp.carnet) + "[label=\"{" + temp.Estado + "|" + temp.hora + "|" + temp.nombre + strconv.Itoa(temp.carnet) + "}\"];\n"
		graphipila1 += "N" + strconv.Itoa(temp.sig.carnet) + "->N" + strconv.Itoa(temp.carnet) + ";\n"
		e += 1

		temp = temp.sig
		p.size2 -= 1
	}

	graphipila1 += "}"
	//fmt.Println(graphipila1)

	file, or := os.Create("pila.dot")
	if or != nil {
		fmt.Println("Error al crear el archivo")
		return
	}
	file.WriteString(graphipila1)
	file.Close()
	//creamos la imagen
	cmd := exec.Command("dot", "-Tpng", "pila.dot", "-o", "pila.png")
	arr := cmd.Run()
	if arr != nil {
		fmt.Println("Error al crear la imagen")
		return
	}

}
func (p *PilaA) VerificarCar(v int) bool {
	temp := p.head
	for temp != nil {
		if temp.carnet == v {
			fmt.Println("EL CARNE DEL ESTUDIANTE YA EXISTE")
			return true

		}
		temp = temp.sig
	}
	return false
}
