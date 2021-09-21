package main

import "fmt"

func main() {

	//Llama a función para ver las variables
	variables()

	//Llama a función de aritmética
	aritmetica()

	//Libreria fmt
	libreriaFmt()
}

func libreriaFmt() {
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
