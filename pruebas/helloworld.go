package main

//package main -> es ejecutable que contiene la funcion main
//todo .go debe empezar con un package <A>
//todos los archivos en una carpeta pertenecen al mismo package.
//La unica excepcion es que puede haber un package NOMBRE_test dentro de la misma carpeta
//los packages pueden no coincidir con el nombre de la carpeta
//por paquete no puedes tener funciones repetidas
//en un package hay un conjunto de funciones

import (
	"fmt"
	"math/rand"
)

var c, java, python bool
var pi float32 = 3.14
var i = false

func main() {
	var a, b, d = true, false, "este un string no un boolean!" //infiere, pero una vez inferido no puedes asignarle algo de otro tipo
	fmt.Printf("Mi nro favorito es ", rand.Intn(10))           //Las funciones con mayuscula (rand.Intn()) son Exports, las minisculas son privadas
	fmt.Printf("\n")
	var i int //Go tiene default, int=0.
	//Esta variable i tiene mas scope que la var global i definida como bool.
	//int es alias de int32 o int64 que lo infiere segun la maquina
	//tipo "rune", alias de int32. Representa un code point de Unicode
	fmt.Printf("vars: ", i, c, java, python, pi)
	fmt.Printf("\n")
	fmt.Printf("var multideclaration: ", a, b, d)
	fmt.Printf("\n")
	fmt.Printf("Sum: ", Sum(1, 1))
	fmt.Printf("\n")
	var a1, b1 string = intercambiar("Mercado", "Libre")
	fmt.Printf("intercambiar: ", a1, b1)
	fmt.Printf("\n")
	var _, b2 string = intercambiar2("Mercado", "Libre") // Identificador vacío: no me interesa la primer palabra
	fmt.Printf("intercambiar2: %b", b2)
	fmt.Printf("\n")
	fmt.Printf("Conversion(casteo): ", float32(1))
	fmt.Println()
	punteros()
}

// Sum ints
func Sum(x, y int) int {
	return x + y
}

// intercambiar devuelve multiples variables
func intercambiar(x, y string) (string, string) {
	return y, x
}

//intercambiar usando nombres a variables de retorno
func intercambiar2(x, y string) (primera string, segunda string) {
	primera = y
	segunda = x
	return
}

func punteros() {
	i := 42

	var p = &i
	fmt.Printf("puntero: ", p)
	fmt.Println()
	fmt.Printf("valor: %b", *p)
	fmt.Println()
	*p = 21
	fmt.Printf("valor modificado: %b", *p)
}
