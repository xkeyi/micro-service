package basic

import (
	"ms-user-service/basic/config"
	"ms-user-service/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}