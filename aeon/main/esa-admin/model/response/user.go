package response

import (
	"esa-go-service/main/esa-admin/model"
)

type Login struct {
	User      model.User `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expiresAt"`
}
