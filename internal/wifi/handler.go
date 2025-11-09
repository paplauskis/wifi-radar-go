package wifi

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetReviews(context *gin.Context) {
	id := context.Query("id")

	parsedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Invalid ID format.")
	}

	wifiReviews, err := QueryWifiReview(parsedId, h.db)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	}

	context.JSON(http.StatusOK, wifiReviews)
}

func (h *Handler) CreateReview(context *gin.Context) {

}

func (h *Handler) GetPasswords(context *gin.Context) {

}

func (h *Handler) AddPassword(context *gin.Context) {

}
