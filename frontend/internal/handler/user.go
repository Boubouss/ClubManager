package handler

import (
	"ClubManager/frontend/internal/views/components"
	"ClubManager/frontend/internal/views/pages"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {}

func NewUserHandler() UserHandler {
  return UserHandler{}
}

func (h *UserHandler) HandleConnexion(c *echo.Context) error {
  return render(c, pages.Connexion())
}

func (h *UserHandler) HandleRegisterForm(c *echo.Context) error {
  return render(c, components.RegisterForm())
}

func (h *UserHandler) HandleRegisterUser(c *echo.Context) error {
  return render(c, pages.Home())
}

func (h *UserHandler) HandleLoginForm(c *echo.Context) error {
  return render(c, components.LoginForm())
}

func (h *UserHandler) HandleLoginUser(c *echo.Context) error {
  return render(c, pages.Home())
}
