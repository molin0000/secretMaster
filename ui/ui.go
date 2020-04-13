package ui

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/molin0000/secretMaster/qlog"
)

func StartUI() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			qlog.Println(err) // 这里的err其实就是panic传入的内容
		}
	}()
	qlog.Println("图形界面启动...")
	http.Handle("/", http.FileServer(rice.MustFindBox("./webpage/dist").HTTPBox()))
	qlog.Println("5000 ListenAndServe...")
	go func() {
		defer func() { // 必须要先声明defer，否则不能捕获到panic异常
			if err := recover(); err != nil {
				qlog.Println(err) // 这里的err其实就是panic传入的内容
			}
		}()
		http.ListenAndServe(":5000", nil)
	}()
}
