package controllers

import (
	"database/sql"
	"net/http"
	"quiz3/config"
	"quiz3/models"
	"time"

	"github.com/gin-gonic/gin"
)


func GetCategories(c *gin.Context) {
    var categories []models.Category

    rows, err := config.DB.Query("SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query categories"})
        return
    }
    defer rows.Close()

    for rows.Next() {
        var category models.Category
        if err := rows.Scan(
            &category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy,
            &category.ModifiedAt, &category.ModifiedBy,
        ); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan category"})
            return
        }
        categories = append(categories, category)
    }

    c.JSON(http.StatusOK, categories)
}


func GetCategoryByID(c *gin.Context) {
    id := c.Param("id")
    var category models.Category

    query := `SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id=$1`
    err := config.DB.QueryRow(query, id).Scan(
        &category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy,
        &category.ModifiedAt, &category.ModifiedBy,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query category"})
        }
        return
    }

    c.JSON(http.StatusOK, category)
}


func AddCategory(c *gin.Context) {
    var category models.Category

    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

   
    category.CreatedAt = time.Now()
    category.ModifiedAt = time.Now()

    query := `
        INSERT INTO categories (name, created_at, created_by, modified_at, modified_by)
        VALUES ($1, $2, $3, $4, $5) RETURNING id
    `

    err := config.DB.QueryRow(query,
        category.Name, category.CreatedAt, category.CreatedBy,
        category.ModifiedAt, category.ModifiedBy,
    ).Scan(&category.ID)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add category"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category added successfully", "category_id": category.ID})
}


func DeleteCategory(c *gin.Context) {
    id := c.Param("id")

 
    var exists bool
    err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM categories WHERE id=$1)", id).Scan(&exists)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check category existence"})
        return
    }

    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }

   
    query := `DELETE FROM categories WHERE id=$1`
    _, err = config.DB.Exec(query, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}


func GetBooksByCategory(c *gin.Context) {
    categoryID := c.Param("id")
    var books []models.Book

    query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by
              FROM books WHERE category_id=$1`
    rows, err := config.DB.Query(query, categoryID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query books"})
        return
    }
    defer rows.Close()

    for rows.Next() {
        var book models.Book
        if err := rows.Scan(
            &book.ID, &book.Title, &book.Description, &book.ImageURL,
            &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness,
            &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy,
        ); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan book"})
            return
        }
        books = append(books, book)
    }

    if len(books) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No books found for this category"})
        return
    }

    c.JSON(http.StatusOK, books)
}
