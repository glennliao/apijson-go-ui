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
	AccessUpdateReq struct {
		g.Meta `method:"PUT" path:"/access"`
		Id     int
		Data   g.Map
	}

	AccessUpdateRes struct {
	}
)

type (
	AccessRoleListReq struct {
		g.Meta `method:"GET" path:"/access/role"`
	}

	AccessRoleListRes struct {
		List []string
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
