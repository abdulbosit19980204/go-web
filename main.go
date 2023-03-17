package main

import (
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Product 1", Price: 100},
	{ID: 2, Name: "Product 2", Price: 200},
	{ID: 3, Name: "Product 3", Price: 300},
}

type Order struct {
	Products []Product `json:"products"`
	Total    int       `json:"total"`
}

func main() {

	// Initialize Gin router
	r := gin.Default()

	// Serve static assets
	r.Static("/assets", "./assets")

	// Define API endpoint
	r.GET("/api/products", func(c *gin.Context) {
		// Return a JSON response with list of products
		products := []string{"Product 1", "Product 2", "Product 3"}
		c.JSON(http.StatusOK, gin.H{"products": products})
	})

	router := gin.Default()

	// GET all products
	router.GET("/products", func(c *gin.Context) {
		c.JSON(200, products)
	})

	// POST add product to cart
	router.POST("/cart/:productID", func(c *gin.Context) {
		productID := c.Param("productID")

		for _, product := range products {
			if strconv.Itoa(product.ID) == productID {
				// TODO: add product to cart
				c.JSON(200, gin.H{"message": "Product added to cart"})
				return
			}
		}

		c.JSON(404, gin.H{"error": "Product not found"})
	})

	// GET cart
	router.GET("/cart", func(c *gin.Context) {
		// TODO: get cart contents
		cart := Order{Products: []Product{}, Total: 0}
		c.JSON(200, cart)
	})

	// POST checkout
	router.POST("/checkout", func(c *gin.Context) {
		// TODO: process order and clear cart
		order := Order{Products: []Product{}, Total: 0}
		c.JSON(200, order)
	})

	router.Run(":8080")
}
