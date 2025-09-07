package v1

import (
	"github.com/Kartikk1127/search-summarizer/internal/server"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/summarize", server.SetupRouter)
}
