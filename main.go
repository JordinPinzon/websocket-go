package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Definimos el upgrader de WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Manejamos las conexiones WebSocket
func handleConnection(w http.ResponseWriter, r *http.Request) {
	// Actualizamos el protocolo HTTP a WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close() // Cerramos la conexión al final

	// Bucle para recibir y responder mensajes
	for {
		// Leer el mensaje enviado por el cliente
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error al leer el mensaje:", err)
			return
		}

		// Responder al cliente (Echo)
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("Error al enviar el mensaje:", err)
			return
		}
	}
}

func main() {
	// Definimos la ruta para manejar conexiones WebSocket
	http.HandleFunc("/ws", handleConnection)

	// Arrancamos el servidor en el puerto 8080
	fmt.Println("Servidor WebSocket corriendo en ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
