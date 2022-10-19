package response

import "server/model"

type ExaFileResponse struct {
	File model.File `json:"file"`
}
