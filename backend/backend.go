package backend

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	e.Add("POST", "/group", PostGroup)
	e.Add("POST", "/activities", PostActivities)
	e.Add("POST", "/chat", PostChat)
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

func StartServer(getGroup func() []*GroupInfo) {
	GetGroupInfoList = getGroup
	e := newEchoServer()
	s := &http.Server{
		Addr:         ":3003",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	go func() {
		if err := e.StartServer(s); err != nil {
			fmt.Printf("shutting down the server: %v", err)
			panic(err)
		}
	}()

	// <-ctx.Done()
	// if err := e.Shutdown(ctx); err != nil {
	// 	e.Logger.Fatal(err)
	// }
}
