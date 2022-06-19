package request

type NoticeRequest struct {
	NoticeId uint `form:"id" json:"id"`
}

type UpdateNoticeRequest struct {
	NoticeId uint   `json:"id"`
	Title    string `json:"title" `
	Desc     string `json:"desc" `
	Content  string `json:"content"`
}
