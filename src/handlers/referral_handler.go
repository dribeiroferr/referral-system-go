package handlers

import (
	"net/http"

	"github.com/dribeiroferr/referral-system-go/src/services"
	"github.com/labstack/echo/v4"
)

type ReferralHandler struct {
	service services.ReferralService
}

func NewReferralHandler(service services.ReferralService) *ReferralHandler {
	return &ReferralHandler{service}
}

func (h *ReferralHandler) GenerateLink(c echo.Context) error {
	user := c.FormValue("user")
	link, _, promoCode, err := h.service.GenerateReferralLink(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"referral_link": link, "promo_code": promoCode})
}

func (h *ReferralHandler) HandleReferral(c echo.Context) error {
	hash := c.Param("hash")
	referral, err := h.service.HandleReferral(hash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":           referral.User,
		"promo_code":     referral.PromoCode,
		"count":          referral.Count,
		"date_generated": referral.DateGenerated,
		"date_registered": referral.DateRegistered,
	})
}


func (h *ReferralHandler) ShowLandingPage(c echo.Context) error{
	return c.File("public/index.html")
}