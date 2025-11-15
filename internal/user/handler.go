package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (uh *UserHandler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	//password hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id int
	err = uh.db.QueryRow(
		"INSERT INTO public.user (username, password, created_at) VALUES ($1, $2, $3) RETURNING id",
		input.Username, string(hashedPassword), time.Now(),
	).Scan(&id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "username": input.Username})
}

func (uh *UserHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &User{}
	err := uh.db.QueryRow(
		`SELECT id, username, password FROM "user" WHERE username = $1`,
		input.Username,
	).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   fmt.Sprint(user.ID),
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (uh *UserHandler) AddFavorite(c *gin.Context) {
	ID := c.Param("id")
	var input struct {
		WifiID int `json:"wifi_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	_, err := uh.db.Exec("INSERT INTO user_favorite_wifi (user_id, wifi_id) VALUES ($1, $2)", ID, input.WifiID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to favorite"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "favorite added"})
}

func (uh *UserHandler) GetFavorite(c *gin.Context) {
	ID := c.Param("id")

	rows, err := uh.db.Query(
		"SELECT wifi_id FROM user_favorite_wifi WHERE user_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get favorite"})
		return
	}
	defer rows.Close()

	var favorite []int
	for rows.Next() {
		var wifiID int
		rows.Scan(&wifiID)
		favorite = append(favorite, wifiID)
	}

	c.JSON(http.StatusOK, gin.H{"favorite": favorite})
}

func (uh *UserHandler) DeleteFavorite(c *gin.Context) {
	ID := c.Param("id")
	WifiID := c.Param("wifi_id")

	if ID == "" || WifiID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id and wifi_id are required"})
		return
	}

	result, err := uh.db.Exec(
		"DELETE FROM user_favorite_wifi WHERE user_id=$1 AND wifi_id=$2",
		ID, WifiID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete favorite"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "favorite not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "favorite deleted"})
}
