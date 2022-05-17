package main

import (
	"github.com/gin-gonic/gin"
	"github.com/merdan2000/internal/handler"
	"github.com/merdan2000/internal/repository"
	"github.com/merdan2000/internal/service"
	"github.com/merdan2000/internal/settings"
	"github.com/merdan2000/migrations"
	"log"
	"net/http"
)

func main() {

	set := settings.NewSettings()
	err := migrations.MigrationUp(set)
	if err != nil {
		log.Fatalln(err)
		return
	}

	repo := repository.NewRepository(set)
	ser := service.NewService(repo)
	hand := handler.NewHandler(ser)
	router := gin.New()
	users := router.Group("user")
	{
		users.GET("me", hand.GetUserById)
	}
	http.ListenAndServe(":8080", router)
}
