package routes

import (
	"be-ruang-warga/config"
	ruangRiungDelivery "be-ruang-warga/internal/ruangriung/delivery"
	ruangRiungUsecase "be-ruang-warga/internal/ruangriung/usecase"
	userDelivery "be-ruang-warga/internal/user/delivery"
	userUsecase "be-ruang-warga/internal/user/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	db := config.DB
	authClient := config.AuthClient
	api := router.Group("/api")

	rrUC := ruangRiungUsecase.NewRuangRiungUsecase(db)
	uUC := userUsecase.NewUserUsecase(db)

	userDelivery.NewUserHandler(api, uUC, authClient)
	ruangRiungDelivery.NewRuangRiungHandler(api, rrUC)
}