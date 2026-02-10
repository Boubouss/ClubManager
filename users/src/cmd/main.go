package main

import (
	"flag"
	"fmt"
)


func main() {
  port := flag.String("port", ":3000", "gRPC port for user service")
  flag.Parse()

  fmt.Println(fmt.Printf("Listening on port %s", *port))

  // svc := service.NewUserService(db *pgx.Conn)
}
