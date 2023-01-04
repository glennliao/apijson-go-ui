package api

import (
	"context"
	. "github.com/glennliao/apijson-go-ui/api/internal"
	"github.com/gogf/gf/v2/frame/g"
)

type RequestApi struct{}

var Request = RequestApi{}

func (a *RequestApi) List(ctx context.Context, req *RequestListReq) (res *RequestListRes, err error) {

	db := g.DB()

	model := db.Model("_request")

	all, err := model.All()

	return &RequestListRes{
		List: all.List(),
	}, err
}

func (a *RequestApi) Update(ctx context.Context, req *RequestUpdateReq) (res *RequestUpdateRes, err error) {

	db := g.DB()

	model := db.Model("_request")

	_, err = model.Update(req.Data, g.Map{"id": req.Data["id"]})

	return &RequestUpdateRes{}, err
}
