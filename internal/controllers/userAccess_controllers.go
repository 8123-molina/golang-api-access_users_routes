package controllers

import (
	"net/http"
	"strconv"

	"github.com/8123-molina/golang-api-access_users_routes/internal/services"
	"github.com/8123-molina/golang-api-access_users_routes/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetUsersAccessController(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	path := c.Query("path")
	method := c.Query("method")

	if err != nil {
		c.JSON(400, gin.H{"error": "user_id deve ser um número válido"})
		return
	}

	if path == "" {
		c.JSON(400, gin.H{"error": "Verifique o campo informado"})
		return
	}

	if method == "" || !utils.IsValidMethod(method) {
		c.JSON(400, gin.H{"error": "Verifique o campo informado methodo é invalido"})
		return
	}

	if path == "" || !utils.IsValidPath(path) {
		c.JSON(400, gin.H{"error": "Verifique o campo informado path é invalido"})
		return
	}

	access := services.GetUsersAccessService(userID, path, method)

	msg := "Acesso permitido"
	if !access {
		msg = "Acesso negado, entre em contato com administrador!"
	}
	c.JSON(http.StatusOK, gin.H{"message": msg, "access": access})

}
