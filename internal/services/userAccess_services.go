package services

import (
	"log"

	"github.com/8123-molina/golang-api-access_users_routes/internal/models"
	"github.com/8123-molina/golang-api-access_users_routes/internal/repositories"
)

func GetUsersAccessService(user_id int, path string, method string) bool {
	var result []models.MethodHTTP

	result, err := repositories.GetUsersAccessRepository(user_id, path)
	log.Println(result)

	if err != nil {
		log.Println("Erro ao obter dados:", err)
		return false
	}
	return true
}
