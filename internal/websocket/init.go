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

type ResumeConnection struct {
	TaskType              string `json:"taskType"`
	ApiKey                string `json:"apiKey"`
	ConnectionSessionUUID string `json:"connectionSessionUUID"`
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

	var conn *websocket.Conn
	var connSession string

	resumeConnection := make(chan struct{}, 1)

	// Create new goroutine for the connection and authentication
	go newConnection(&conn, &connSession, resumeConnection)

	resumeConnection <- struct{}{}

	// Start listening goroutine
	go listenWebsocket(&conn, &connSession, resumeConnection, responseCh)

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	tickerPayload := []Ping{
		{
			TaskType: "ping",
			Ping:     true,
		},
	}

	for {

		currentLocalConn := conn

		select {

		case request := <-requestCh:

			if currentLocalConn == nil {
				time.Sleep(500 * time.Millisecond)
				continue
			}
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

// This function will create new connection with websocket
func newConnection(connPtr **websocket.Conn, connSessionPtr *string, resumeConnChan chan struct{}) {
	apiUrl := config.Config("WEBSOCKET_API_URL")
	apiKey := config.Config("CREATIVEDREAM_WEBSOCKET_API_KEY")

	for {
		<-resumeConnChan
		log.Println("Establish new connection:")

		if *connPtr != nil {
			(*connPtr).Close()
			*connPtr = nil
		}

		// Establish new connection
		dialedConn, resp, err := websocket.DefaultDialer.Dial(apiUrl, nil)
		if err != nil {
			log.Printf("New connection dialing: %v", err)
			*connPtr = nil
			time.Sleep(5 * time.Second)
			continue
		}
		log.Printf("New connection dialing: %s", resp.Status)

		*connPtr = dialedConn
		var payload any

		// Send Authentication message
		if *connSessionPtr == "" {
			payload = []Authentication{
				{
					TaskType: "authentication",
					ApiKey:   apiKey,
				},
			}
		} else {
			payload = []ResumeConnection{
				{
					TaskType:              "authentication",
					ApiKey:                apiKey,
					ConnectionSessionUUID: *connSessionPtr,
				},
			}
		}

		if err := (*connPtr).WriteJSON(payload); err != nil {
			log.Printf("Authentication wrong: %v", err)
		}
		log.Println("Successful active connection!")
	}
}

func listenWebsocket(connPtr **websocket.Conn, connSessionPtr *string, resumeConn chan struct{}, responseCh chan models.GeneratedImage) {

	log.Println("Websocket listener started")

	for {
		currentLocalConn := *connPtr
		if currentLocalConn == nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		_, message, err := currentLocalConn.ReadMessage()
		if err != nil {
			log.Printf("Reading from the websocket: %v", err)
			select {
			case resumeConn <- struct{}{}:
				log.Println("Sent resume signal.")
			default:
				log.Println("resumeConnChan full or new connection not ready")
			}
			time.Sleep(200 * time.Millisecond)
			continue
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
				if data.ConnectionSessionUUID != "" {
					log.Printf("Authentication successful: %s", data.ConnectionSessionUUID)
					*connSessionPtr = data.ConnectionSessionUUID
				}
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
