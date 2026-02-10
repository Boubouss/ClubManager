package main

import (
	"ClubManager/frontend/internal/handler"

	"github.com/labstack/echo/v5"
)

func main() {
  app := echo.New()

  homeHandler := handler.NewHomeHandler()
  userHandler := handler.NewUserHandler()

  app.GET("/", homeHandler.HandleLandingPage)
  app.GET("/home", homeHandler.HandleHomePage)

  auth := app.Group("/user")

  auth.GET("/connexion", userHandler.HandleConnexion)
  auth.GET("/register", userHandler.HandleRegisterForm)
  auth.GET("/login", userHandler.HandleLoginForm)
  auth.POST("/login", userHandler.HandleLoginUser)
  auth.POST("/register", userHandler.HandleRegisterUser)
  
  app.Start(":3000")
}
