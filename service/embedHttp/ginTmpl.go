package embedHttp

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
)

//go:embed templates
var ginTmpl embed.FS

func ginTmplServe() {
	r := gin.Default()
	t, _ := template.ParseFS(ginTmpl, "templates/*.tmpl")
	r.SetHTMLTemplate(t)
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.tmpl", gin.H{"title": "Golang Embed 测试"})
	})
	r.Run(":8080")
}
