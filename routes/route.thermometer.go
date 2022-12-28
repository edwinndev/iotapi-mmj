package routes

import (
	"encoding/json"
	"net/http"

	"github.com/edwinndev/iotapi-mmj/commons"
	"github.com/edwinndev/iotapi-mmj/database"
	"github.com/edwinndev/iotapi-mmj/middlewares"
	"github.com/edwinndev/iotapi-mmj/models"
	"github.com/gorilla/mux"
)

func ConfigureThermometerRouter(router *mux.Router) {
	router.HandleFunc("", findAllThermometers).Methods(http.MethodGet)
	router.HandleFunc("/{device}", findLastValuesByDevice).Methods(http.MethodGet)
	router.HandleFunc("", saveThermometerValue).Methods(http.MethodPost)
	router.HandleFunc("/{device}", middlewares.JWTMiddlweare(deleteThermometerTableValue)).Methods(http.MethodDelete)
}

func findAllThermometers(w http.ResponseWriter, r *http.Request) {
	var ths []models.Thermometer

	err := database.Mysql.Find(&ths)
	if err.Error != nil {
		commons.ApiBadRequest(w, "Error al obtener dispositivos")
	} else {
		commons.ApiOK(w, ths)
	}
}

func _(w http.ResponseWriter, r *http.Request) {
	deviceName := mux.Vars(r)["device"]
	if len(deviceName) < 1 {
		commons.ApiBadRequest(w, "Solicitud invalida")
	} else {
		var ths []models.Thermometer

		err := database.Mysql.Find(&ths, "device=?", deviceName)
		if err.Error != nil {
			commons.ApiBadRequest(w, "Error al obtener valores")
		} else {
			commons.ApiOK(w, ths)
		}
	}
}

func findLastValuesByDevice(w http.ResponseWriter, r *http.Request) {
	deviceName := mux.Vars(r)["device"]
	var th models.Thermometer
	database.Mysql.Order("id DESC").Limit(1).First(&th)
	exec := database.Mysql.Exec("SELECT * FROM thermometers WHERE device=? ORDER BY id DESC LIMIT 1", deviceName)
	if exec.Error != nil {
		commons.ApiBadRequest(w, "Error al obtener valores")
	} else {
		commons.ApiOK(w, th)
	}
}

func saveThermometerValue(w http.ResponseWriter, r *http.Request) {
	var th models.Thermometer

	err := json.NewDecoder(r.Body).Decode(&th)
	if err != nil {
		commons.ApiBadRequest(w, "Formato de solicitud invalido")
	} else {
		errDB := database.Mysql.Create(&th)
		if errDB.Error != nil {
			commons.ApiBadRequest(w, "Error al insertar valor")
		} else {
			commons.ApiCreated(w, th)
		}
	}
}

func deleteThermometerTableValue(w http.ResponseWriter, r *http.Request) {
	deviceName := mux.Vars(r)["device"]
	if len(deviceName) < 1 {
		commons.ApiBadRequest(w, "Es necesario el nombre del dispositivo")
	} else {
		var ths []models.Thermometer

		err := database.Mysql.Exec("DELETE FROM thermometers WHERE device=?", deviceName)
		if err.Error != nil {
			commons.ApiBadRequest(w, "Error al eliminar datos")
		} else {
			commons.ApiOK(w, ths)
		}
	}
}
