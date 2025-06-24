package delivery

import (
	"be-ruang-warga/internal/ruangriung/domain"
	"be-ruang-warga/internal/ruangriung/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RuangRiungHandler struct {
	UC usecase.RuangRiungUsecase
}

func NewRuangRiungHandler(router *gin.RouterGroup, uc usecase.RuangRiungUsecase) {
	h := &RuangRiungHandler{UC: uc}

	router.GET("/ruangriung", h.GetAll)
	router.POST("/ruangriung", h.Create)
	router.PATCH("/ruangriung", h.Update)
}

func (h *RuangRiungHandler) GetAll(c *gin.Context) {
	data, err := h.UC.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *RuangRiungHandler) Create(c *gin.Context) {
	var input domain.RuangRiung

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UC.Create(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Acara ruang riung berhasil di tambahkan"})
}

func (h *RuangRiungHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var update domain.RuangRiung
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UC.Update(id, &update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Acara ruang riung berhasil di perbarui"})
}
