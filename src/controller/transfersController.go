package controller

import "github.com/LGuilhermeMoreira/bank_api/src/database"

type TransferController struct {
	conn *database.Connection
}
