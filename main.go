package main

import (
	"cart-api/config"
	"cart-api/internal/model/entity"
	"cart-api/pkg/db"
	"cart-api/pkg/log"
	"fmt"
)

func main() {
	cfg := config.MustLoad()
	customLogger := log.NewLogger()
	dbConnection := db.MustStartDB(&cfg.DBConf, customLogger)
	defer db.MustCloseDB(dbConnection, customLogger)

	err := dbConnection.AutoMigrate(&entity.Product{}, &entity.Cart{}, &entity.Status{}, &entity.Order{})
	if err != nil {
		panic(err)
	}

	var cart entity.Cart

	dbConnection.Where("is_ordered = ?", false).Preload("Products").First(&cart)
	fmt.Println(cart.Products)

	//var products []entity.Product

	//err = dbConnection.Model(&cart).Preload("Products").Find(&products).Error

	//fmt.Println(products)

}
