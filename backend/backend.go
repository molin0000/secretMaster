package backend

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/molin0000/secretMaster/qlog"
)

func loadRoutes(e *echo.Echo) {
	//---Get------------
	e.Add("GET", "/", GetInterfaceList)
	e.Add("GET", "/version", GetVersion)
	e.Add("GET", "/count", GetCount)
	e.Add("GET", "/supermaster", GetSuperMaster)
	e.Add("GET", "/delay", GetDelay)
	e.Add("GET", "/imageMode", GetImageMode)
	e.Add("GET", "/textSegment", GetTextSegment)
	e.Add("GET", "/moneyMap", GetMoneyMap)
	e.Add("GET", "/group", GetGroup)
	e.Add("GET", "/activities", GetActivities)
	e.Add("GET", "/locked", GetLocked)

	//----Post--------------------------
	e.Add("POST", "/password", PostPassword)
	e.Add("POST", "/supermaster", PostSuperMaster)
	e.Add("POST", "/delay", PostDelay)
	e.Add("POST", "/imageMode", PostImageMode)
	e.Add("POST", "/textSegment", PostTextSegment)
	e.Add("POST", "/moneyMap", PostMoneyMap)
	e.Add("POST", "/activities", PostActivities)
	e.Add("POST", "/chat", PostChat)

	e.Add("POST", "/globalSwitch", PostGlobalSwitch)
	e.Add("POST", "/globalSilent", PostGlobalSilent)
	e.Add("POST", "/groupSwitch", PostGroupSwitch)
	e.Add("POST", "/groupSilent", PostGroupSilent)
	// e.Add("POST", "/groupExit", PostGroupExit)
}

func newEchoServer() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	loadRoutes(e)

	return e
}

var echoServer *echo.Echo
var httpServer *http.Server

func recoverFunc() {
	if err := recover(); err != nil {
		qlog.Println(err) // 这里的err其实就是panic传入的内容
	}
}

func StartServer(getGroup func() []*GroupInfo) {
	defer recoverFunc()
	qlog.Println("后台服务启动...")
	GetGroupInfoList = getGroup
	e := newEchoServer()
	s := &http.Server{
		Addr:         ":3003",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	echoServer = e
	httpServer = s

	go func() {
		defer recoverFunc()
		if err := e.StartServer(s); err != nil {
			qlog.Printf("StartServer error: %v", err)
		}
	}()

	qlog.Println("后台服务启动完成...")
}

func StopServer() {
	defer recoverFunc()
	qlog.Println("后台服务准备停止")
	ctx, stop := context.WithCancel(context.Background())
	stop()
	echoServer.Shutdown(ctx)
	httpServer.Shutdown(ctx)
	// if err := echoServer.Shutdown(ctx); err != nil {
	// 	echoServer.Logger.Fatal(err)
	// }
	// httpServer.Close()
	// echoServer.Listener.Close()
	// echoServer.Server.Close()
	// echoServer.Close()
	qlog.Println("后台服务停止完毕")
}
