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

func ConfigureSwitchsRouter(router *mux.Router) {
	router.HandleFunc("", createSwitch).Methods(http.MethodPost)
	router.HandleFunc("", findAllSwitchs).Methods(http.MethodGet)
	router.HandleFunc("/{device}", findSwitchValue).Methods(http.MethodGet)
	router.HandleFunc("", updateSwitchValue).Methods(http.MethodPut)
	router.HandleFunc("/{device}", middlewares.AdminMiddleware(deleteSwitch)).Methods(http.MethodDelete)
}

func createSwitch(w http.ResponseWriter, r *http.Request) {
	var device models.Switch
	jsonError := json.NewDecoder(r.Body).Decode(&device)
	if jsonError != nil {
		commons.ApiBadRequest(w, "Los datos enviados son invalidos")
	} else {
		exists := database.Mysql.First(&models.Switch{}, "device=?", device.Device)
		if exists.Error == nil {
			commons.ApiBadRequest(w, "El dispositivo ya existe")
		} else {
			errorDB := database.Mysql.Create(&device)
			if errorDB.Error != nil {
				commons.ApiBadRequest(w, "Error al registrar dispositivo")
			} else {
				commons.ApiCreated(w, device)
			}
		}
	}
}

func findAllSwitchs(w http.ResponseWriter, r *http.Request) {
	var devices []models.Switch
	errorDB := database.Mysql.Find(&devices)
	if errorDB.Error != nil {
		commons.ApiBadRequest(w, "Error al obtener dispositivos")
	} else {
		commons.ApiOK(w, devices)
	}
}

func findSwitchValue(w http.ResponseWriter, r *http.Request) {
	var device models.Switch
	id := mux.Vars(r)["device"]
	errorDB := database.Mysql.First(&device, "device=?", id)
	if errorDB.Error != nil {
		commons.ApiNotFound(w, "El dispositivo no existe")
	} else {
		commons.ApiOK(w, device)
	}
}

func updateSwitchValue(w http.ResponseWriter, r *http.Request) {
	var device models.Switch
	var value bool
	jsonError := json.NewDecoder(r.Body).Decode(&device)
	if jsonError != nil {
		commons.ApiBadRequest(w, "Los datos enviados son invalidos")
	} else {
		value = device.Value
		exists := database.Mysql.First(&device, "device=?", device.Device)
		if exists.Error != nil {
			commons.ApiNotFound(w, "El dispositivo no existe")
		} else {
			device.Value = value
			errDB := database.Mysql.Save(&device)
			if errDB.Error != nil {
				commons.ApiBadRequest(w, "No se puedo actualizar el ultimo valor")
			} else {
				commons.ApiOK(w, device)
			}
		}
	}
}

func deleteSwitch(w http.ResponseWriter, r *http.Request) {
	var device models.Switch
	id := mux.Vars(r)["device"]
	errorDB := database.Mysql.First(&device, "device=?", id)
	if errorDB.Error != nil {
		commons.ApiNotFound(w, "El dispositivo no existe")
	} else {
		deleteError := database.Mysql.Delete(&device, "device=?", device.Device)
		if deleteError.Error != nil {
			commons.ApiNotFound(w, "Error al eliminar el dispositivo")
		} else {
			commons.ApiOK(w, device)
		}
	}
}
