package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<H1>Hello Test!"))
	})
	http.ListenAndServe(":80", nil)
}
