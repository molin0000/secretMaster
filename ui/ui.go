package ui

import (
	"fmt"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

func MakeDoc() {
	http.Handle("/", http.FileServer(rice.MustFindBox("./webpage/dist").HTTPBox()))
	fmt.Println("8080 ListenAndServe...")
	http.ListenAndServe(":8080", nil)
}
