package request

type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}
