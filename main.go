package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	model "github.com/DhanakotiS/go-authentication-system/models"
)

var (
	Infolog  *log.Logger
	Warnlog  *log.Logger
	Errorlog *log.Logger
)

var (
	port = flag.String("port", "3032", "Port to which to listen for")
	host = flag.String("host", "http://localhost", "Host name of the url")
)

func init() {
	file, err := os.OpenFile(".logs/debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	Infolog = log.New(file, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)
	Warnlog = log.New(file, "WARNING : ", log.Ldate|log.Ltime|log.Lshortfile)
	Errorlog = log.New(file, "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	flag.Parse()

	kd := model.User{
		ID:         1,
		Name:       "Name",
		Email:      "mail@test.com",
		UserName:   "username",
		Password:   base64.StdEncoding.EncodeToString([]byte("<Password>")),
		TokenHash:  fmt.Sprintf("%x", sha256.Sum256([]byte("<Token goes here>"))),
		IsVerified: true,
		CreatedAt:  time.Now(),
	}
	Infolog.Println("Server Configured with port ", *port)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(kd.TokenHash)
	})

	Infolog.Println("Starting Server, Listening to port " + *host + ":" + *port)

	if err := http.ListenAndServe(":"+*port, r); err != nil {
		Errorlog.Println("Error listening to Port", *port, err)
		os.Exit(1)
	}

}
