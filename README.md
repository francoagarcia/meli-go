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

# Links utiles

https://dave.cheney.net/resources-for-new-go-programmers Varios recursos para empezar

https://github.com/golang/go/wiki LA wiki de go

https://github.com/EbookFoundation/free-programming-books/blob/master/free-programming-books.md#go Textos gratuitos

https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1 Una manera de organizar los paquetes

https://peter.bourgon.org/go-best-practices-2016/ Tips y experiencias de uso

https://github.com/tmrts/go-patterns Listado de patrones con algunas implementaciones

https://github.com/miguellgt/books/blob/master/go/go-design-patterns.pdf Libro sobre patrones

https://github.com/jbuberel/go-patterns Ejemplos de patrones (puede estar desactualizado)

https://talks.golang.org/2014/names.slide#1 Consejos sobre naming

https://blog.golang.org/slices Sobre slices y arrays

https://research.swtch.com/godata Cómo funcionan internamente las estructuras

https://golang.org/doc/effective_go.html Un montón de info sobre de todo un poco

https://talks.golang.org/2014/organizeio.slide#1 Cómo organizar el código en paquetes

https://dmitri.shuralyov.com/idiomatic-go conjunto de buenas prácticas

https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back sobre punteros

http://kunalkushwaha.github.io/2015/09/11/understanding-golang-interfaces/ Sobre interfaces

http://golangtutorials.blogspot.com.ar/2011/06/interfaces-in-go.html Tutorial de interfaces

https://www.golang-book.com/books/intro/9 Interfaces y tipos embebidos

https://talks.golang.org/2012/concurrency.slide#1 Patrones de concurrencia

http://guzalexander.com/2013/12/06/golang-channels-tutorial.html Tutorial de canales

https://go-talks.appspot.com/github.com/davecheney/high-performance-go-workshop/high-performance-go-workshop.slide#1 Workshop sobre performance

https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go Cómo escribir benchmarks

https://github.com/google/pprof/blob/master/doc/pprof.md Profiling

https://blog.golang.org/defer-panic-and-recover Defer, panic, recover

http://www.codegist.net/snippet/shell/golang-total-coverage-html-output_toefel18_shell script para medir cobertura en todos los paquetes

https://talks.golang.org/2009/go_talk-20091030.pdf
