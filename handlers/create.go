package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/swxtz/api-go/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		log.Printf("Erro ao inserir no banco: %v", err)
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Erro ao inserir no banco: %v", err),
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo criado com sucesso! ID: %d", id),
			"ID":      id,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
