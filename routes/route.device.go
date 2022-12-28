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

func ConfigureDeviceRouter(router *mux.Router) {
	router.HandleFunc("", finDevices).Methods(http.MethodGet)
	router.HandleFunc("", createDevice).Methods(http.MethodPost)
	router.HandleFunc("", updateDevice).Methods(http.MethodPut)
	router.HandleFunc("/{id}", finDevice).Methods(http.MethodGet)
	router.HandleFunc("/{id}", middlewares.AdminMiddleware(deleteDevice)).Methods(http.MethodDelete)
	router.HandleFunc("/truncate", middlewares.AdminMiddleware(truncateTableDevices)).Methods(http.MethodPatch)
}

func finDevices(w http.ResponseWriter, r *http.Request) {
	var devices []models.Device

	err := database.Mysql.Find(&devices)
	if err.Error != nil {
		commons.ApiInternalServerError(w)
	} else {
		commons.ApiOK(w, devices)
	}
}

func finDevice(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) < 1 {
		commons.ApiBadRequest(w, "Especificar ID del dispositivo")
	} else {
		var device models.Device
		err := database.Mysql.First(&device, "id=?", id)
		if err.Error != nil {
			commons.ApiNotFound(w, "Dispositivo no encontrado")
		} else {
			commons.ApiOK(w, device)
		}
	}
}

func createDevice(w http.ResponseWriter, r *http.Request) {
	var device models.Device

	err := json.NewDecoder(r.Body).Decode(&device)
	if err != nil {
		commons.ApiBadRequest(w, "Formato de solicitud invalido")
		return
	}

	exists := database.Mysql.First(&models.Device{}, "name=?", device.Name)
	if exists.Error == nil {
		commons.ApiBadRequest(w, "El dispositivo ya existe")
		return
	}

	inserted := database.Mysql.Create(&device).Error
	if inserted != nil {
		commons.ApiBadRequest(w, "Error al guardar dispositivo")
	} else {
		commons.ApiCreated(w, device)
	}
}

func updateDevice(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	err := json.NewDecoder(r.Body).Decode(&device)
	if err != nil {
		commons.ApiBadRequest(w, "Formato de solicitud invalido")
	} else {
		var first models.Device
		res := database.Mysql.First(&first, "id=?", device.ID)
		if res.Error != nil {
			commons.ApiNotFound(w, "El dispositivo a actualizar no existe")
		} else {
			first.Code = device.Code
			first.Name = device.Name
			first.Description = device.Description
			dbError := database.Mysql.Save(&first).Error
			if dbError != nil {
				commons.ApiBadRequest(w, "No se puede actualizar datos del dispositivo")
			} else {
				commons.ApiOK(w, first)
			}
		}
	}
}

func deleteDevice(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	id := mux.Vars(r)["id"]

	if len(id) < 1 {
		commons.ApiBadRequest(w, "Solicitud invalida")
	} else {
		res := database.Mysql.Delete(&device, "id=?", id)
		if res.Error != nil {
			commons.ApiBadRequest(w, "Error al eliminar el dipositivo")
		} else {
			commons.ApiOK(w, device)
		}
	}
}

func truncateTableDevices(w http.ResponseWriter, r *http.Request) {
	erro := database.Mysql.Exec("TRUNCATE TABLE devices")
	if erro.Error != nil {
		commons.ApiBadRequest(w, "Error, no se puede truncar la tabla DEVICES")
	} else {
		commons.ApiOK(w, nil)
	}
}
