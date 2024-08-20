package routes

import (
    "github.com/gin-gonic/gin"
    "93HW/controllers"
    "93HW/middleware"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    router.Use(middleware.CORSMiddleware())
    router.Use(middleware.RateLimiter())
    
    userRoutes := router.Group("/user")
    {
        userRoutes.Use(middleware.AuthMiddleware())
        userRoutes.GET("/", controllers.GetUser)
    }

    return router
}
