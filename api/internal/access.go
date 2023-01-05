package internal

import "github.com/gogf/gf/v2/frame/g"

type (
	AccessListReq struct {
		g.Meta `method:"GET" path:"/access"`
	}

	AccessListRes struct {
		List []g.Map
	}
)

type (
	AccessItem struct {
		Alias     string
		RowKey    []string
		FieldsGet g.Map
	}

	RequestItem struct {
		Tag     string
		Request []g.Map
	}

	AddItem struct {
		TableName string
		Access    AccessItem
		Request   RequestItem
	}

	AccessAddReq struct {
		g.Meta `method:"POST" path:"/access"`
		Data   []AddItem
	}

	AccessAddRes struct {
	}
)

type (
	AccessUpdateReq struct {
		g.Meta `method:"PUT" path:"/access"`
		Id     int
		Data   g.Map
	}

	AccessUpdateRes struct {
	}
)

type (
	AccessOptionListReq struct {
		g.Meta `method:"GET" path:"/access/options"`
	}

	AccessOptionListRes struct {
		RoleList           []string
		RowKeyGenList      []string
		QueryExecutorList  []string
		ActionExecutorList []string
	}
)

type (
	AccessTableFieldsListReq struct {
		g.Meta `method:"GET" path:"/access/table-fields"`
		Table  string
	}

	AccessTableFieldsListRes struct {
		List []g.Map
	}
)

type (
	AccessTableListReq struct {
		g.Meta `method:"GET" path:"/access/table"`
		Table  string
	}

	AccessTableListRes struct {
		List []g.Map
	}
)
