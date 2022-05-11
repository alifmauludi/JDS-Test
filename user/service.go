package user

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	length := 6

	generatedPassword := make([]byte, length)
	for i := range generatedPassword {
		generatedPassword[i] = charset[seededRand.Intn(len(charset))]
	}

	password := string(generatedPassword)

	user := User{}
	user.Username = input.Username
	user.Role = input.Role
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.CreatedDate = time.Now()

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	newUser.PlainPassword = password

	return newUser, nil
}
