package main

import "fmt"

func main() {

	fmt.Println("Este es el curso de Platzi.")
	fmt.Println("Elegir el número cuál de las funcionalidades implementadas quieres ejecutar.")
	fmt.Println("1: Variables")
	fmt.Println("2: Aritmética")
	fmt.Println("3: Librería fmt")
	fmt.Println("4: Funciones")
	fmt.Println("5: Ciclo FOR")
	fmt.Println("6: Condicionales")
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
