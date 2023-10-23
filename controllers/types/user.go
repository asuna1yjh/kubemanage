package types

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type UserRegisterResponse struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}
