package response

import "server/model"

type AuthorityResponse struct {
	Authority model.Authority `json:"authority"`
}
