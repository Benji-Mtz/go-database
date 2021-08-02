package main

import (
	"github/Benji-Mtz/go-database/pkg/product"
	"github/Benji-Mtz/go-database/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

}
