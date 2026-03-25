package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 1. Verificamos que escribas la operación completa
	if len(os.Args) < 4 {
		fmt.Println("❌ Faltan datos. Ejemplo de uso: go run main.go 10 + 5")
		return 
	}

	// 2. Capturamos los números y el signo
	// Nota: En Go, el primer número está en la posición [1] (en Node era [2])
	numero1, _ := strconv.ParseFloat(os.Args[1], 64)
	operacion := os.Args[2]
	numero2, _ := strconv.ParseFloat(os.Args[3], 64)

	var resultado float64

	// 3. Hacemos las matemáticas (usamos "x" y "div" igual que en Node)
	if operacion == "+" {
		resultado = numero1 + numero2
	} else if operacion == "-" {
		resultado = numero1 - numero2
	} else if operacion == "x" {
		resultado = numero1 * numero2
	} else if operacion == "div" {
		resultado = numero1 / numero2
	} else {
		fmt.Println("❌ Operación no válida. Usa +, -, x, o div")
		return
	}

	// 4. Mostramos el resultado
	fmt.Printf("✅ El resultado es: %v\n", resultado)
}