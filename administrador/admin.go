package administrador

import (
	//"encoding/csv"
	"encoding/csv"
	"fmt"
	"os"
	"time"

	//"os"

	"proyecto.com/proyecti/PPila"
	"proyecto.com/proyecti/colita"
)

func Menuadmi(colita *colita.Cola, pila *PPila.PilaA) {

	var opcion int
	validar := 0
	var nom, cont, Ap string
	var car string
	for opcion != 5 {

		fmt.Println("═══════════════════════ ADMINISTRADOR - EDD GODRIVE ══════════════════════ ")
		fmt.Println("❤                   1. VER ESTUDIANTES ELIMINANDO                        ❤")
		fmt.Println("❤                   2. VER ESTUDIANTES DEL SISTEMA                       ❤")
		fmt.Println("❤                   3. REGISTRAR NUEVOS ESTUDIANTES                      ❤")
		fmt.Println("❤                   4. CARGA MASIVA DE ESTUDIANTES                       ❤")
		fmt.Println("❤                   5. CERRAR SESION                                     ❤")
		fmt.Println("═══════════════════════════════════════════════════════════════════════════ ")

		fmt.Println("Ingrese una opcion:   ")
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			validar = 0
			for validar != 3 {
				if !colita.Vacia() {
					//colita.Graph()
					temp, size := colita.Eliminar()

					fmt.Println("═══════════════════════ ESTUDIANTES PENDIENTES ═════════", size+1, "═════════")

					fmt.Printf("Carnet: %s\n", temp.Carnet)
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
						pila.Push(t, "se acepto a el estudiante")

					case 2:
						date := time.Now()
						t := fmt.Sprintf("%02d/%02d/%02d      %02d:%02d     ", date.Day(), date.Month(), date.Year(), date.Hour(), date.Minute())

						fmt.Println("SE RECHAZO AL ESTUDIANTE")
						fmt.Println("")
						pila.Push(t, "se rechazo a el estudiante")
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
				pila.Print()

			}

		case 2:
			fmt.Println("Estudiantes del sistema")
			fmt.Println("")

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
	registrar, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
	for _, registro := range registrar {
		fmt.Println("")
		//fmt.Println(registro[1])

		colita.Agregar(registro[0], registro[1], registro[2], registro[3])

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
