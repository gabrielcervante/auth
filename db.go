package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=cervante password=cervantepswd dbname=auth port=5432"

var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

func init() {

	if err != nil {
		log.Fatal(err)
	}
}
