package routes

import (
	"encoding/json"
	"net/http"

	"github.com/edwinndev/iotapi-mmj/commons"
	"github.com/edwinndev/iotapi-mmj/database"
	"github.com/edwinndev/iotapi-mmj/handlers"
	"github.com/edwinndev/iotapi-mmj/models"
	"github.com/gorilla/mux"
)

func ConfigureAuthRouter(router *mux.Router) {
	router.HandleFunc("/register", authRegister).Methods(http.MethodPost)
	router.HandleFunc("/login", authLogin).Methods(http.MethodPost)
}

func authRegister(w http.ResponseWriter, r *http.Request) {
	var user models.User
	jsonError := json.NewDecoder(r.Body).Decode(&user)
	if jsonError != nil {
		commons.ApiBadRequest(w, "Los datos enviados no son validos")
	} else {
		if database.UserExists(user.Email) == nil {
			commons.ApiBadRequest(w, "El email ya esta en uso")
			return
		}
		user.Role = commons.RoleReader
		errorDB := database.Mysql.Create(&user)
		if errorDB.Error != nil {
			commons.ApiBadRequest(w, "No se pudo registra el usuario")
		} else {
			createUserWithToken(w, user)
		}
	}
}

func authLogin(w http.ResponseWriter, r *http.Request) {
	var login models.LoginRequest

	jsonError := json.NewDecoder(r.Body).Decode(&login)
	if jsonError != nil {
		commons.ApiBadRequest(w, "Los datos enviados no son validos")
	} else {
		var user models.User
		errorDB := database.Mysql.First(&user, "email=? AND password=?", login.Username, login.Password)
		if errorDB.Error != nil {
			commons.ApiNotFound(w, "El usuario no existe")
		} else {
			createUserWithToken(w, user)
		}
	}
}

func createUserWithToken(w http.ResponseWriter, user models.User) {
	token, tokenError := handlers.MakeToken(user)
	if tokenError != nil {
		commons.ApiBadRequest(w, "No se puede generar el token")
	} else {
		var response = models.LoginResponse{}
		response.ID = user.ID
		response.Names = user.Names
		response.Surnames = user.Surnames
		response.Email = user.Email
		response.Phone = user.Phone
		response.Role = user.Role
		response.CreatedAt = user.CreatedAt
		response.UpdatedAt = user.UpdatedAt
		response.DeletedAt = ""
		response.Token = token
		commons.ApiOK(w, response)
	}
}
