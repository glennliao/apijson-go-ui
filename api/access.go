package api

import (
	"context"
	. "github.com/glennliao/apijson-go-ui/api/internal"
	"github.com/glennliao/apijson-go/config"
	"github.com/glennliao/apijson-go/config/executor"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
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

func (a *AccessApi) Add(ctx context.Context, req *AccessAddReq) (res *AccessAddRes, err error) {

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		db := g.DB().Ctx(ctx)
		accessModel := db.Model("_access")

		for _, item := range req.Data {

			tableName := item.TableName
			alias := item.Access.Alias

			one, err := accessModel.Clone().One("alias", alias)
			if err != nil {
				return err
			}
			if !one.IsEmpty() {
				g.Log().Warning(ctx, "已存在alias:"+alias)
				return gerror.New("已存在alias:" + alias)
			}

			_, err = accessModel.Clone().Insert(g.Map{
				"name":        tableName,
				"alias":       alias,
				"get":         "[]",
				"head":        "[]",
				"gets":        "[]",
				"heads":       "[]",
				"post":        "[]",
				"put":         "[]",
				"delete":      "[]",
				"detail":      tableName,
				"row_key":     strings.Join(item.Access.RowKey, ","),
				"row_key_gen": "",
				"field_get":   item.Access.FieldsGet,
				"executor":    "",
			})
			if err != nil {
				return err
			}

			// request

			requestModel := db.Model("_request")
			for _, reqItem := range item.Request.Request {
				_, err := requestModel.Clone().Insert(g.Map{
					"version": "1",
					"method":  reqItem["method"],
					"tag":     item.Request.Tag,
					"structure": g.Map{
						item.Request.Tag: g.Map{
							"MUST":   g.Map{},
							"UPDATE": g.Map{},
						},
					},
				})
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	return &AccessAddRes{}, err
}

func (a *AccessApi) Update(ctx context.Context, req *AccessUpdateReq) (res *AccessUpdateRes, err error) {

	db := g.DB()

	model := db.Model("_access")

	_, err = model.Update(req.Data, g.Map{"id": req.Id})

	return &AccessUpdateRes{}, err
}

func (a *AccessApi) OptionsList(ctx context.Context, req *AccessOptionListReq) (res *AccessOptionListRes, err error) {

	return &AccessOptionListRes{
		RoleList:           config.RoleList(),
		RowKeyGenList:      config.RowKeyGenList(),
		ActionExecutorList: executor.ActionExecutorList(),
		QueryExecutorList:  executor.QueryExecutorList(),
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

func (a *AccessApi) TableList(ctx context.Context, req *AccessTableListReq) (res *AccessTableListRes, err error) {

	fields, err := g.DB().Tables(ctx, req.Table)

	var list []g.Map
	for _, table := range fields {
		list = append(list, g.Map{
			"label": table,
			"value": table,
		})
	}

	return &AccessTableListRes{
		List: list,
	}, err
}
