package internal

import "github.com/gogf/gf/v2/frame/g"

type (
	RequestListReq struct {
		g.Meta `method:"GET" path:"/request"`
	}

	RequestListRes struct {
		List []g.Map
	}
)

type (
	RequestUpdateReq struct {
		g.Meta `method:"PUT" path:"/request"`
		Data   g.Map
	}

	RequestUpdateRes struct {
	}
)
