package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_map "wifi-radar-go/internal/map"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	r := gin.Default()

	mapHandler := _map.NewHandler(db)
	//userHandler := user.NewHandler(db)
	//wifiHandler := wifi.NewHandler(db)

	api := r.Group("/api")
	{
		// Map routes
		mapRoutes := api.Group("/map")
		{
			mapRoutes.GET("/search", mapHandler.Search)
			mapRoutes.GET("/coordinates", mapHandler.GetCoordinates)
		}

		//userRoutes := api.Group("/user")
		//{
		//	userRoutes.POST("/auth/login", userHandler.Login)
		//	userRoutes.POST("/auth/register", userHandler.Register)
		//	userRoutes.POST("/:userId/favorites", userHandler.AddFavorite)
		//	userRoutes.GET("/:userId/favorites", userHandler.GetFavorites)
		//	userRoutes.DELETE("/:userId/favorites/:wifiId", userHandler.DeleteFavorite)
		//}

		//wifiRoutes := api.Group("/wifi")
		//{
		//	wifiRoutes.POST("/reviews", wifiHandler.CreateReview)
		//	wifiRoutes.GET("/reviews", wifiHandler.GetReviews)
		//	wifiRoutes.POST("/passwords", wifiHandler.AddPassword)
		//	wifiRoutes.GET("/passwords", wifiHandler.GetPasswords)
		//}
	}

	return r
}
