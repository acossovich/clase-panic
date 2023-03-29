package main

import (
	"fmt"
	"os"
	"io"
)

// Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones. Para ello, cuentan con todo el detalle necesario en un archivo .txt. 
// Tendrás que desarrollar la funcionalidad para poder leer el archivo .txt que nos indica el cliente, sin embargo, no han pasado el archivo a leer por nuestro programa. 
// Desarrollá el código necesario para leer los datos del archivo llamado “customers.txt” (recordá lo visto sobre el pkg “os”).
// Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso, el programa deberá arrojar un panic al intentar leer un archivo que no existe, mostrando el mensaje “el archivo indicado no fue encontrado o está dañado”.
// Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.

func abrirArchivo(nombre string) (*os.File, error){
	
	defer func(){
		dato := recover()

		if dato != nil{
			fmt.Println(dato)
		}		
	}()
	
	
	f1, err := os.Open(nombre)
	if err != nil{
		panic("El archivo indicado no fue encontrado o esta dañado")
	}

	return f1,err
}

// A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio. 
// Ahora que el archivo sí existe, el panic no debe ser lanzado. 
// Creamos el archivo “customers.txt” y le agregamos la información de los clientes. 
// Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga. En el caso de no poder leerlo, se debe lanzar un “panic”.
// Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener un “defer” que nos indique que la ejecución finalizó. También recordemos cerrar los archivos al finalizar su uso.

func leerArchivo(archivo *os.File) (string, error){

	defer func(){
		err := recover()

		if err != nil{
			fmt.Println(err)
		}
	}()

	texto,err := io.ReadAll(archivo)

	if err != nil{
		panic("El archivo no se pudo leer:" + err.Error())
	}

	return string(texto),nil

}

func main(){
	fmt.Println("Iniciando... ")

	fileCustomers,ok := abrirArchivo("customers.txt")

	if ok != nil{
		fmt.Println(ok)
	}

	datosArchivo,_ := leerArchivo(fileCustomers)

	fmt.Println(datosArchivo)

	fmt.Println("Ejecucion finalizada")
}