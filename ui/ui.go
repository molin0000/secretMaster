package ui

import (
	"context"
	"net/http"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/molin0000/secretMaster/qlog"
)

var server *http.Server

func recoverFunc() {
	if err := recover(); err != nil {
		qlog.Println(err) // 这里的err其实就是panic传入的内容
	}
}

func StartUI() {
	defer recoverFunc()
	qlog.Println("图形界面启动...")
	handle := http.FileServer(rice.MustFindBox("./webpage/dist").HTTPBox())
	srv := &http.Server{
		Addr:           ":5000",
		Handler:        handle,
		ReadTimeout:    time.Second * 15,
		WriteTimeout:   time.Second * 15,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 1 << 20,
	}

	server = srv
	qlog.Println("5000 ListenAndServe...")
	go func() {
		defer recoverFunc()
		err := srv.ListenAndServe()
		qlog.Println(err)
	}()
}

func StopUI() {
	defer recoverFunc()

	qlog.Println("关闭前端服务")

	ctx, stop := context.WithCancel(context.Background())
	stop()
	server.Shutdown(ctx)
}
