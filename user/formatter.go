package user

type UserFormatter struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{
		Username: user.Username,
		Role:     user.Role,
		Password: user.PlainPassword,
	}

	return formatter
}
