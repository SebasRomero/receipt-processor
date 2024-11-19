package receipt

import (
	"net/http"
)

func Process(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Process"))
}

func Points(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Points"))
}
