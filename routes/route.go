package routes

import (
	"quiz3/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

   
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Welcome to the Book Management API!",
        })
    })

    // Routes untuk buku
    router.GET("/api/books", controllers.GetBooks)
    router.POST("/api/books", controllers.AddBook)
    router.GET("/api/books/:id", controllers.GetBookByID)
    router.DELETE("/api/books/:id", controllers.DeleteBook)

    // Routes untuk kategori
    router.GET("/api/categories", controllers.GetCategories)
    router.POST("/api/categories", controllers.AddCategory)
    router.GET("/api/categories/:id", controllers.GetCategoryByID)
    router.DELETE("/api/categories/:id", controllers.DeleteCategory)
    router.GET("/api/categories/:id/books", controllers.GetBooksByCategory)

    return router
}

