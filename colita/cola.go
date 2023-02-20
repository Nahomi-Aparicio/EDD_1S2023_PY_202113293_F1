package colita

import (
	"fmt"

	"proyecto.com/proyecti/estructura"
)

type Cola struct {
	primero *NODOcola
	size    int // para indicar e tamaño de la cola de ususarios
}

func (c *Cola) Agregar(carnet string, nombre string, contraseña string) {
	fmt.Println("Agregando estudiante a la cola")
	//creo un estudiante
	nuevoEst := &estructura.Estudiante{Carnet: carnet, Nombre: nombre, Contraseña: contraseña}

	//creo un nodo
	nuevonodo := &NODOcola{Estudiante: nuevoEst, Siguiente: nil}
	//verifico si la cola esta vacia
	if c.primero == nil {
		c.primero = nuevonodo
	} else {
		// si no esta vacia recorro la cola hasta el ultimo nodo
		aux := c.primero
		for aux.Siguiente != nil {
			aux = aux.Siguiente
		}
		//agrego el nuevo nodo al final de la cola
		aux.Siguiente = nuevonodo
		c.size += 1
	}

}

// aqui imprimimos  yey
func (c *Cola) Imprimir() {

	//recorremos la cola
	aux := c.primero
	//verificamos que no este vacia
	for aux != nil {
		//imprimimos el nodo
		aux.Estudiante.Mostrar()
		//pasamos al siguiente nodo
		aux = aux.Siguiente
	}
}

func (c *Cola) Eliminar() (*estructura.Estudiante, int) {

	//verificamos si la cola esta vacia
	if c.primero != nil {

		temp := c.primero
		//si no esta vacia asignamos el segundo nodo como primero
		c.primero = c.primero.Siguiente
		c.size -= 1
		// retornamos el nodo eliminado
		return temp.Estudiante, c.size
	}

	//si esta vacia retornamos nil
	return nil, c.size

}

func (c *Cola) Vacia() bool {
	return c.primero == nil
}










/*func (c *Cola) Graph() {
	file, err := os.Create("cola.dot")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString("digraph G{\n")
	file.WriteString("rankdir=LR;\n")
	file.WriteString("node [shape=record];\n")
	file.WriteString("subgraph cluster_0 {\n")
	file.WriteString("style=filled;\n")
	file.WriteString("color=lightgrey;\n")
	file.WriteString("node [style=filled,color=white];\n")
	file.WriteString("label = \"Cola\";\n")
	file.WriteString("N0[label=\"carne" + c.primero.Estudiante.Nombre + "\"];\n")
	file.WriteString("N0->N1;\n")
	file.WriteString("N1[label=\"" + c.primero.Siguiente.Estudiante.Nombre + "\"];\n")
	file.WriteString("N1->N2;\n")
	file.WriteString("N2[label=\"" + c.primero.Siguiente.Siguiente.Estudiante.Nombre + "\"];\n")
	file.WriteString("N2->N3;\n")
	file.WriteString("N3[label=\"" + c.primero.Siguiente.Siguiente.Siguiente.Estudiante.Nombre + "\"];\n")
	file.WriteString("N3->N4;\n")
	file.WriteString("N4[label=\"" + c.primero.Siguiente.Siguiente.Siguiente.Siguiente.Estudiante.Nombre + "\"];\n")
	file.WriteString("N4->N5;\n")
	file.WriteString("N5[label=\"" + c.primero.Siguiente.Siguiente.Siguiente.Siguiente.Siguiente.Estudiante.Nombre + "\"];\n")
	file.WriteString("N5->N6;\n")
	file.WriteString("N6[label=\"" + c.primero.Siguiente.Siguiente.Siguiente.Siguiente.Siguiente.Siguiente.Estudiante.Nombre + "\"];\n")
	file.WriteString("N6->N7;\n")
	return
}




digraph G{
rankdir=LR;
node [shape=record];
subgraph cluster_0 {
style=filled;
color=lightgrey;
node [style=filled,color=white];
label = "Cola";
N0[label="carnemangera"];
N0->N1;
N1[label="marcelino"];
N1->N2;






*/

// / HASTA AQUI LLEGE HOY DEBO DE VER COMO IMPRIMIR ESTO EN EL MAIN NO ESTA TERMINADO  LAS CARPETAS COLITA Y ESTRUCTURA
