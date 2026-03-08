package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Простой эндпоинт для проверки
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Заглушка для создания тикета (пока просто возвращаем ID)
	r.POST("/api/tickets", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"ticket_id": 123})
	})

	log.Println("API Service starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
