package main

import (
	"log" 
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func OpenConnection() *gorm.DB {
    // Please define your username and password for MySQL.
    db, err := gorm.Open("mysql", "admin:admin@tcp(localhost:3306)/Pessoas?charset=utf8&parseTime=True")
    // NOTE: See we’re using = to assign the global var
    // instead of := which would assign it only in this function

    if err!=nil{
    log.Println("Connection Failed to Open")
    }else{ 
    log.Println("Connection Established")
    }

	return db
}