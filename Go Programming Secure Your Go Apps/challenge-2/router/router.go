package router

import (
	"challenge-2/controllers"
	"challenge-2/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productId", middlewares.UserAuthorization(), controllers.UpdateProduct)
		productRouter.GET("/:productId", middlewares.ProductAuthorization(), controllers.GetProduct)
		productRouter.DELETE("/:productId", middlewares.UserAuthorization(), controllers.DeleteProduct)
	}

	return r
}
