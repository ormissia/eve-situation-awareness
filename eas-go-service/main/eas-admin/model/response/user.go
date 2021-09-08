package response

import (
	"eas-go-service/main/eas-admin/model"
)

type Login struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
	ExpiresAt int64      `json:"expiresAt"`
}
