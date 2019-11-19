package main

import (
	"github.com/Linchpins-team/ilumi-backend/pkg/config"
	"github.com/Linchpins-team/ilumi-backend/pkg/database"
	"github.com/Linchpins-team/ilumi-backend/pkg/router"

	"log"
)

func main() {
	conf := config.Testing()
	log.Println("Got conf", conf)
	db := database.New(conf)
	log.Println("Got db", db)
	r := router.SetupServer(conf, db)
	log.Printf("Server has started on http://%s:%s\n", conf.Host, conf.Port)
	if err := r.Run(":" + conf.Port); err != nil {
		panic(err)
	}
}
