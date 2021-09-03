package editor

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.Engine) *gin.Engine {
	// Simple group: v2
	// v2 := r.Group("/v2")
	// {
	// 	v2.POST("/login", loginEndpoint)
	// 	v2.POST("/submit", submitEndpoint)
	// 	v2.POST("/read", readEndpoint)
	// }

	v1 := r.Group("/editor")
	{
		v1.POST("/read", ReadHandle)
	}
	return r
}
