package ginExample

import (
	"flag"
	"io"
	"net/http"
	"os"
	"test/service/ginExample/editor"
	"test/service/ginExample/example"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func GetRouter() *gin.Engine {

	// Default With the Logger and Recovery middleware already attached
	// r := gin.Default()

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	// r.Use(gin.Logger())

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r = example.SetupRouter(r)
	r = editor.SetupRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome for you!")
	})

	r.Static("/assets", "./assets")
	r.StaticFS("/more_static", http.Dir("my_file_system"))
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	return r
}
func Main() {
	webPort := flag.String("port", "8000", "gin server port")
	flag.Parse()
	r := GetRouter()
	r.Run(":" + *webPort)
}
