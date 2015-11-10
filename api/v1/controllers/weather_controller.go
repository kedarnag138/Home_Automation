package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/d2r2/go-dht"
	"github.com/kedarnag13/Home_Automation/api/v1/models"
	"log"
	"net/http"
)

type WeatherController struct{}

var Weather WeatherController

func (w *WeatherController) Monitor(rw http.ResponseWriter, req *http.Request) {
	sensorType := dht.DHT11
	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(sensorType, 4, true, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)
	b, err := json.Marshal(models.WeatherMessage{
		Success:     "True",
		Message:     "Temperature and Humidity updated",
		Temperature: temperature,
		Humidity:    humidity,
	})
	if err != nil {
		log.Fatal(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(b)
}
