package embedHttp

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed webDist
var ginDist embed.FS

func ginDistServe() {
	r := gin.Default()

	r.StaticFS("/", http.FS(ginDist))

	r.StaticFS("/static/", http.FS(ginDist))

	r.Run(":8080")
}

// type webFS struct {
// 	content embed.FS
// }

// func (c webFS) Open(name string) (fs.File, error) {
// 	return c.content.Open(path.Join("web", name))
// }
// func main() {

// 	app.Use("/", filesystem.New(filesystem.Config{
// 		Root: http.FS(webFS{f}),
// 	}))
// }
