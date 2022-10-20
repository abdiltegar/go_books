package haloController

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type haloController struct {
}

func NewHaloController() *haloController {
	return &haloController{}
}

// get halo
func (ctrl *haloController) GetHalo(c echo.Context) error {
	return c.HTML(http.StatusOK, "Ini adalah Halo Page (Updated)")
}
