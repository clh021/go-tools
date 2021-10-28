package embedHttp

import (
	"embed"
	"flag"
	"fmt"
	"log"
)

func Main() {

	t := flag.String("tag", "default", "embed example tag")
	flag.Parse()

	switch *t {
	case "ginDist":
		ginDistServe()

	case "ginTmpl":
		ginTmplServe()

	case "template":
		tmplServe()

	case "webDist":
		webDistServe()

	default:
		defaultServe()
	}
}

//go:embed hello.txt
var s string

//go:embed hello.txt
var b []byte

//go:embed hello.txt
//go:embed webDist
var webfs embed.FS

func defaultServe() {
	log.Println(version)
	print(s)
	print(string(b))
	data, _ := f.ReadFile("hello.txt")
	print(string(data))

	entries, err := webfs.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			panic(err)
		}
		fmt.Println(info.Name(), info.Size(), info.IsDir())
	}
}
