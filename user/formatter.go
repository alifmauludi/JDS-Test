package user

type UserFormatter struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type UserLoggedinFormatter struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type TokenValidationFormatter struct {
	IsValid   bool   `json:"is_valid"`
	ExpiredAt string `json:"expired_at"`
	Username  string `json:"username"`
}

func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{
		Username: user.Username,
		Role:     user.Role,
		Password: user.PlainPassword,
	}

	return formatter
}

func FormatLoggedinUser(user User, token string) UserLoggedinFormatter {
	formatter := UserLoggedinFormatter{
		Username: user.Username,
		Role:     user.Role,
		Token:    token,
	}

	return formatter
}

func FormatTokenValidation(user User, expired_at string, is_valid bool) TokenValidationFormatter {
	formatter := TokenValidationFormatter{
		IsValid:   is_valid,
		ExpiredAt: expired_at,
		Username:  user.Username,
	}

	return formatter
}
