package dto

import (
	"github.com/LGuilhermeMoreira/bank_api/src/utils/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginAccountInput struct {
	UserMail     string `json:"user_mail"`
	UserPassword string `json:"user_password"`
}

func (d LoginAccountInput) ConvertInputToModel() (*models.LoginAccountModel, error) {

	bytePassword, err := bcrypt.GenerateFromPassword([]byte(d.UserPassword), 10)

	if err != nil {
		return nil, err
	}

	return &models.LoginAccountModel{
		ID:           uuid.New(),
		UserMail:     d.UserMail,
		UserPassword: string(bytePassword),
	}, nil
}
