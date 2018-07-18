package main

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

func main() {
	hashKey := securecookie.GenerateRandomKey(64)
	blockKey := securecookie.GenerateRandomKey(32)

	s := securecookie.New(hashKey, blockKey) // the blockKey is used to encrypt the cookie's value

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if encoded, err := s.Encode("mycookie", "myvalue"); err == nil {
			cookie := http.Cookie{
				Name:  "mycookie",
				Value: encoded,
				Path:  "/",
			}
			// playing with decode
			var result string
			if err = s.Decode("mycookie", cookie.Value, &result); err == nil {
				println(result)
			}
			http.SetCookie(w, &cookie)
		}
		w.Write([]byte("Setting cookies"))
	})

	http.ListenAndServe(":3000", nil)
}
