package main

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

func main() {
	hashKey := securecookie.GenerateRandomKey(64)

	s := securecookie.New(hashKey, nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if encoded, err := s.Encode("mycookie", "myvalue"); err == nil {
			cookie := http.Cookie{
				Name:  "mycookie",
				Value: encoded,
				Path:  "/",
			}
			http.SetCookie(w, &cookie)
		}
		w.Write([]byte("Setting cookies"))
	})

	http.ListenAndServe(":3000", nil)
}
