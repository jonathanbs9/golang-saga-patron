package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
)

var d int

func main() {

	app := fiber.New()
	// Creo la ruta /temperatures
	r := app.Group("/temperatures")
	// /temperatures/local
	r.Get("/local", func(c *fiber.Ctx) {
		c.JSON(buildDegrees())
	})

	r.Post("/local/:temp", func(c *fiber.Ctx) {
		temperature := c.Params("temp")
		_temperature, err := strconv.Atoi(temperature)

		if err != nil {
			c.Status(400)
			log.Printf(err.Error())
			return
		}
		// Si todo fue ok, tengo la temp
		log.Printf("La temperatura enviada es : %d", _temperature)
		setD(_temperature)
	})

	log.Fatal(app.Listen(9192))
}

type Temperature struct {
	Degrees int `json:"degrees"`
}

// buildDegrees func =>  returns a temperature
func buildDegrees() Temperature {
	return Temperature{
		Degrees: d,
	}
}

func setD(_d int) {
	d = _d
}
