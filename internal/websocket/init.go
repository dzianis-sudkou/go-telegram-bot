package websocket

// import (
// 	"encoding/json"
// 	"log"
// 	"time"

// 	"github.com/dzianis-sudkou/go-telegram-bot/internal/config"
// 	"github.com/dzianis-sudkou/go-telegram-bot/internal/models"
// 	"github.com/gorilla/websocket"
// )

// type RequestPayload struct {
// 	ApiKey         string  `json:"apiKey"`
// 	TaskType       string  `json:"taskType"`
// 	TaskUUID       string  `json:"taskUUID"`
// 	OutputType     string  `json:"outputType"`
// 	Width          int     `json:"width"`
// 	Height         int     `json:"height"`
// 	OutputFormat   string  `json:"outputFormat"`
// 	PositivePrompt string  `json:"positivePrompt"`
// 	NegativePrompt string  `json:"negativePrompt"`
// 	Model          string  `json:"model"`
// 	Steps          int     `json:"steps"`
// 	CFGScale       float64 `json:"CFGScale"`
// 	NumberResults  int     `json:"numberResults"`
// 	Ping           bool    `json:"ping"`
// 	IncludeCost    bool    `json:"includeCost"`
// }

// type ResponsePayload struct {
// 	TaskType              string  `json:"taskType"`
// 	TaskUUID              string  `json:"taskUUID"`
// 	ImageUUID             string  `json:"imageUUID"`
// 	ImageURL              string  `json:"imageURL"`
// 	ImageBase64Data       string  `json:"imageBase64Data"`
// 	ImageDataURI          string  `json:"imageDataURI"`
// 	Seed                  uint64  `json:"seed"`
// 	NSFWContent           bool    `json:"NSFWContent"`
// 	Cost                  float64 `json:"cost"`
// 	Pong                  bool    `json:"pong"`
// 	ConnectionSessionUUID string  `json:"connectionSessionUUID"`
// }

// // The Response we receive from the server
// type OutWebsocket struct {
// 	Data []ResponsePayload `json:"data"`
// }

// func Init(chIn chan models.FreeRequest, chOut chan models.GeneratedImage) {

// 	// Establish the connection with websocket
// 	conn, resp, err := websocket.DefaultDialer.Dial(config.Config("WEBSOCKET_API_URL"))
// 	if err != nil {
// 		log.Printf("Websocket connection: %v", err)
// 	}
// 	log.Printf("Successfully connected to WebSocket: %s", resp.Status)

// 	defer conn.Close() // Close the connection

// 	// Start the goroutine listen to WebSocket
// 	go listenWebsocket(conn)

// 	go writeWebsoket(conn)
// 	for {
// 		select {
// 		case prompt := <-chIn:
// 		case <-chOut:
// 		}
// 	}
// }

// func listenWebsocket(conn *websocket.Conn) {
// 	// defer close()
// 	time.Sleep(1 * time.Second)
// 	defer log.Println("Read goroutine Finished")
// 	log.Println("Listen websocket goroutine started.")
// 	for {
// 		_, message, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Printf("Reading from the websocket: %v", err)
// 		}

// 		// Unmarshal the Response
// 		var receivedPayload OutWebsocket
// 		if err = json.Unmarshal(message, &receivedPayload); err != nil {
// 			log.Printf("Unmarshal websocket response: %v", err)
// 		}
// 		for _, val := range receivedPayload.Data {
// 			switch val.TaskType {
// 			case "ping":
// 				log.Printf("Pong: %t", val.Pong)
// 			case "authentication":
// 				log.Printf("Authentication successfull: %s", val.ConnectionSessionUUID)
// 			case "imageInference":
// 				log.Printf("Task: %s\nImageURL: %s Cost: %.4f", val.TaskUUID, val.ImageURL, val.Cost)
// 			}
// 		}
// 	}
// }

// func writeWebsoket(conn *websocket.Conn) {
// 	log.Println("Write websocket goroutine started.")

// }
