UNIVERSIDAD SAN CARLOS DE GUATEMALA
<br>FACULTAD DE INGENIERIA 
<br>ESCUELA DE CIENCIAS Y SISTEMAS
<br>ESTRUCTURA DE DATOS 
![Image text](https://upload.wikimedia.org/wikipedia/commons/4/4a/Usac_logo.png)
<div style ="text-aling:rigth"> GENESIS NAHOMI APARICIO ACAN
  202113293
  SECCION: B
  </div>
   
  
<div style ="text-aling:justify">
  <p>    
    Este programa fue desarrollado en el lenguaje de programacion Go(golang). El objetivo principal de este programa fue el crear una forma de almacenar archivos importantes. Siendo capaz de usarse en cualquier sistema operativo. En el presente manual se descriven los metodos,funciones y clases que fueron utilizadas en el desarrollo del proyecto en donde se pueden  observar los distintos tipod de datos que se usaron .Al igual que se hizo uso de la herramienta Graphviz para graficar las estructuras de dato de el mismo modo se hizo uso de varias librerias para leer archivos  como para generar los reportes 
  </p>
  </div>
    <div style ="text-aling:center">
  <p>
    <h3><b> LOGICA DEL PROGRAMA</b></h3>.
</p>
  </div>
  
<div style ="text-aling:justify">
  <p>    
   Se usaron 6 paquetes  y varios metodos para realizar este programa entre los paquetes que se usaron estan :
    <br>administrador:cuenta con un archivo admin.go
    <br>colita: cuenta con dos archivos  cola.go y nodo.go
   <br> listaDo: cuenta con dos archivos listaDD.go y nododo.go
     <br> estructura:cuenta con un archivo estruct.go
   <br> PilaDoble: cuenta con dos archivos nodo.go y pilaDo.go
    <br>PPila : cuenta con dos archivos nodo.go y pilaA.go
     <br>fuera de estos paquetes se encuenta un archivo llamado go.mod y el archivo main.go que son indispensables para el funcionamiento del programa        
  </p>
  </div>
</div>
    <div style ="text-aling:center">
  <p>
    <h3><b> MAIN.GO </b></h3>.
</p>
  </div>
  <div style ="text-aling:justify">
  <p>    
  Este archivo contiene en su interior la funcion main la cual hace que el programa comience a funcionar a la hora de iniciarlo, cuenta con dos metodos incluyendo la main. Aqui se muestra el primer menu el cual es el menu de inicio de secion 
  </p>
  </div>


   ![Image text](https://raw.githubusercontent.com/Nahomi-Aparicio/EDD_1S2023_PY_202113293/main/EDD_Proyecto1_Fase1/imagenesdel%20manual/main.PNG?token=GHSAT0AAAAAAB7AG43LOHR5SBINGYJXNK6MY76RTAA)
 
 
<div style ="text-aling:center">
  <p>
    <h3><b> ADMIN.GO </b></h3>.
</p>
  </div>
  
  
<div style ="text-aling:justify">
  <p>    
  Este archivop contiene los menus donde el administrador puede gestionar a los usuario , tiene dos metodos uno llamado Menuadmi el cual tiene como parametros la cola,la pila y la lista doble que son usadas en estre programa ademas  en este se pueden aceptar rechazar e incluir a los usuarios a sus respecxtivas pilas, de igual forma llama a los metodos de otros archivos para mostrar los usuarios registrados o para generar los respectivos reportes en "dot" , "png" y "Json".
    </p>
  </div>

   ![Image text](https://raw.githubusercontent.com/Nahomi-Aparicio/EDD_1S2023_PY_202113293/main/EDD_Proyecto1_Fase1/imagenesdel%20manual/admin.PNG?token=GHSAT0AAAAAAB7AG43LMX76EHQXLH5YY2QIY76RTKA)
 

 
  <div style ="text-aling:justify">
  <p>    
 Este archivo de igual forma tiene otro metodo llamado Carga masiva el cual lee el nombre del archivo el cual se desea ingresar al sistema y separa los datos por comas y por espacios que este tenga para guardarlo en la cola de estudiantes pendientes 
    </p>
  </div>
 
 
 
 
 
  ![Image text](https://raw.githubusercontent.com/Nahomi-Aparicio/EDD_1S2023_PY_202113293/main/EDD_Proyecto1_Fase1/imagenesdel%20manual/carga%20masiva.PNG?token=GHSAT0AAAAAAB7AG43LGTPSJDVPYWVXY7YWY76RTUQ)
  
  <div style ="text-aling:center">
  <p>
    <h3><b> ESTRUCT.GO </b></h3>.
</p>
  </div>
  
   <div style ="text-aling:justify">
  <p>    
Este archivo contiene la estructutura de cuales seran los datos que seran ingresados en el sistema teniendo un metodo que mostrara el carnet, el nombre , apellido y la contraseña del nuevo usuario o el usuario que ya se encuentra en el sistema
    </p>
  </div>
  
   ![Image text](https://raw.githubusercontent.com/Nahomi-Aparicio/EDD_1S2023_PY_202113293/main/EDD_Proyecto1_Fase1/imagenesdel%20manual/estudiantes.PNG?token=GHSAT0AAAAAAB7AG43KWNO22UQXBEIQ3DCWY76RUBA)
  
  
 <div style ="text-aling:center">
  <p>
    <h3><b> COLA.GO </b></h3>.
</p>
  </div>

   <div style ="text-aling:justify">
  <p>    
  Este archivo contiene 5 metodos agregar,eliminar, vacia, graph1 y repocol.
  <br>  Agregar: en este metodo agrega los estudiantes ya sea de carga masiva o de forma manual a una cola de forma ordenada
   <br> Eliminar: los elementos en la cola se eliminan conforme el primero en salir es el primero que entro 
     <br> vacia: este es un metodo tipo bool el cual nos indica si la cola esta vacia o si aun existen datos en ella
     <br> graph1 y repocol: estos metodos son usados para generar lso reportes de cola ya sea en un archivo .dot y un archivo.png
    </p>
  </div>
  
 ![Image text](https://raw.githubusercontent.com/Nahomi-Aparicio/EDD_1S2023_PY_202113293/main/EDD_Proyecto1_Fase1/imagenesdel%20manual/cola.PNG?token=GHSAT0AAAAAAB7AG43LFYX2XUTFRHCEVQB4Y76RUVQ)
 
 
 <div style ="text-aling:center">
  <p>
    <h3><b> PILAA.GO </b></h3>.
</p>
  </div>
  
  
   <div style ="text-aling:justify">
  <p>    
Este archivo contiene los metodos Push,VerificarCar,Print,Graph:
    <br>Push : este metodo es igual que el de agregar de cola solo que los datos se incertan al inicio del la pila no al final de esta
    <br> Print: este metodo muestra los datos que se encuentran en la pila 
    <br>Graph: este metodo grafica los reportes  que genera el archivo mostrando los estudiantes que fueron aceptados o rechazados conforme fueron entrando estos datos
    <br> VerificarCar: este metodo verifica si los carnets de los nuevos usuarios registrados son repeito o no     </p>
</div>
  
  
![Image text](https://raw.githubusercontent.com/Nahomi-Aparicio/EDD_1S2023_PY_202113293/main/EDD_Proyecto1_Fase1/imagenesdel%20manual/pilaA.PNG?token=GHSAT0AAAAAAB7AG43LAQ5UUSW4OCTUVECGY76RP3A)


<div style ="text-aling:center">
  <p>
    <h3><b> NODODO.GO </b></h3>.
</p>
  </div>
  <div style ="text-aling:justify">
  <p>    
Este archivo contiene la estructutura del nodo de la lista doble en el cual tiene como apuntadores sigiente, anterior y pila al igual que contiene los 4 datos de el susuario y un metodo el cual muesta estros cuatro datos 
    </p>
  </div>
  
 ![Image text](https://raw.githubusercontent.com/Nahomi-Aparicio/EDD_1S2023_PY_202113293/main/EDD_Proyecto1_Fase1/imagenesdel%20manual/nododo.PNG?token=GHSAT0AAAAAAB7AG43L2AT3PI7MG4N4DHI6Y76RVDQ)


<div style ="text-aling:center">
  <p>
    <h3><b> LISTADD.GO </b></h3>.
</p>
  </div>
  
   <div style ="text-aling:justify">
  <p> 
    Este archivo contiene la lista doble el cual contiene varios metodos como el Añadir, Imprimir, BuscaryagregarHora, OrdenarPorCarnet, CrearJon y Graficar
    <br>
  <br> Añadir: añade los usuario segun como estos son aceptados al sistema a la lista doble
    <br> Imprimir:muestra los usuarios en el sistema cuando el administrador quiera ver cuales son estos   
   </p>
</div>

 ![Image text](https://raw.githubusercontent.com/Nahomi-Aparicio/EDD_1S2023_PY_202113293/main/EDD_Proyecto1_Fase1/imagenesdel%20manual/listaDD.PNG?token=GHSAT0AAAAAAB7AG43L6FAOCD2STBEQJWF4Y76RVQA)

  <div style ="text-aling:justify">
  <p> 
 <br>  BuscaryagregarHora: es una verificacion el cual añade la hora y la fecha que los usuarios ingresaron al sistema y los guarda en la pila que esta como un parametro en esta lista
  <br>  OrdenarPorCarnet:ordena los usuario segun el numero de carnet que estos poseen 
   <br>  CrearJon:crea el JSon de reporte pormedio de algunas librerias 
   <br>  Graficar:al igual que en pila y cola este metodo se utiliza para poder graficar los reportes de lista doble en donde muestra los usuarios en orden y la pila de cuando cada uno de estos ingresaron al sistema
     </p>
</div>

 ![Image text](https://raw.githubusercontent.com/Nahomi-Aparicio/EDD_1S2023_PY_202113293/main/EDD_Proyecto1_Fase1/imagenesdel%20manual/listaDD.PNG?token=GHSAT0AAAAAAB7AG43KBGNPGZRX57MFQSX2Y76RQ3A)
