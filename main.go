package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Generamos un bucle for y numeros aleatorios
	for {
		t := temp()
		localTemp := generateTemp(t)
		//_temp := strconv.Itoa(localTemp.Temp)
		url := fmt.Sprintf("http://localhost:9192/temperatures/local/%d", localTemp.Temp)
		_, err := http.Post(url, "text/plain", nil)

		if err != nil {
			log.Printf("No puede enviar temperatura => %s ", err.Error())
		}
		log.Printf("Temperatura enviada => %d ", localTemp.Temp)
		time.Sleep(time.Second * 5)
	}

}

type LocalTemperature struct {
	Temp int `json:"temp"`
}

func generateTemp(t int) LocalTemperature {
	return LocalTemperature{Temp: t}
}

// Generamos una temperatura n entero aleatorio
func temp() int {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	return r.Intn(40)
}
