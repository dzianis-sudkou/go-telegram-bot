package websocket

import (
	"encoding/json"
	"log"
	"time"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/config"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"github.com/gorilla/websocket"
)

type RequestImage struct {
	TaskType string `json:"taskType"`
	TaskUUID string `json:"taskUUID"`

	OutputType     string  `json:"outputType"`
	OutputFormat   string  `json:"outputFormat"`
	OutputQuality  int     `json:"outputQuality"`
	CheckNSFW      bool    `json:"checkNSFW"`
	IncludeCost    bool    `json:"includeCost"`
	PositivePrompt string  `json:"positivePrompt"`
	NegativePrompt string  `json:"negativePrompt"`
	Model          string  `json:"model"`
	Strength       float64 `json:"strength"`
	Steps          int     `json:"steps"`
	Scheduler      string  `json:"scheduler"`
	CFGScale       float64 `json:"CFGScale"`
	ClipSkip       int     `json:"clipSkip"`
	NumberResults  int     `json:"numberResults"`
	Lora           []Lora  `json:"lora"`
	Width          int     `json:"width"`
	Height         int     `json:"height"`
}

type Lora struct {
	Model  string  `json:"model"`
	Weight float64 `json:"weight"`
}

type Authentication struct {
	TaskType string `json:"taskType"`
	ApiKey   string `json:"apiKey"`
}

type Ping struct {
	TaskType string `json:"taskType"`
	Ping     bool   `json:"ping"`
}

type ResponsePayload struct {
	TaskType string `json:"taskType"`

	// Authentication
	ConnectionSessionUUID string `json:"connectionSessionUUID"`

	// Keeping Connection Alive
	Pong bool `json:"pong"`

	// Image Generation
	TaskUUID        string  `json:"taskUUID"`
	ImageUUID       string  `json:"imageUUID"`
	ImageURL        string  `json:"imageURL"`
	ImageBase64Data string  `json:"imageBase64Data"`
	ImageDataURI    string  `json:"imageDataURI"`
	Seed            uint64  `json:"seed"`
	NSFWContent     bool    `json:"NSFWContent"`
	Cost            float64 `json:"cost"`
}

// The Response we receive from the server
type OutWebsocket struct {
	Data []ResponsePayload `json:"data"`
}

func Init(requestCh chan models.GeneratedImage, responseCh chan models.GeneratedImage) {

	// Establish the connection with websocket
	conn, resp, err := websocket.DefaultDialer.Dial(config.Config("WEBSOCKET_API_URL"), nil)
	if err != nil {
		log.Printf("Websocket connection: %v", err)
		return
	}

	log.Printf("Successfully connected to WebSocket: %s", resp.Status)
	defer conn.Close() // Close connection when function exits

	// Send Authentication message
	authPayload := []Authentication{
		{
			TaskType: "authentication",
			ApiKey:   config.Config("CREATIVEDREAM_WEBSOCKET_API_KEY"),
		},
	}
	if err := conn.WriteJSON(authPayload); err != nil {
		log.Printf("Authentication wrong: %v", err)
	}

	// Start listening goroutine
	go listenWebsocket(conn, responseCh)

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	tickerPayload := []Ping{
		{
			TaskType: "ping",
			Ping:     true,
		},
	}

	for {
		select {
		case request := <-requestCh:

			log.Printf("Processing generation request for TaskUUID: %s", request.TaskUUID)

			model := newRequestImage(&request)

			payload := []RequestImage{
				model,
			}

			if err := conn.WriteJSON(payload); err != nil {
				log.Printf("Sending request: %v", err)
			}

		case <-ticker.C:
			if err := conn.WriteJSON(tickerPayload); err != nil {
				log.Printf("Sending ping: %v", err)
			}
		}
	}
}

func listenWebsocket(conn *websocket.Conn, responseCh chan models.GeneratedImage) {
	log.Println("Websocket listener started")

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Reading from the websocket: %v", err)
		}
		log.Printf("Received data #1: %s", message)

		// Parse the response
		var receivedPayload OutWebsocket
		if err = json.Unmarshal(message, &receivedPayload); err != nil {
			log.Printf("Unmatching websocket response: %v", err)
			continue
		}

		// Process received payloads
		for _, data := range receivedPayload.Data {
			switch data.TaskType {
			case "ping":
			case "authentication":
			case "imageInference":
				// Send the result to the response channel
				generatedImage := models.GeneratedImage{
					NSFW:     data.NSFWContent,
					TaskUUID: data.TaskUUID,
					ImageURL: data.ImageURL,
					Done:     true,
				}
				responseCh <- generatedImage
			default:
				log.Println("Unknown Task Type: ", data.TaskType)
			}
		}
	}
}
