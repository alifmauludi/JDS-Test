package user

import (
	"errors"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	GetUserByID(id int) (User, error)
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

	usernameCheck, err := s.repository.FindByUsername(input.Username)
	if err != nil {
		return usernameCheck, err
	}

	if usernameCheck.Username != "" {
		return user, errors.New("Username has been registered!")
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	newUser.PlainPassword = password

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	username := input.Username
	password := input.Password

	user, err := s.repository.FindByUsername(username)

	if err != nil {
		return user, err
	}

	if user.UserID == 0 {
		return user, errors.New("User not found!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.UserID == 0 {
		return user, errors.New("No user found with that ID")
	}

	return user, nil
}
