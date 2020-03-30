package ui

import (
	"fmt"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

func MakeDoc() {
	http.Handle("/doc", http.FileServer(rice.MustFindBox("../doc").HTTPBox()))
	fmt.Println("8080 ListenAndServe...")
	http.ListenAndServe(":8080", nil)
}
