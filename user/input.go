package user

type RegisterUserInput struct {
	Username string `json:"username" binding:"required" validate:"required"`
	Role     string `json:"role" binding:"required" validate:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}
