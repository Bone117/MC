package response

import "server/model"

type ReviewResponse struct {
	ReviewUserId   uint       `json:"reviewUserId"`
	ReviewUserName string     `json:"reviewUserName"`
	Sign           model.Sign `json:"sign"`
}
