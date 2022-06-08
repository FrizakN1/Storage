package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"storage/setting"
	"storage/utils"
)

var Link *sql.DB

func Connection(setting *setting.Setting) {
	var e error
	Link, e = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		setting.DbAddress,
		setting.DbPort,
		setting.DbUser,
		setting.DbPass,
		setting.DbName,
	))
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	e = Link.Ping()
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	errors := make([]string, 0)

	errors = append(errors, prepareRequest()...)

	if len(errors) > 0 {
		for _, i := range errors {
			utils.Logger.Println(i)
		}
	}
}
