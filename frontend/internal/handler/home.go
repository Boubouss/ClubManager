package handler

import (
	"ClubManager/frontend/internal/views/pages"

	"github.com/labstack/echo/v5"
)

type HomeHandler struct {}

func NewHomeHandler() HomeHandler {
  return HomeHandler{}
}

func (h *HomeHandler) HandleLandingPage(c *echo.Context) error {
  return render(c, pages.Landing())
}

func (h *HomeHandler) HandleHomePage(c *echo.Context) error {
  return render(c, pages.Home())
}
