# meli-go

## Arrays
* Arrays son de longitud fija.

## Slices
* Slice, es un puntero a un array, por lo tantosu zero-value es un nil. Es un array sin longitud fija.
* `make()`, crea punteros a cosas.
* `a := make([]int, 0)`, crea un slice vacio.
* longitud: `len()`, nro de elemento del array.
* capacidad: `cap()`, nro de elementos del arrat subyavente.
* el operador `:` sirve para mover el puntero de inicio del slice.
* `s := arr[ :0 ]`, dame punteros del array desde la nada hasta el elemento cero.
* `s := s[ :4 ]`, dame hasta el cuarto elemento. 
* `s := s[ 2: ]`, remover los 2 primeros elementos.
* **Agregar elementos al slice**: `append(array, elementosQueQuieroAgregar)`, los agrega al final. 

## Metodos
* Función asociada a un método.
* (t TipoStruct), se llama receiver.

## Interfaces
* El tipo `var asdf interface{}`, es cualquier cosa. Para validar que tipo de dato es podemos hacer `asdf.(type)` y obtener su tipo. O podemos hacer `valor, ok := asdf.(int)` (el parametro ok es importante porque si no puede castearlo paniquea).
* Los structs que implementan los metodos de una interface, ya implementan automaticamente la interfaz. No hay notación especial para esto.
* **Todas las interfaces son punteros** 

## Estructuras embebidas
* composición por sobre herencia (no existe en Go): un struct A que tiene como atributo otro struct B, podemos acceder a las propiedades de B directamente como si fueran parte de A.

## Concurrencia
### Go Routine 
* Una Go Rutina es un "hilo liviano" administrado por Go, que crea un pull de threads y administra las rutinas en ese pull, por eso una rutina no es exactamente un hilo.
* El Runtime de Go decide crear y administrar los hijos necesarios.
* Para ejecutar una funcion concurrente se hace `go f(x, y)`.

### Channels
* Conductos tipados de coumunicación entre rutinas. Que sean tipados quiere decir que si declaro un canal de ints, solo puedo enviar ints por ese canal.
* `ch := make(chan int, capacidad)` declaro un canal de ints.
* Son bidireccionales: 
    - puede enviar `ch <- variable` 
    - y recibir  info 
        + `variable := <- ch` (paniquea si el canal está cerrado) 
        + o `variable, ok := <- ch` (no paniquea) info al canal.
* Es un punto de sincronización entre Go Rutinas, son sincrónicos: se bloquea el sender si el canal está lleno, se bloquea el receiver hasta que llegue un mensaje.
* Para leer infinitamente de un canal podemos usar:
    - `for v := range ch`. 
* Cuando el dueño del canal termina de escribir debería cerrarlo `close(ch)`.

### Select
* **select** es como un switch sobre canales. No hay garantía del orden de ejecución entre los *cases* del **select**.

### Patron Fan In
* Problema: veo como si se ejecutaran secuencialmente pero en realidad se ejecutan aleatoriamente. 
* Solución: se crea un canal que es concentrador de los otros 2 canales. Este concentrador se podría hacer un **select** o creamos N nuevas rutinas que reciban mensajes de todos los canales y los ponga en el canal homogéneo.

