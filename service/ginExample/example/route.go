package example

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.Engine) *gin.Engine {
	r.GET("/ping", PingHandle)
	r.GET("/welcome", WelcomHandle)
	r.POST("/form_post", FormPostHandle)
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", UploadHandle)

	// $ curl -v localhost:8088/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
	// $ curl -v localhost:8088/thinkerou/not-uuid
	r.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})

	r.GET("/moreJSON", MoreJSONHandle)

	r.GET("/panic", func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		panic("foo")
	})

	return r
}
