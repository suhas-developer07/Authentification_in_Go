package main

import (
	"net/http"

	"github.com/suhas-developer07/Authentification_in_Go/db"
	"github.com/suhas-developer07/Authentification_in_Go/routes"
)

func main(){

	handler := routes.MountRoutes()
	db.InitDB()

	server := &http.Server{
		Addr: ":8080",
		Handler: handler,
	}

	
	server.ListenAndServe()
	
}
