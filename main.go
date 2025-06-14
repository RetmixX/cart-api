package main

import (
	"cart-api/config"
	"cart-api/internal/repository"
	"cart-api/internal/service"
	v1 "cart-api/internal/transport/rest/controllers/v1"
	"cart-api/internal/transport/server"
	"cart-api/pkg/db"
	"cart-api/pkg/log"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	customLogger := log.NewLogger()
	dbConnection := db.MustStartDB(&cfg.DBConf, customLogger)
	defer db.MustCloseDB(dbConnection, customLogger)

	productRepo := repository.New(dbConnection)
	productService := service.New(productRepo, customLogger)

	productCtrl := v1.New(productService, customLogger)

	srv := server.New(cfg, customLogger, productCtrl)

	go srv.StartServer()

	defer srv.StopServer()

	sgn := make(chan os.Signal, 1)

	signal.Notify(sgn, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-context.Background().Done():
	case <-sgn:
	}

	/*err := dbConnection.AutoMigrate(&entity.Product{}, &entity.Cart{}, &entity.Status{}, &entity.Order{})
	if err != nil {
		panic(err)
	}

	var cart entity.Cart

	dbConnection.Where("is_ordered = ?", false).Preload("Products").First(&cart)
	fmt.Println(cart.Products)*/

	//var products []entity.Product

	//err = dbConnection.Model(&cart).Preload("Products").Find(&products).Error

	//fmt.Println(products)

}
