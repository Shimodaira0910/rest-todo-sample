package main

import (
	"fmt"
	"rest-todo-sample/db"
	"rest-todo-sample/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully! Migrated!")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
