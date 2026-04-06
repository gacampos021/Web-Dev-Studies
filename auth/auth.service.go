package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  string `json:"user"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email e password são obrigatórios",
		})
		return
	}

	// 🔴 AQUI VOCÊ VERIFICARIA NO BANCO DE DADOS
	// Por enquanto, vamos aceitar qualquer email/password
	// MUDE ISSO DEPOIS!
	if req.Email != "user@example.com" || req.Password != "senha123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "email ou senha incorretos",
		})
		return
	}

	token, err := GenerateToken("user123", req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Token: token,
		User:  req.Email,
	})
}
