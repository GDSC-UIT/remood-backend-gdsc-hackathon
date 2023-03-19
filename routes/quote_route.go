package routes 

import (
	"remood/handlers"

	"github.com/gin-gonic/gin"
)

func QuoteRouter(r *gin.RouterGroup) {
	quoteRouter := r.Group("quotes")
	{
		quoteRouter.POST("/", handlers.CreateManyQuotes)
		quoteRouter.GET("/", handlers.GetRandomQuotes)
	}
}

