package main

import (
	"fmt"
	"log"

	"github.com/alinepinheir05/empresa-gestao-produtos/internal/domain"
	"github.com/alinepinheir05/empresa-gestao-produtos/internal/repository"
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

	productRepo := repository.NewProductRepository(db)

	p := &domain.Product{
		Name:         "Produto A",
		Code:         "A001",
		Description:  "Descrição do Produto A",
		Unit:         "un",
		CostEstimate: 50.75,
	}

	if err := productRepo.Create(p); err != nil {
		log.Fatal("Erro criando produto:", err)
	}
	fmt.Println("Produto criado:", p.Name)

	products, err := productRepo.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Produtos no banco:")
	for _, prod := range products {
		fmt.Printf("- %s (%s)\n", prod.Name, prod.Code)
	}
}
