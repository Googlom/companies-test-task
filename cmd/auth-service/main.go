package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func main() {
	hmacSecret := os.Getenv("HMAC_SECRET")
	http.HandleFunc("/login", login([]byte(hmacSecret)))

	listen := ":8090"
	fmt.Printf("server is listening at %s\n", listen)
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func login(hmacSecret []byte) http.HandlerFunc {
	if len(hmacSecret) == 0 {
		log.Fatal("HMAC secret must be set by envvar HMAC_SECRET")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"test": "companies",
			"nbf":  time.Now().Unix(),
		})
		tokenString, err := token.SignedString(hmacSecret)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}
