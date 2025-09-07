package main

import (
	v1 "github.com/Kartikk1127/search-summarizer/internal/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1Group := r.Group("/api/v1")
	v1.RegisterRoutes(v1Group)

	r.Run(":8080")
}
