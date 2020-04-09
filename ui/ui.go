package ui

import (
	"fmt"
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

func StartUI() {
	fmt.Println("图形界面启动...")
	http.Handle("/", http.FileServer(rice.MustFindBox("./webpage/dist").HTTPBox()))
	fmt.Println("5000 ListenAndServe...")
	go func() {
		http.ListenAndServe(":5000", nil)
	}()
}
