package main

import (
	"database/sql"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	host     = "10.254.17.176"
	port     = 5432
	user     = "postgres"
	password = "senha"
	dbname   = "postgres"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

type Teste struct {
	tested bool
	sec    int
}

func publiStart(sec int, msg string) {

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

	token := client.Publish("test/topic", 1, false, msg)
	token.Wait()
	fmt.Println("Publicado:", msg)
	time.Sleep(2 * time.Second)
	select {
	case <-time.After(time.Second * time.Duration(sec)):
		fmt.Println("Cliente desconectado.")
		client.Disconnect(250)

	}
}

type SensorData struct {
	value    float32
	typeName string
}

func db(sensor_ SensorData) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// insert_row()
	sqlStatement := `INSERT INTO sensor(Dados, Timestamp, Tipo) 
	VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, sensor_.value, time.Now().Format(time.TimeOnly), sensor_.typeName)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nRow inserted successfully!")
	}
}

type Sensor struct {
	nameType        string
	delta           float32
	min             float32
	max             float32
	connection      bool
	counter         int
	status          string
	startWave       bool
	currentPosition float32
	goPosition      float32
	waveSize        int
}

func NewSensor(name string, min float32, max float32) *Sensor {
	return &Sensor{
		nameType:        name,
		min:             min,
		max:             max,
		connection:      false,
		counter:         10,
		status:          "new wave",
		startWave:       true,
		currentPosition: 0,
		goPosition:      0,
		waveSize:        0,
	}
}

func (s *Sensor) On(test Teste) {
	// Connect to the broker
	s.connection = true
	s.counter = test.sec
	for s.connected(test) {
		message := s.generateWaveData()
		db(SensorData{value: message, typeName: "temperatura"})
		strMessage := fmt.Sprintf("%.2f", message)
		publiStart(test.sec, strMessage)

	}
}

func (s *Sensor) Off() {
	s.connection = false
	fmt.Println("Publicação encerrada")
}

func (s *Sensor) connected(test Teste) bool {
	if test.tested {
		s.counter -= 1
		if s.counter < 0 {
			s.Off()
		}
	}
	return s.connection
}

func (s *Sensor) generateWaveData() float32 {
	var value float32 = s.goPosition

	if s.status == "new wave" {
		s.waveSize = 5
		if s.startWave {
			s.currentPosition = float32(rand.Intn(int(s.max-s.min)) + int(s.min))
			s.startWave = false
			value = s.currentPosition
		}
		s.status = "wave"
	} else if s.status == "wave" {
		s.waveSize -= 1
		value = s.currentPosition + float32(math.Round(math.Abs(rand.NormFloat64()*5)*math.Cos(time.Hour.Minutes())*10)/10)
		s.delta += 1
		if s.waveSize <= 0 {
			s.status = "transtion"
			s.goPosition = float32(rand.Intn(int(s.max-s.min)) + int(s.min))
		}
	} else if s.status == "transtion" {
		fmt.Printf("Indo de %.2f para %.2f\n", s.currentPosition, s.goPosition)
		posOrNeg := 0.0
		difference := s.currentPosition - s.goPosition

		if difference > 0 {
			posOrNeg = 1.0
		} else {
			posOrNeg = -1.0
		}

		fmt.Println(difference, posOrNeg)
		s.currentPosition -= float32(5.0 * posOrNeg)
		value = s.currentPosition
		if s.currentPosition == s.goPosition {
			s.currentPosition = s.goPosition
			s.status = "new wave"
		}
	} else {
		fmt.Println("Erro")
	}

	if value > s.max {
		value = s.max
	}
	if value < s.min {
		value = s.min
	}

	return value
}

func main() {
	test := Teste{tested: false, sec: 3}
	s := NewSensor("sensor", -100, 100)
	s.On(test)
}
