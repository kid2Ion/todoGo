package main

import (
	"fmt"
	"todoGo/db"
	"todoGo/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("migrate成功にゃ")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
