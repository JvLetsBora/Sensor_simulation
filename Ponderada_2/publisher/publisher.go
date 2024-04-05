package main

import (
	"fmt"
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
	TesteDuracao int
	Msg          string
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

	text := teste.Msg //time.Now().Format(time.RFC3339)
	token := client.Publish("qualidadeAr", 1, false, text)
	token.Wait()
	fmt.Println("Publicado:", text)

	select {
	case <-time.After(time.Second * time.Duration(teste.TesteDuracao)):
		fmt.Println("Cliente desconectado.")
		client.Disconnect(250)

	}
}
