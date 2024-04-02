package controllers

import (
	"curso/devbook/src/db"
	"curso/devbook/src/models"
	"curso/devbook/src/repositories"
	"curso/devbook/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	ID, err := repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	user.ID = ID
	responses.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuários retornados"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário retornado"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário atualizado"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário removido"))
}
