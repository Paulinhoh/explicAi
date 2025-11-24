package configuration

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Aplication struct {
	server *echo.Echo
}

func NewAplication() *Aplication {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	return &Aplication{
		server: server,
	}
}

func (a *Aplication) Start() {
	fmt.Println("explicAi is starting on 0.0.0.0:8080...")
	a.server.Start("0.0.0.0:8080")
}
