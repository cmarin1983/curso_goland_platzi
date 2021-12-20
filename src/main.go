package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Este es el curso de Platzi.")
	fmt.Println("Elegir el número cuál de las funcionalidades implementadas quieres ejecutar.")
	fmt.Println("1: Variables")
	fmt.Println("2: Aritmética")
	fmt.Println("3: Librería fmt")
	fmt.Println("4: Funciones")
	fmt.Println("5: Ciclo FOR")
	fmt.Println("6: Condicionales")
	fmt.Println("7: Keywords")
	fmt.Println("8: Arreglos")
	fmt.Println("9: Recorrer arreglo")
	fmt.Println("10: Llave valor")
	fmt.Println("11: Como son las clases en GO => Structs")
	fmt.Print("Opción: ")
	var n int
	fmt.Scanln(&n)

	switch n {
	case 1:
		//Llama a función para ver las variables
		variables()
	case 2:
		//Llama a función de aritmética
		aritmetica()
	case 3:
		//Libreria fmt
		libreriaFmt("Platzi")
	case 4:
		//Funciones
		funciones()
	case 5:
		//Ciclo FOR
		ciclos()
	case 6:
		//Condicionales
		condicional()
	case 7:
		//Palabras clave
		keywords()
	case 8:
		//Arreglos
		arreglos()
	case 9:
		//range
		sliceRange()
	case 10:
		llaveValor()
	case 11:
		//Aca se llaman structs
		clases()
	}

}

type car struct {
	brand string
	year  int
}

func clases() {
	myCar := car{brand: "Toyota", year: 2021}
	fmt.Println(myCar)

	var otherCar car
	otherCar.brand = "Nissan"
	otherCar.year = 2012
	fmt.Println(otherCar)
}

func llaveValor() {
	m := make(map[string]int)

	m["Arianna"] = 10
	m["Danna"] = 9

	fmt.Println(m)

	// recorrer mapa
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar valor en mapa
	valor, ok := m["Arianna"]
	fmt.Println(valor, ok)
}

func isPalindromo(texto string) {
	var textoReverso string
	texto = strings.ToLower(texto)

	for i := len(texto) - 1; i >= 0; i-- {
		textoReverso += string(texto[i])
	}

	if texto == textoReverso {
		fmt.Println("Es palindromo.")
	} else {
		fmt.Println("No es palindromo.")
	}
}

func sliceRange() {
	slice := []string{"hola", "que", "hace"}
	fmt.Println(slice)
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	isPalindromo("AmA")
}

func arreglos() {
	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)
	fmt.Println(len(array), cap(array))

	//Slides
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(slice[3])
	fmt.Println(slice[:3])
	fmt.Println(slice[3:5])
	fmt.Println(slice[4:])

	//Append
	slice = append(slice, 8)
	fmt.Println(slice)

	//Append nueva lista
	newSlice := []int{9, 10, 11}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}

func keywords() {
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break
		if i == 8 {
			fmt.Println("break")
			break
		}

	}
}

func ciclos() {

	//En reversa
	counterReverse := 10
	for counterReverse >= 0 {
		fmt.Println(counterReverse)
		counterReverse--
	}

	//Condicional for
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//For forever
	/*counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}*/
}

func funciones() {
	normalFuncion("Hola mundo.")
	tripleFuncion(45, 88, "números")

	value := returnFuncion(4)
	fmt.Println("Valor es: ", value)

	valor1, valor2 := dosValores(4)
	fmt.Println("Valor1 y Valor2 son:", valor1, valor2)

	valor3, _ := dosValores(4)
	fmt.Println("Valor3 es:", valor3)
}

func normalFuncion(message string) {
	fmt.Println(message)
}

func tripleFuncion(a, b int, c string) {
	fmt.Printf("El %d y %d son %s.\n", a, b, c)
}

func returnFuncion(a int) int {
	return a * 2
}

func dosValores(a int) (c, d int) {
	return a, a * 2
}

func libreriaFmt(mensaje string) {
	//Declaración de variables
	helloMessage := "Hello"
	worldMessage := "world"

	//Prinln
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Prinf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos.\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos.\n", nombre, cursos)

	//SPrintf
	message := fmt.Sprintf("%v tiene más de %v cursos.", nombre, cursos)
	fmt.Println(message)

	//Tipo de dato
	fmt.Printf("helloMessge: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	fmt.Println("Cuál es el nomnre de este curso?")
	var s string = mensaje
	//fmt.Scanln(&s)
	fmt.Printf("El nombre del curso es %s\n", s)
}

func variables() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Área de un cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Área cuadrado: ", areaCuadrado)
}

func aritmetica() {
	const pi float64 = 3.14
	x := 10
	y := 50
	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = y - x
	fmt.Println("Resta: ", result)

	//Multiplcación
	result = y * x
	fmt.Println("Multiplicación: ", result)

	//División
	result = y / x
	fmt.Println("División: ", result)

	//Módulo (residuo)
	result = y % x
	fmt.Println("Módulo: ", result)

	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Decremental
	x--
	fmt.Println("Decremental: ", x)

	//Calcular el área de rectángulo, trapecio y círculo
	//Rectángulo
	baseRectangulo := 40
	alturaRectangulo := 15
	result = baseRectangulo * alturaRectangulo
	fmt.Println("Área de rectángulo: ", result)

	//Trapecio
	base1Trapecio := 40
	base2Trapecio := 30
	alturaTrapecio := 25
	result = ((base1Trapecio * base2Trapecio) * alturaTrapecio) / 2
	fmt.Println("Área de trapecio: ", result)

	//Círculo
	diametroCirculo := 20
	radioCirculo := diametroCirculo / 2
	areaCirculo := pi * float64(radioCirculo*radioCirculo)
	fmt.Println("Área de círculo: ", areaCirculo)
}

func condicional() {
	valor1 := 1

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//Con and
	valor2 := 2
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("El valor1 es 1 y el valor2 es 2")
	} else {
		fmt.Println("Uno de los valores no es igual.")
	}

	//Si es par o no
	if valor1%2 == 0 {
		fmt.Println("Valor1 es par")
	} else {
		fmt.Println("Valor1 no es par")
	}
	if valor2%2 == 0 {
		fmt.Println("Valor2 es par")
	} else {
		fmt.Println("Valor2 no es par")
	}
}
