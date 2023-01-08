package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/matiasbenary/chatActivity/database"
	"github.com/matiasbenary/chatActivity/repository"
)

var ctx = context.Background()

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	var addr = os.Getenv("API_PORT")

	database.CreateRedisClient()
	db := database.InitDBMaria()
	defer db.Close()

	wsServer := NewWebsocketServer(&repository.RoomRepository{Db: db}, &repository.UserRepository{Db: db}, &repository.MessageRepository{Db: db})
	go wsServer.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	})

	http.HandleFunc("/moreMessage", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		msjId, ok := r.URL.Query()["msjId"]

		if !ok || len(msjId[0]) < 1 {
			log.Println("Url Param 'msjId' is missing")
			return
		}

		msj := wsServer.messageRepository.MoreMessage(msjId[0])
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msj)
	})

	http.HandleFunc("/lastMessage", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		roomId, ok := r.URL.Query()["roomId"]
		if !ok || len(roomId[0]) < 1 {
			log.Println("Url Param 'roomId' is missing")
			return
		}

		msj := wsServer.messageRepository.LastMessage(roomId[0])
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msj)
	})

	http.HandleFunc("/deleteMessage", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		msjId, ok := r.URL.Query()["msjId"]

		if !ok || len(msjId[0]) < 1 {
			log.Println("Url Param 'msjId' is missing")
			return
		}

		msj := wsServer.messageRepository.DeleteMessage(msjId[0])
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msj)
	})

	http.HandleFunc("/lastActivity", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		msj := wsServer.messageRepository.LastActivity()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msj)
	})
	// http.HandleFunc("/getUser", func(w http.ResponseWriter, r *http.Request) {
	// 	roomId, ok := r.URL.Query()["email"]

	// 	if !ok || len(roomId[0]) < 1 {
	// 		log.Println("Url Param 'email' is missing")
	// 		return
	// 	}
	// 	// userParams := NewUser("tute", roomId[0], "1")
	// 	// user := wsServer.userRepository.AddUser(userParams)
	// 	user := wsServer.userRepository.FindUserByEmail(roomId[0])
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(user)
	// })

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":"+addr, nil))
}
