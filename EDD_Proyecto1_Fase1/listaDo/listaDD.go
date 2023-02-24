package listaDo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"proyecto.com/proyecti/PilaDoble"
)

type DoublyList struct {
	head *NodoAD
	tail *NodoAD
	size int
}

func (l *DoublyList) Añadir(carnet int, nombre string, Apellido string, contraseña string) {

	nuevonodo := &NodoAD{Carnet: carnet, Nombre: nombre, Apellido: Apellido, Contraseña: contraseña, pilaD: &PilaDoble.PilaDob{}, sigue: nil, anteior: nil}
	if l.head == nil {
		l.head = nuevonodo
		l.tail = nuevonodo
	} else {
		l.tail.sigue = nuevonodo
		nuevonodo.anteior = l.tail
		l.tail = nuevonodo
		l.size += 1

	}

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
			date := time.Now()
			t := fmt.Sprintf("%02d/%02d/%02d      %02d:%02d     ", date.Day(), date.Month(), date.Year(), date.Hour(), date.Minute())

			fmt.Println("                          ❤ ESTUDIANTE:  " + carnet + "  ❤                          ")

			fmt.Println("                      ❤ BIEVENIDO AL SISTEMA ❤                          ")
			fmt.Println("")
			aux.pilaD.AgregarP("se inicio secion", t, num)
			aux.pilaD.ImprimirPila()
			fmt.Println("")

		}
		aux = aux.sigue

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
			if actual.Carnet > actual.sigue.Carnet {

				actual.Carnet, actual.sigue.Carnet = actual.sigue.Carnet, actual.Carnet
				actual.Nombre, actual.sigue.Nombre = actual.sigue.Nombre, actual.Nombre
				actual.Apellido, actual.sigue.Apellido = actual.sigue.Apellido, actual.Apellido
				actual.Contraseña, actual.sigue.Contraseña = actual.sigue.Contraseña, actual.Contraseña

				cambiado = true
			}
			actual = actual.sigue
		}

	}
}

type Estudiante struct {
	Carne        string `json:"Carne"`
	Nombre       string `json:"Nombre"`
	Contraseña   string `json:"Contraseña"`
	Carpeta_raiz string `json:"Carpeta_raiz"`
}
type listaE struct {
	Lista []Estudiante `json:"ALUMNOS"`
}

func (P *DoublyList) CrearJon() {
	lisi := []string{}
	var str string
	aux := P.head

	for aux != nil {
		str += aux.Nombre + "," + strconv.Itoa(aux.Carnet) + "," + aux.Contraseña + "," + "/"
		lisi = append(lisi, str)
		aux = aux.sigue
		str = ""
	}
	lista1 := listaE{}
	for _, v := range lisi {
		campo := strings.Split(v, ",")
		estudiante := Estudiante{
			Carne:        campo[1],
			Nombre:       campo[0],
			Contraseña:   campo[2],
			Carpeta_raiz: campo[3],
		}
		lista1.Lista = append(lista1.Lista, estudiante)
	}

	jsonData, err := json.Marshal(lista1)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("lista_estudiantes.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
		return
	}
}

func (L *DoublyList) Graficar() {
	L.size += 1
	var graphipila1 string

	temp := L.head

	//e := 1
	graphipila1 = "digraph G {\n"
	graphipila1 += "rankdir=LR;\n"
	graphipila1 += "node [shape=box];\n"

	graphipila1 += "label = \"BITACORA DE ESTUDIANTES\";\n"
	for temp != nil && temp.sigue != nil {
		graphipila1 += strconv.Itoa(temp.Carnet) + "[label=\"{" + temp.Nombre + " " + temp.Apellido + " " + strconv.Itoa(temp.Carnet) + "}\"];\n"
		graphipila1 += strconv.Itoa(temp.Carnet) + " -> " + strconv.Itoa(temp.sigue.Carnet) + ";\n"
		graphipila1 += strconv.Itoa(temp.sigue.Carnet) + " -> " + strconv.Itoa(temp.Carnet) + ";\n" //N1 -> N2;

		z := temp.pilaD.GraficarPila(temp.Carnet)
		graphipila1 += z

		temp = temp.sigue
	}
	for temp != nil && temp.sigue == nil {
		graphipila1 += strconv.Itoa(temp.Carnet) + "[label=\"{" + temp.Nombre + " " + temp.Apellido + " " + strconv.Itoa(temp.Carnet) + "}\"];\n"
		temp = temp.sigue
	}
	//fmt.Println(graphipila1)
	graphipila1 += "}"

	file, or := os.Create("lista.dot")
	if or != nil {
		fmt.Println("Error al crear el archivo")
		return
	}
	file.WriteString(graphipila1)
	file.Close()
	//creamos la imagen
	cmd := exec.Command("dot", "-Tpng", "lista.dot", "-o", "lista.png")
	arr := cmd.Run()
	if arr != nil {
		fmt.Println("Error al crear la imagen")
		return
	}

}
