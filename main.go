package main

import (
	"github.com/dribeiroferr/referral-system-go/src/configs"
	"github.com/dribeiroferr/referral-system-go/src/handlers"
	"github.com/dribeiroferr/referral-system-go/src/repositories"
	"github.com/dribeiroferr/referral-system-go/src/routes"
	"github.com/dribeiroferr/referral-system-go/src/services"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := configs.InitDB()
	defer db.Close()

	referralRepo := repositories.NewReferralRepository(db)
	referralService := services.NewReferralService(referralRepo)
	referralHandler := handlers.NewReferralHandler(referralService)

	routes.InitRoute(e, referralHandler)

	e.Logger.Fatal(e.Start(":9001"))
}