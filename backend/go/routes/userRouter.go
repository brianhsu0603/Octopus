package routes

import (
	"octopus/middleware"

	"octopus/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	//incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.PUT("/users/:user_id", controllers.UpdateUser())
	incomingRoutes.DELETE("/users/:user_id", controllers.DeleteUser())

	incomingRoutes.POST("/logout", controllers.Logout())
	incomingRoutes.POST("/search", controllers.Search())

	incomingRoutes.GET("/watchlist/:user_id", controllers.GetWatchlist())
	incomingRoutes.POST("/watchlist/:user_id", controllers.AddItem())
	incomingRoutes.DELETE("/watchlist/:user_id/*item_id", controllers.DeleteItem())

}
