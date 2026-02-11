package main

import (
	"ClubManager/users/internal/middlewares"
	"ClubManager/users/internal/server"
	"ClubManager/users/internal/service"
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)


func main() {
  port := flag.String("port", ":3000", "gRPC port for user service")
  flag.Parse()

  ctx := context.Background()

  db, err := pgx.Connect(ctx, os.Getenv("DB_URL"))
  if err != nil {
    fmt.Println("Connection to database failed")
    fmt.Println(os.Getenv("DB_URL"))
    return
  }

  svc := middlewares.NewLoggingService(service.NewUserService(db))

  server.MakeServerAndRun(*port, svc)
}
