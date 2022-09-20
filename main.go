package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/matiasbenary/chatActivity/database"
	"github.com/matiasbenary/chatActivity/repository"
)

var ctx = context.Background()

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

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":"+addr, nil))
}
