package router

import (
	"database/sql"
	_map "wifi-radar-go/internal/map"
	"wifi-radar-go/internal/middleware"
	"wifi-radar-go/internal/user"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	r := gin.Default()

	mapHandler := _map.NewHandler(db)
	userHandler := user.NewUserHandler(db)
	wifiHandler := wifi.NewHandler(db)

	api := r.Group("/api")
	{
		// Map routes
		mapRoutes := api.Group("/map")
		{
			mapRoutes.GET("/search", mapHandler.Search)
			mapRoutes.GET("/coordinates", mapHandler.GetCoordinates)
		}

		// User routes
		userRoutes := api.Group("/user")
		{
			userRoutes.POST("/auth/login", userHandler.Login)
			userRoutes.POST("/auth/register", userHandler.Register)

			protected := userRoutes.Group("/")
			protected.Use(middleware.AuthenticationMiddleware())
			{
				protected.POST("/:id/favorite", middleware.OwnData(), userHandler.AddFavorite)
				protected.GET("/:id/favorite", middleware.OwnData(), userHandler.GetFavorite)
				protected.DELETE("/:id/favorite/:wifi_id", middleware.OwnData(), userHandler.DeleteFavorite)

			}
		}

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
