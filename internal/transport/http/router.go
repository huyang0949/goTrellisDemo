package httptransport

import "github.com/gin-gonic/gin"

func NewRouter(greeterHandler *GreeterHandler, userHandler *UserHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/hello", greeterHandler.Hello)
	router.POST("/hello", greeterHandler.Hello)
	router.POST("/login", userHandler.Login)
	return router
}
