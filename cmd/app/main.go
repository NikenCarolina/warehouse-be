package main

import (
	"github.com/NikenCarolina/warehouse-be/internal/config"
	"github.com/NikenCarolina/warehouse-be/internal/handler"
	"github.com/NikenCarolina/warehouse-be/internal/repository/postgres"
	"github.com/NikenCarolina/warehouse-be/internal/router"
	"github.com/NikenCarolina/warehouse-be/internal/server"
)

func main() {
	config := config.InitConfig()
	db := postgres.Init(config)
	defer db.Close()
	handler := handler.Init(db, config)
	router := router.Init(handler, config)
	server := server.NewServer(config, router)
	server.Run()
}
