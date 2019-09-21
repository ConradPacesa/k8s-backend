package handlers

import (
	"io"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200")
}
