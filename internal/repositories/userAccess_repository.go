package repositories

import (
	"fmt"
	"log"

	"github.com/8123-molina/golang-api-access_users_routes/configs/database"
	"github.com/8123-molina/golang-api-access_users_routes/internal/models"
)

func GetUsersAccessRepository(user_id int, path string) ([]models.MethodHTTP, error) {

	if database.DB == nil {
		log.Fatal("Database connection is not established.")
		fmt.Println("error: database connection is not established")
	}

	rows, err := database.DB.Query("SELECT id, method FROM methods_http")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var methods []models.MethodHTTP

	for rows.Next() {
		var method models.MethodHTTP

		err = rows.Scan(&method.ID, &method.Method)
		if err != nil {
			return nil, err
		}

		methods = append(methods, method)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return methods, nil
}
