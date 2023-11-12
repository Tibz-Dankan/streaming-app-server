package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")

		allowedOrigins := []string{
			"http://localhost:5173",
			"https://owino-dev.netlify.app",
			"http://http://127.0.0.1:5500/index.html",
		}

		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				return true // Allow the connection for this origin
			}
		}

		// Deny the connection for any other origins
		return false
	},
}

func WSWebRTCHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// type Message struct {
	// 	Type         string              `json:"type"`
	// 	Offer        models.Offer        `json:"offer"`
	// 	IceCandidate models.ICECandidate `json:"iceCandidate"`
	// }

	// // type ClientOffer struct {
	// // 	Offer models.Offer `json:"offer"`
	// // }
	// type ServerAnswer struct {
	// 	Answer models.Answer `json:"answer"`
	// }

	// type ServerIceCandidate struct {
	// 	IceCandidate models.ICECandidate `json:"iceCandidate"`
	// }

	type Message struct {
		Type         string                    `json:"type"`
		Offer        webrtc.SessionDescription `json:"offer"`
		Answer       webrtc.SessionDescription `json:"answer"`
		IceCandidate webrtc.ICECandidate       `json:"iceCandidate"`
	}

	type ServerAnswer struct {
		Answer webrtc.SessionDescription `json:"answer"`
	}

	type ServerIceCandidate struct {
		IceCandidate webrtc.ICECandidate `json:"iceCandidate"`
	}

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Received message from client: \n", message)
		fmt.Println("Received message from client (raw JSON): \n", string(message))

		var receivedMessage Message
		if err := json.Unmarshal(message, &receivedMessage); err != nil {
			log.Println("Failed to unmarshal JSON of message: \n", err)
			continue
		}
		fmt.Println("receivedMessage ", receivedMessage)

		// Check for the "offer" property
		if receivedMessage.Type == "offer" {
			fmt.Println("Received client offer: \n", receivedMessage.Offer)
			// receive offer send answer
			answer := ReceiveOfferCreateAnswer(receivedMessage.Offer)
			message := Message{Type: "answer", Answer: answer}

			jsonMessage, err := json.Marshal(message)
			if err != nil {
				fmt.Println(err)
				return
			}
			if err := conn.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
				fmt.Println(err)
				return
			}
		}

		// Check for the "iceCandidate" property
		if receivedMessage.Type == "iceCandidate" {
			fmt.Println("Received client iceCandidate: \n", receivedMessage.IceCandidate)
			// receive client iceCandidate send send server iceCandidate

			iceCandidate := ReceiveCreateIceCandidate(receivedMessage.IceCandidate)
			message := Message{Type: "iceCandidate", IceCandidate: iceCandidate}

			jsonMessage, err := json.Marshal(message)
			if err != nil {
				fmt.Println(err)
				return
			}
			if err := conn.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
				fmt.Println(err)
				return
			}
		}

		fmt.Println("messageType from client: ", messageType)
		// Send back message to the client
		// err = conn.WriteMessage(messageType, message)
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }
	}
}
