package main

import (
	_ "gf-init/boot"
	_ "gf-init/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
