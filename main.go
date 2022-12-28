package main

import (
	"log"
	"net/http"
	"os"

	"github.com/edwinndev/iotapi-mmj/database"
	"github.com/edwinndev/iotapi-mmj/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	startApp()
}

func startApp() {
	database.Connect()

	//_ = database.Mysql.AutoMigrate(&models.Device{})
	//_ = database.Mysql.AutoMigrate(&models.Thermometer{})
	//_ = database.Mysql.AutoMigrate(&models.Switch{})
	//_ = database.Mysql.AutoMigrate(&models.User{})
	//_ = database.Mysql.AutoMigrate(&models.Upload{})

	router := mux.NewRouter().PathPrefix("/api/v2").Subrouter()
	configureRouter(router)

	PORT := os.Getenv("PORT")
	if len(PORT) < 1 {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

func configureRouter(router *mux.Router) {
	routes.ConfigureDeviceRouter(router.PathPrefix("/devices").Subrouter())
	routes.ConfigureThermometerRouter(router.PathPrefix("/ths").Subrouter())
	routes.ConfigureSwitchsRouter(router.PathPrefix("/switchs").Subrouter())
	routes.ConfigureAuthRouter(router.PathPrefix("/auth").Subrouter())
	routes.ConfigureUsersRouter(router.PathPrefix("/users").Subrouter())
	routes.ConfigureUploadsRouter(router.PathPrefix("/uploads").Subrouter())
}
