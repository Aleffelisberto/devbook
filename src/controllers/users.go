package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário criado"))
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
