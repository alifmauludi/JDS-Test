package user

type RegisterUserInput struct {
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
