package main

import (
	"encoding/json"
	"net/http"
	"os"
	"wayfarer/database"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string
	Password string
}

func addUser(ctx *gin.Context) {
	body := User{}
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = database.Db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", body.Username, body.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}

func main() {
	r := gin.Default()
	r.POST("/add-user", addUser)
	database.ConnectDatabase()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
