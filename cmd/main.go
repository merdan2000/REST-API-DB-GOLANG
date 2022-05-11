package main

import (
	"github.com/merdan2000/internal/settings"
	"github.com/merdan2000/migrations"
	"log"
)

func main() {
	set := settings.NewSettings()
	err := migrations.MigrationUp(set)
	if err != nil {
		log.Fatalln(err)
		return
	}

}
