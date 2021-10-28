// +build !prod

// $ go run .
// Version "dev"

// embed 之前我们可以采取通过 -ldflags 的方法
// 将版本动态的赋值到变量。
// 现在，可以通过 embed 方式赋值
package embedHttp

var version string = "dev"
