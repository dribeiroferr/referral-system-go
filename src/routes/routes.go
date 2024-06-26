package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/dribeiroferr/referral-system-go/src/handlers"
)

func InitRoute(e *echo.Echo, referralHandler *handlers.ReferralHandler){
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/generate", referralHandler.GenerateLink)
	e.GET("/referral/:hash", referralHandler.HandleReferral)
	e.GET("/landing", referralHandler.ShowLandingPage)
}
