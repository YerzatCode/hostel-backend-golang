package main

import (
	"net/http"

	"github.com/YerzatCode/sql_/internal/config"
	"github.com/YerzatCode/sql_/internal/database"
	"github.com/YerzatCode/sql_/internal/handler"
	"github.com/YerzatCode/sql_/internal/repository"
	"github.com/YerzatCode/sql_/internal/service"
)

func main() {

	cfg := config.InitConfig()
	db := database.DatabaseInit(&cfg.Database)

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	r := handler.InitRoutes()

	http.ListenAndServe(cfg.Port, r)
}
