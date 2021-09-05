package crudApi

import (
	"fmt"

	"github.com/clh021/crud-api/cmd"
)

var build = "not set"

func Main() {
	fmt.Printf("Build: %s\n", build)
	cmd.Execute()
}
