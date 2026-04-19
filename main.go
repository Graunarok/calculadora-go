package main

import (
	"github.com/gofiber/fiber/v2"
)

// Definimos constantes para valores que no cambian
const Port = ":3000"
const AppName = "Calculadora API"

// Estructura para recibir los datos del JSON
type OperacionRequest struct {
	Num1      float64 `json:"num1"`
	Num2      float64 `json:"num2"`
	Operacion string  `json:"op"`
}

func main() {
	// Creamos la instancia de Fiber
	// fiber.New() crea el "motor" de nuestra aplicación web
	app := fiber.New()

	// Ruta principal de bienvenida
	// Definimos una ruta GET (cuando entras desde el navegador al "/" inicial)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Bienvenido a " + AppName)
	})

	// Ruta para calcular
	// Definimos una ruta POST para recibir datos y procesar el cálculo
	app.Post("/calcular", func(c *fiber.Ctx) error {
		// Declaramos una variable de tipo OperacionRequest para guardar los datos recibidos
		// 'datos' es una variable que contendrá los números y la operación
		var datos OperacionRequest

		// Intentamos parsear el cuerpo de la petición (JSON)
		// BodyParser toma el JSON que envía el usuario y lo "llena" en nuestra variable 'datos'
		if err := c.BodyParser(&datos); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Datos inválidos"})
		}

		var resultado float64

		switch datos.Operacion {
		case "+":
			resultado = datos.Num1 + datos.Num2
		case "-":
			resultado = datos.Num1 - datos.Num2
		case "x":
			resultado = datos.Num1 * datos.Num2
		case "div":
			if datos.Num2 == 0 {
				return c.Status(422).JSON(fiber.Map{"error": "División por cero no permitida"})
			}
			resultado = datos.Num1 / datos.Num2
		default:
			return c.Status(400).JSON(fiber.Map{"error": "Operación no soportada. Use: +, -, x, div"})
		}

		return c.JSON(fiber.Map{
			"app":       AppName,
			"resultado":  resultado,
			"operacion": datos.Operacion,
		})
	})

	// Iniciamos el servidor usando nuestra constante
	app.Listen(Port)
}