package main

import (
	"api/src/config"
	"api/src/router"
	"crypto/rand"
	base642 "encoding/base64"
	"fmt"
	"log"
	"net/http"
)

func init() {
	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	base64 := base642.StdEncoding.EncodeToString(key)

	fmt.Println(base64)
}

func main() {
	config.Load()
	r := router.Generate()

	fmt.Printf("Server Listening on port %d", config.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}
