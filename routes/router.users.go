package routes

import (
	"net/http"

	"github.com/edwinndev/iotapi-mmj/commons"
	"github.com/edwinndev/iotapi-mmj/database"
	"github.com/edwinndev/iotapi-mmj/middlewares"
	"github.com/edwinndev/iotapi-mmj/models"
	"github.com/gorilla/mux"
)

func ConfigureUsersRouter(router *mux.Router) {
	router.HandleFunc("", middlewares.AdminMiddleware(findAllUsers)).Methods(http.MethodGet)
	router.HandleFunc("/{id}", middlewares.AdminMiddleware(findUserById)).Methods(http.MethodGet)
}

func findAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	errorDB := database.Mysql.Find(&users)
	if errorDB.Error != nil {
		commons.ApiBadRequest(w, "Error al obtener usuarios")
	} else {
		var usersResponse []models.User

		for _, value := range users {
			value.Password = ""
			usersResponse = append(usersResponse, value)
		}
		commons.ApiOK(w, usersResponse)
	}
}

func findUserById(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var id = mux.Vars(r)["id"]

	errorDB := database.Mysql.Find(&user, "id=?", id)
	if errorDB.Error != nil {
		commons.ApiNotFound(w, "El usuario no existe")
	} else {
		user.Password = ""
		commons.ApiOK(w, user)
	}
}
