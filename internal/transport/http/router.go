package httptransport

import "github.com/gin-gonic/gin"

func NewRouter(greeterHandler *GreeterHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/hello", greeterHandler.Hello)
	router.POST("/hello", greeterHandler.Hello)
	return router
}
