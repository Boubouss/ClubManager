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

  user := app.Group("/user")

  user.GET("/connexion", userHandler.HandleConnexion)
  user.GET("/register", userHandler.HandleRegisterForm)
  user.GET("/login", userHandler.HandleLoginForm)
  user.POST("/login", userHandler.HandleLoginUser)
  user.POST("/register", userHandler.HandleRegisterUser)
  
  app.Start(":3000")
}
