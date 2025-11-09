package wifi

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetReviews(context *gin.Context) {

}

func (h *Handler) CreateReview(context *gin.Context) {

}

func (h *Handler) GetPasswords(context *gin.Context) {

}

func (h *Handler) AddPassword(context *gin.Context) {

}
