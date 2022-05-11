package user

type UserFormatter struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

type UserLoggedinFormatter struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{
		Username: user.Username,
		Role:     user.Role,
		Password: user.PlainPassword,
	}

	return formatter
}

func FormatLoggedinUser(user User) UserLoggedinFormatter {
	formatter := UserLoggedinFormatter{
		Username: user.Username,
		Role:     user.Role,
	}

	return formatter
}
