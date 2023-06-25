package models

import (
	"GoRestAPI/ja"
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (params *LoginParams) EncodePassword() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (params *LoginParams) Validates() error {
	emailError := params.emailValidates()
	if emailError != nil {
		return emailError
	}

	passwordError := params.passwordValidates()
	if passwordError != nil {
		return passwordError
	}

	return nil
}

func (params *LoginParams) emailValidates() error {
	emailRegex := `[\w\-._]+@[\w\-._]+\.[A-Za-z]+`
	regex := regexp.MustCompile(emailRegex)

	if regex.MatchString(params.Email) {
		return nil
	}
	return errors.New(ja.EmailValidationErrorMessage)
}

func (params *LoginParams) passwordValidates() error {
	passwordMinLength := 8
	passwordMaxLength := 16

	var passwordLength = len(params.Password)
	if passwordLength >= passwordMinLength && passwordLength < passwordMaxLength {
		return nil
	}
	return errors.New(ja.PasswordValidationErrorMessage)
}
