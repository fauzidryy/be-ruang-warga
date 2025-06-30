package delivery

import (
	"be-ruang-warga/internal/user/usecase"
	"be-ruang-warga/utils"
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

type UserHandler struct {
	UC usecase.UserUsecase
}

func NewUserHandler(router *gin.RouterGroup, uc usecase.UserUsecase) {
	h := &UserHandler{UC: uc}

	router.POST("/auth/google", h.GoogleAuthHandler)
}

func (h *UserHandler) GoogleAuthHandler(c *gin.Context) {
	var req struct {
		IdToken string `json:"id_token"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload, err := idtoken.Validate(context.Background(), req.IdToken, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	email := payload.Claims["email"].(string)
	name := payload.Claims["name"].(string)

	user, err := h.UC.FindOrCreateUser(email, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}
