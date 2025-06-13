package main

import (
	"log"

	"github.com/alinepinheir05/empresa-gestao-produtos/internal/domain"
	"github.com/alinepinheir05/empresa-gestao-produtos/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "alinepinheiro:1Tartaruga!@tcp(127.0.0.1:3306)/gestao?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro conexão banco:", err)
	}

	err = db.AutoMigrate(&domain.Product{})
	if err != nil {
		log.Fatal("Erro na migração:", err)
	}

	r := router.StartRouter(db)
	r.Run(":8080")
}
