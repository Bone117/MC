package request

type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

type GetByTime struct {
	currentTime string `json:"currentTime"` // 主键ID
}

type GetStage struct {
	CurrentTime string `json:"currentTime"`
}

type PageInfo struct {
	Page     int                    `json:"offset" form:"offset"`   // 页码
	PageSize int                    `json:"size" form:"size"`       // 每页大小
	Keyword  map[string]interface{} `json:"keyword" form:"keyword"` //关键字
}
