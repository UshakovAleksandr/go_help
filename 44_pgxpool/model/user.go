package model

// User model.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserResp dto.
type UserResp struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UserUpdate struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
