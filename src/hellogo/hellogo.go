// go install
// go run hellogo.go
// go build src\hellogo\hellogo.go -> he tenido que hacerlo desde src
// hellogo.exe

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello Go!")

	// declarar ints
	var x int = 5
	var y int
	y = 0
	z := 5

	// los arrays empiezan en [0]
	var a [5]int
	a[0] = 1
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// los slices no tienen longitud fija
	c := []int{1, 2, 3}
	c = append(c, 4)

	// maps son los hashmap o diccionarios en pyhon
	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["pentagon"] = 5

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])

	delete(vertices, "pentagon")
	fmt.Println(vertices)

	// loops
	for i := 0; i < 5; i++ {
		fmt.Println("Hello Go!")
	}

	for y < 3 && z == 5 {
		fmt.Println("Hello Go!")
		y++
	}

	// para arrays
	for index, value := range b {
		fmt.Println("index: ", index, "value: ", value)
	}

	// para maps
	for key, value := range vertices {
		fmt.Println("key: ", key, "value: ", value)
	}

	// condicionales
	if x > 3 {
		fmt.Println("Hello Go!")
	}

	// llamada a funciones
	result := sum(z, y)
	fmt.Println(result)

	res, err := sqrt(81)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	// declaracion de objetos y acceso a sus propiedades
	p := person{name: "Joshua", age: 32, hobby: "Golang"}
	fmt.Println(p)
	fmt.Println(p.hobby)

	// punteros
	fmt.Println(&x) // direccion de memoria donde está x
	k := 1
	increment(k)
	fmt.Println(k) // no cambia k porque no hacemos ningun return

	increment_pointer(&k)
	fmt.Println(k)

}

// func nombrefunc (arg1 tipo1, arg2 tipo2) returntype
func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers") // go no tiene excepciones, pero puedes crear tus propios errores
	}

	return math.Sqrt(x), nil // nil indica que no hay error
}

// creacion de objetos propios
type person struct {
	name  string
	age   int
	hobby string
}

func increment(x int) {
	x++ // no devolvemos nada
}

func increment_pointer(x *int) {
	*x++ // no devolvemos nada, sin * aumentamos la direccion de memoria, con * aumentamos el valor guardado en esa dir
}

// float64
// float32
// int
