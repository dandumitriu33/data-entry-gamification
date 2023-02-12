package model

import "mime/multipart"

type UserInfo struct {
	UserID   int64  `json:"user_id"`
	Points   int64  `json:"points"`
	Level    int64  `json:"level"`
	ImageURI string `json:"img_uri"`
}

type UserAvatar struct {
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}
