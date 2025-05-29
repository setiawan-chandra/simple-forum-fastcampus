package main

import (
	"docker-compose-mysql/internal/handlers/memberships"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	membershipsHandler := memberships.NewHandler(r)
	membershipsHandler.RegisterRoute()
	r.Run()
}
