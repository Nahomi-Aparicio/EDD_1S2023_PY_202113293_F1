package administrador

import (
	//"encoding/csv"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	//"os"

	"proyecto.com/proyecti/PPila"
	"proyecto.com/proyecti/colita"
	"proyecto.com/proyecti/listaDo"
)

func Menuadmi(colita *colita.Cola, pila *PPila.PilaA, lis *listaDo.DoublyList) {

	var opcion int
	validar := 0
	var nom, cont, Ap string
	var car int
	for opcion != 6 {

		fmt.Println("═══════════════════════ ADMINISTRADOR - EDD GODRIVE ══════════════════════ ")
		fmt.Println("❤                   1. VER ESTUDIANTES PENDIENTES                        ❤")
		fmt.Println("❤                   2. VER ESTUDIANTES DEL SISTEMA                       ❤")
		fmt.Println("❤                   3. REGISTRAR NUEVOS ESTUDIANTES                      ❤")
		fmt.Println("❤                   4. CARGA MASIVA DE ESTUDIANTES                       ❤")
		fmt.Println("❤                   5. VER REPORTE DE ESTUDIANTES                        ❤")
		fmt.Println("❤                   6. CERRAR SESION                                     ❤")
		fmt.Println("═══════════════════════════════════════════════════════════════════════════ ")

		fmt.Println("Ingrese una opcion:   ")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			//colita.Graph1()

			validar = 0

			for validar != 3 {

				if !colita.Vacia() {

					temp, size := colita.Eliminar()
					colita.Graph1()
					fmt.Println("═══════════════════════ ESTUDIANTES PENDIENTES ═════════", size+1, "═════════")

					fmt.Printf("Carnet: %2d\n", temp.Carnet)
					fmt.Printf("Nombre: %s\n", temp.Nombre)
					fmt.Println(" ═══════════════════════ ELIJA UNA OPCION PARA EL ESTUDIANTE ══════════════════════ ")
					fmt.Println("❤                   1. ACEPTAR                                                    ❤")
					fmt.Println("❤                   2. RECHAZAR                                                   ❤")
					fmt.Println("❤                   3. REGRESAR                                                   ❤")
					fmt.Println(" ═══════════════════════════════════════════════════════════════════════════════════")
					fmt.Print("Elige una opcion: ")

					fmt.Scanln(&validar)
					switch validar {
					case 1:
						date := time.Now()
						t := fmt.Sprintf("%02d/%02d/%02d      %02d:%02d     ", date.Day(), date.Month(), date.Year(), date.Hour(), date.Minute())

						fmt.Println("SE ACEPTO AL ESTUDIANTE")
						fmt.Println("")
						pila.Push(t, "se acepto a el estudiante", temp.Carnet, temp.Nombre)
						lis.Añadir(temp.Carnet, temp.Nombre, temp.Apellido, temp.Contraseña)

					case 2:
						date := time.Now()
						t := fmt.Sprintf("%02d/%02d/%02d      %02d:%02d     ", date.Day(), date.Month(), date.Year(), date.Hour(), date.Minute())

						fmt.Println("SE RECHAZO AL ESTUDIANTE")
						fmt.Println("")
						pila.Push(t, "se rechazo a el estudiante", temp.Carnet, temp.Nombre)
					case 3:
						fmt.Println("Regresando al menu principal")
						fmt.Println("")
					default:
						fmt.Println("Opcion no valida")
						fmt.Println("")
					}

				} else {
					validar = 3
					fmt.Println("═══════════════════════ ESTUDIANTES PENDIENTES ══════════════════")
					fmt.Println("")
					fmt.Println("❤                 NO HAY ESTUDIANTES PENDIENTES               ❤")
					fmt.Println("")
				}
				//pila.Print()

				//lis.BuscaryagregarHora("201780044", t)

				//

			}

		case 2:
			fmt.Println("═══════════════════════ ADMINISTRADOR - ESTUDIANTES EN EL SISTEMA ══════════════════════ ")
			//lis.Imprimir()
			lis.OrdenarPorCarnet()
			//lis.Gurdarcarnet()
			lis.Imprimir()
		case 3:

			fmt.Println("═══════════════════════ ADMINISTRADOR - registro de estudiantes ══════════════════════ ")

			fmt.Println("❤ INGRESE NOMBRE: ")
			fmt.Scanln(&nom)
			fmt.Println("❤ INGRESE APELLIDO: ")
			fmt.Scanln(&Ap)
			fmt.Println("❤ INGRESE CARNE: ")
			fmt.Scanln(&car)
			fmt.Println("❤ INGRESE CONTRASEÑA	: ")
			fmt.Scanln(&cont)
			colita.Agregar(car, nom, Ap, cont)

			fmt.Println("")
			fmt.Println("                              ❤ ESTUDIANTE REGISTRADO❤                               ")
			fmt.Println("")

		case 4:

			Cargamasiva(colita)

		case 5:
			colita.RepoCol()
			pila.Graph()

		case 6:
			fmt.Println("Cerrando sesion")
			fmt.Println("")
			return

		}

	}
}

func Cargamasiva(colita *colita.Cola) {
	var dato string
	fmt.Println("═══════════════════════ CARGA MASIVA DE ESTUDIANTES ══════════════════════ ")
	fmt.Println("")
	fmt.Println("❤ INGRESE NOMBRE DEL ARCHIVO: ")
	fmt.Scanln(&dato)

	file, err := os.Open("C:\\Users\\genes\\Desktop\\" + dato + ".csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}

	defer file.Close()
	registrar := csv.NewReader(file)
	if _, err := registrar.Read(); err != nil {
		fmt.Println(err)
		return
	}
	for {
		// Leer una línea del archivo
		row, err := registrar.Read()
		if err != nil {
			// Si hemos llegado al final del archivo, salir del bucle
			if err == io.EOF {
				break
			}
			// Si hay un error diferente, mostrarlo y salir del programa
			fmt.Println(err)
			return
		}
		words := strings.Split(strings.TrimSpace(row[1]), " ")
		num, err := strconv.Atoi(row[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		colita.Agregar(num, words[0], words[1], row[2])
	}
}

//func registro() {

//}

/*
func Cargamasiva(colita *colita.Cola) {
	var dato string
	fmt.Println("═══════════════════════ CARGA MASIVA DE ESTUDIANTES ══════════════════════ ")

	fmt.Println("❤ INGRESE NOMBRE DEL ARCHIVO: ")
	fmt.Scanln(&dato)

	file, err := os.Open("C:\\Users\\genes\\Desktop\\proyecto2.1\\datos.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}

	defer file.Close()
	registrar, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
	for _, registro := range registrar {
		colita.Agregar(registro[0], registro[1], registro[2])
	}

*/
