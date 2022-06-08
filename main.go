package main

import (
	"storage/database"
	"storage/router"
	"storage/setting"
)

func main() {
	s := setting.LoadSetting("setting.json")
	database.Connection(s)
	_ = router.Initialized().Run(s.Address + ":" + s.Port)
}
