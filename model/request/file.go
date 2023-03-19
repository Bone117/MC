package request

type GetFileRequest struct {
	FileId uint `form:"id" json:"id"`
}
