package request

type ReviewRequest struct {
	ReviewId uint   `json:"reviewId"`
	UserId   []uint `json:"userId" `
	SignId   []uint `json:"signId" `
	JieCiId  uint   `json:"jieCiId"`
}
