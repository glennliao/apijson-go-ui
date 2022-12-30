package api

import (
	. "apijson-go-ui/api/internal"
	"context"
	"github.com/glennliao/apijson-go/config"
	"github.com/gogf/gf/v2/frame/g"
)

type AccessApi struct{}

var Access = AccessApi{}

func (a *AccessApi) List(ctx context.Context, req *AccessListReq) (res *AccessListRes, err error) {

	db := g.DB()

	model := db.Model("_access")

	all, err := model.All()

	return &AccessListRes{
		List: all.List(),
	}, err
}

func (a *AccessApi) Update(ctx context.Context, req *AccessUpdateReq) (res *AccessUpdateRes, err error) {

	db := g.DB()

	model := db.Model("_access")

	_, err = model.Update(req.Data, g.Map{"id": req.Id})

	return &AccessUpdateRes{}, err
}

func (a *AccessApi) RoleList(ctx context.Context, req *AccessRoleListReq) (res *AccessRoleListRes, err error) {

	return &AccessRoleListRes{
		List: config.RoleList(),
	}, err
}

func (a *AccessApi) TableFieldsList(ctx context.Context, req *AccessTableFieldsListReq) (res *AccessTableFieldsListRes, err error) {

	fields, err := g.DB().TableFields(ctx, req.Table)

	var list []g.Map
	for _, field := range fields {
		list = append(list, g.Map{
			"label": field.Name,
			"value": field.Name,
		})
	}

	return &AccessTableFieldsListRes{
		List: list,
	}, err
}
