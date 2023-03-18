package main

import (
	"context"
	"fmt"
	"os"

	"remood/pkg/auth"	
	"remood/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"remood/pkg/database"
	"remood/pkg/middleware"
)

func main() {
	//DATABASE CONNECTION
	db := database.GetMongoInstance()
	defer db.Client.Connect(context.Background())
	fmt.Println("MONGODB CONNECTED")
 
	// GENERATE JWT SECRET KEY
	auth.GenerateJWTKey()

	//GIN DEFINE
	router := gin.Default()
	router.Use(middleware.CorsMiddleware)
	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Ping successful",
			})
		})
	}

	//ROUTER DEFINE
	routes.UserRouter(api)
	routes.DiaryNoteRouter(api)
	routes.ReviewNoteRouter(api)
	routes.DayReviewRouter(api)
	routes.ArticleRouter(api)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if port == "" {
        	port = "8080"
   	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))	
	router.Run("localhost" + port)
}
