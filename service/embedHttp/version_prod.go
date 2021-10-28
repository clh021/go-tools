// +build prod

// $ go run -tags prod .
// Version "0.0.1"

// embed 之前我们可以采取通过 -ldflags 的方法
// 将版本动态的赋值到变量。
// 现在，可以通过 embed 方式赋值
package embedHttp

import (
	_ "embed"
)

//go:embed version.txt
var version string
