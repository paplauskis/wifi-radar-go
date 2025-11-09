package router

import (
	"database/sql"
	_map "wifi-radar-go/internal/map"
	_wifi "wifi-radar-go/internal/wifi"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	r := gin.Default()

	mapHandler := _map.NewHandler(db)
	//userHandler := user.NewHandler(db)
	wifiHandler := _wifi.NewHandler(db)

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

		wifiRoutes := api.Group("/wifi")
		{
			wifiRoutes.GET("/reviews", wifiHandler.GetReviews)
			wifiRoutes.POST("/reviews", wifiHandler.CreateReview)
			wifiRoutes.GET("/passwords", wifiHandler.GetPasswords)
			wifiRoutes.POST("/passwords", wifiHandler.AddPassword)
		}
	}

	return r
}
