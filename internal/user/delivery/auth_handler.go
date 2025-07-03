package delivery

import (
	"be-ruang-warga/internal/user/usecase"
	"be-ruang-warga/utils"
	"context"
	"fmt"
	"net/http"

	firebaseAuth "firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UC         usecase.UserUsecase
	AuthClient *firebaseAuth.Client
}

func NewUserHandler(router *gin.RouterGroup, uc usecase.UserUsecase, authClient *firebaseAuth.Client) {
	h := &UserHandler{UC: uc, AuthClient: authClient}

	router.POST("/auth/google", h.GoogleAuthHandler)
}

func (h *UserHandler) GoogleAuthHandler(c *gin.Context) {
	fmt.Println("GoogleAuthHandler called.")
	var req struct {
		IdToken string `json:"id_token"`
	}

	if err := c.BindJSON(&req); err != nil {
		fmt.Println("Error binding JSON:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(req.IdToken) > 30 {
		fmt.Println("Received ID Token (first 30 chars):", req.IdToken[:30]+"...")
	} else {
		fmt.Println("Received ID Token:", req.IdToken)
	}

	token, err := h.AuthClient.VerifyIDToken(context.Background(), req.IdToken)
	if err != nil {
		fmt.Println("Firebase ID Token verification failed:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired ID token."})
		return
	}

	email := token.Claims["email"].(string)
	name := token.Claims["name"].(string)

	fmt.Println("Firebase ID Token verified successfully for email:", email)

	user, err := h.UC.FindOrCreateUser(email, name)
	if err != nil {
		fmt.Println("Error finding or creating user:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jwtToken, err := utils.GenerateJWT(*user)
	if err != nil {
		fmt.Println("Error generating JWT:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Successfully processed Google Auth. User:", user.Email, "Generated JWT:", jwtToken[:20]+"...")

	c.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
		"user":  user,
	})
}
