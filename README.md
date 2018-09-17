# meli-go

## Arrays

## Slices
* Arrays son de longitud fija.
* Slice, es un puntero a un array, por lo tantosu zero-value es un nil. Es un array sin longitud fija.
* `make()`, crea punteros a cosas.
* `a := make([]int, 0)`, crea un slice vacio.
* longitud: `len()`, nro de elemento del array.
* capacidad: `cap()`, nro de elementos del arrat subyavente.
* el operador `:` sirve para mover el puntero de inicio del slice.
* `s := arr[ :0 ]`, dame punteros del array desde la nada hasta el elemento cero.
* `s := s[ :4 ]`, dame hasta el cuarto elemento. 
* `s := s[ 2: ]`, remover los 2 primeros elementos.

## Agregar elementos al slice
* `append(array, elementosQueQuieroAgregar)`, los agrega al final. 
