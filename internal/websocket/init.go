package websocket

import (
	"encoding/json"
	"log"
	"time"

	"github.com/dzianis-sudkou/go-telegram-bot/internal/config"
	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
	"github.com/gorilla/websocket"
)

type RequestPayload struct {
	ApiKey         string  `json:"apiKey"`
	TaskType       string  `json:"taskType"`
	TaskUUID       string  `json:"taskUUID"`
	OutputType     string  `json:"outputType"`
	Width          int     `json:"width"`
	Height         int     `json:"height"`
	OutputFormat   string  `json:"outputFormat"`
	PositivePrompt string  `json:"positivePrompt"`
	NegativePrompt string  `json:"negativePrompt"`
	Model          string  `json:"model"`
	Steps          int     `json:"steps"`
	CFGScale       float64 `json:"CFGScale"`
	Scheduler      string  `json:"scheduler"`
	NumberResults  int     `json:"numberResults"`
	Ping           bool    `json:"ping"`
	IncludeCost    bool    `json:"includeCost"`
	CheckNSFW      bool    `json:"checkNSFW"`
	ClipSkip       int     `json:"clipSkip"`
}

type ResponsePayload struct {
	TaskType              string  `json:"taskType"`
	TaskUUID              string  `json:"taskUUID"`
	ImageUUID             string  `json:"imageUUID"`
	ImageURL              string  `json:"imageURL"`
	ImageBase64Data       string  `json:"imageBase64Data"`
	ImageDataURI          string  `json:"imageDataURI"`
	Seed                  uint64  `json:"seed"`
	NSFWContent           bool    `json:"NSFWContent"`
	Cost                  float64 `json:"cost"`
	Pong                  bool    `json:"pong"`
	ConnectionSessionUUID string  `json:"connectionSessionUUID"`
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
	authPayload := []RequestPayload{
		{
			ApiKey:   config.Config("CREATIVEDREAM_WEBSOCKET_API_KEY"),
			TaskType: "authentication",
		},
	}

	if err := conn.WriteJSON(authPayload); err != nil {
		log.Printf("Authentication wrong: %v", err)
	}

	// Start listening goroutine
	go listenWebsocket(conn, responseCh)

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	tickerPayload := []RequestPayload{
		{
			TaskType: "ping",
			Ping:     true,
		},
	}

	for {
		select {
		case request := <-requestCh:
			log.Printf("Processing generation request for TaskUUID: %s", request.TaskUUID)

			// Create request payload
			payload := []RequestPayload{
				{
					TaskType:      "imageInference",
					TaskUUID:      request.TaskUUID,
					OutputType:    "URL",
					OutputFormat:  "JPG",
					Width:         640,
					Height:        1152,
					CFGScale:      3.5,
					NumberResults: 1,
					IncludeCost:   true,
					CheckNSFW:     true,
				},
			}
			switch request.Model {
			case "CreativeDream": // TODO
			case "realism":
				payload[0].ClipSkip = 0
				payload[0].Scheduler = "Default"
				payload[0].PositivePrompt = request.Prompt
				payload[0].NegativePrompt = "score_6, score_5, score_4, text, censored, deformed, bad hand, watermark"
				payload[0].Model = "civitai:372465@914390"
				payload[0].Steps = 28
			case "anime":
				payload[0].ClipSkip = 2
				payload[0].Scheduler = "DPM++ 2M"
				payload[0].PositivePrompt = "score_9, score_8_up, score_7_up, score_6_up, score_5_up, score_4_up, source_anime, " + request.Prompt
				payload[0].NegativePrompt = `3D, CGI, render, photo, text, watermark, low-quality, signature, moiré pattern, downsampling, aliasing,
							distorted, blurry, glossy, blur, jpeg artifacts, compression artifacts, poorly drawn, low-resolution, bad, distortion,
							twisted, excessive, exaggerated pose, exaggerated limbs, grainy, symmetrical, duplicate, error, pattern, beginner, pixelated,
							fake, hyper, glitch, overexposed, high-contrast, bad-contrast`
				payload[0].Model = "civitai:439889@828380"
				payload[0].Steps = 25
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
				log.Printf("Pong: %t", data.Pong)
			case "authentication":
				log.Printf("Authentication successful: %s", data.ConnectionSessionUUID)
			case "imageInference":
				log.Printf("Image generation completed - TaskUUID: %s, URL: %s", data.TaskUUID, data.ImageURL)

				// Send the result to the response channel
				generatedImage := models.GeneratedImage{
					NSFW:     data.NSFWContent,
					TaskUUID: data.TaskUUID,
					ImageURL: data.ImageURL,
					Done:     true,
				}
				responseCh <- generatedImage
			}
		}

	}
}
