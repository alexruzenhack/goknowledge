package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookie := http.Cookie{
			Name:  "mycookie",
			Value: "myvalue",
			Path:  "/",
		}

		http.SetCookie(w, &cookie)

		w.Write([]byte("Setting cookies"))
	})

	http.ListenAndServe(":3000", nil)
}
