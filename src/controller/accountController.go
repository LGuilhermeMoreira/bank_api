package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LGuilhermeMoreira/bank_api/src/utils/dto"
)

//var db = database.NewConnection()

type accountController struct {
}

func (a accountController) HandlePostAccount(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	// criando o modelo que vai receber o
	var account dto.DtoAccount

	if err := json.NewDecoder(body).Decode(&account); err != nil {
		panic(err)
	}

	fmt.Println(account.ID, account.Owner, account.Balance)

	w.WriteHeader(http.StatusOK)

}

func NewAccountController() *accountController {
	return &accountController{}
}
