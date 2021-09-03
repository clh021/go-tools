package fileModule

import (
	"runtime"

	"github.com/filebrowser/filebrowser/v2/cmd"
)

func Main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
