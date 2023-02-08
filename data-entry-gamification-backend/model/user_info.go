package model

type UserInfo struct {
	UserID int64 `json:"user_id"`
	Points int64 `json:"points"`
	Level  int64 `json:"level"`
}
