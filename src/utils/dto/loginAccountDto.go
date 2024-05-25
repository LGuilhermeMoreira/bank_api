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

type LoginAccountOutput struct {
	ID    uuid.UUIDs `json:"login_id"`
	Token string     `json:"token"`
}

type LoginAccountVerifyOutput struct {
	ID        string `json:"login_id"`
	AccountID string `json:"account_id"`
	Owner     string `json:"owner"`
	Balance   string `json:"balance"`
	Token     string `json:"token"`
}

func NewLoginAccountVerify(id, account_id, owner, balance, token string) *LoginAccountVerifyOutput {
	return &LoginAccountVerifyOutput{
		ID:        id,
		AccountID: account_id,
		Owner:     owner,
		Balance:   balance,
		Token:     token,
	}
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
