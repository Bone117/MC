package request

type ReviewRequest struct {
	ReviewId uint   `json:"reviewId"`
	UserId   []uint `json:"userId" `
	SignId   []uint `json:"signId" `
	JieCiId  uint   `json:"jieCiId"`
}
type EvaluateRequest struct {
	ReviewId uint   `json:"reviewId" `
	UserId   []uint `json:"userId" binding:"required"`
	SignId   []uint `json:"signId" binding:"required"`
	JieCiId  uint   `json:"jieCiId" binding:"required"`
}
