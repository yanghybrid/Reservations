package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// Reservation struct
type Reservation struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Guests   int       `json:"guests"`
	DateTime time.Time `json:"dateTime"`
}

// Store reservations
var reservations = []Reservation{}
var nextID = 1

func main() {
	r := gin.Default()

	// ✅ Allow CORS for frontend requests (React)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3001", "http://10.0.0.139:3001"}, // Add allowed frontend URLs
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// API Routes
	r.GET("/reservations", func(c *gin.Context) {
		c.JSON(http.StatusOK, reservations)
	})

	r.POST("/reservations", func(c *gin.Context) {
		var newReservation Reservation
		if err := c.ShouldBindJSON(&newReservation); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newReservation.ID = nextID
		nextID++
		newReservation.DateTime = time.Now().Add(time.Hour * 24)
		reservations = append(reservations, newReservation)

		c.JSON(http.StatusCreated, gin.H{"message": "Reservation successful", "reservation": newReservation})
	})

	fmt.Println("Server running on :8080")
	r.Run(":8080") // Start server on port 8080
}
