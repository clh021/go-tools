package embedHttp

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates
var tmpl embed.FS

func tmplServe() {
	t, err := template.ParseFS(tmpl, "templates/*.tmpl")
	if err != nil {
		panic(err)
	}

	// /hello?lang=xx.tmpl
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		t.ExecuteTemplate(rw, r.FormValue("lang"), nil)
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(rw, "index.tmpl", map[string]string{"title": "Golang Embed 测试"})
	})
	http.ListenAndServe(":8080", nil)
}
