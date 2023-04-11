package request

type ReviewRequest struct {
	ReviewId uint   `json:"reviewId"`
	UserId   []uint `json:"userId"  binding:"required"`
	SignId   []uint `json:"signId"  binding:"required"`
	JieCiId  uint   `json:"jieCiId" binding:"required"`
}
type EvaluateRequest struct {
	ReviewId uint   `json:"reviewId" `
	UserId   []uint `json:"userId" binding:"required"`
	SignId   []uint `json:"signId" binding:"required"`
	JieCiId  uint   `json:"jieCiId" binding:"required"`
}

type UpdateEvaluateRequest struct {
	UserId   uint
	SignId   uint   `json:"signId" binding:"required"`
	Score    uint   `json:"score" `
	Comments string `json:"comments"`
}
