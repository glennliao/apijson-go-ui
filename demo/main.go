package main

import (
	"github.com/glennliao/apijson-go-ui"
	"github.com/glennliao/apijson-go/framework"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	framework.Init()

	s := g.Server()
	apijson_go_ui.Handler(s)
	s.Run()
}
