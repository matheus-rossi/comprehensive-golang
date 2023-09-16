package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type User struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

func saveUserToFile(user User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = os.WriteFile("user.json", userJSON, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Dados do usuário salvos em disco: %s\n", userJSON)
	return nil
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Lista de Usuários")
}

func insertUsers(w http.ResponseWriter, r *http.Request) {
	var user User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Falha ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Falha ao decodificar JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Dados do usuário recebidos: %+v\n", user)

	err = saveUserToFile(user)
	if err != nil {
		http.Error(w, "Falha ao salvar os dados do usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Dados do usuário recebidos e salvos com sucesso!")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", listUsers).Methods("GET")
	router.HandleFunc("/users", insertUsers).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
