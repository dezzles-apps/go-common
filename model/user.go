package model

type User struct {
	UserId      string `json:"user_id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}
