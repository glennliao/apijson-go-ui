package main

import (
	"apijson-go-ui/api"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func response(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()

	defer func() {
		err := recover()
		if err != nil {
			g.Log().Error(r.Context(), err)
		}
	}()

	code := 200
	msg := ""
	data := r.GetHandlerResponse()
	err := r.GetError()
	if err != nil {
		code = 500
		msg = err.Error()
		data = g.Map{}
	}
	r.Response.WriteJson(Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(response)
		group.Bind(api.Access)
		group.Bind(api.Request)
	})

	s.Run()

}
