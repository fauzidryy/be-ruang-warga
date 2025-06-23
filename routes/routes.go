package routes

import (
	"be-ruang-warga/config"
	ruangriungDelivery "be-ruang-warga/internal/ruangriung/delivery"
	ruangriungUsecase "be-ruang-warga/internal/ruangriung/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	db := config.DB
	api := router.Group("/api")

	rrUC := ruangriungUsecase.NewRuangRiungUsecase(db)
	ruangriungDelivery.NewRuangRiungHandler(api, rrUC)
}
