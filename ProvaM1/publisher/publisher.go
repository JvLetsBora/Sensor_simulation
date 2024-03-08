package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

type Teste struct {
	testeDuracao int
}

type Loja struct {
	Freezer   SensorData `json:"Freezer"`
	Geladeira SensorData `json:"Geladeira"`
}

type SensorData struct {
	SensorId    string `json:"sensor_id"`
	Timestamp   int64  `json:"timestamp"`
	Tipo        string `json:"Tipo"`
	Temperatura int    `json:"Temperatura"`
	Limite      int
	min         int
	max         int
}

// func (sensor SensorData) getData() int {
// 	return sensor.Limite
// }

func NovaLoja() Loja {
	return Loja{
		Freezer:   NovoSensor("Freezer", -28, 0, "lj98b01"),
		Geladeira: NovoSensor("Geladeira", -2, 20, "lj01f01"),
	}
}

func NovoSensor(tipo string, min int, max int, id string) SensorData {
	return SensorData{
		SensorId:    id,
		Timestamp:   time.Now().Unix(),
		Tipo:        tipo,
		Temperatura: rand.Intn(int(max-min)) + int(min),
		Limite:      8,
		min:         min,
		max:         max,
	}
}

func publiStart(teste Teste) {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("Publisher")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {

		time.Sleep(time.Duration(1 * time.Second))

		currentData := NovaLoja()
		jsonData, _ := json.Marshal(currentData)
		token := client.Publish("test/topic", 1, false, jsonData)
		token.Wait()
		fmt.Println(string(jsonData))

		// if token := client.Publish("sensor/data", 1, true, jsonData); token.Wait() && token.Error() != nil {
		// 	panic(token.Error())
		// }

	}
	select {
	// case <-time.After(time.Second * time.Duration(teste.testeDuracao)):
	// 	fmt.Println("Cliente desconectado.")
	// 	client.Disconnect(250)

	}
}

func main() {
	teste := Teste{testeDuracao: 30}
	publiStart(teste)
}
