package controllers

import (
	"database/sql"
	"net/http"
	"quiz3/config"
	"quiz3/models"
	"time"

	"github.com/gin-gonic/gin"
)


func AddBook(c *gin.Context) {
    var book models.Book

   
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
        return
    }

    if book.TotalPage > 100 {
        book.Thickness = "tebal"
    } else {
        book.Thickness = "tipis"
    }

   
    book.CreatedAt = time.Now()
    book.ModifiedAt = time.Now()

    
    query := `
        INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id
    `

    // Eksekusi query dan handle error jika ada
    err := config.DB.QueryRow(query,
        book.Title, book.Description, book.ImageURL, book.ReleaseYear,
        book.Price, book.TotalPage, book.Thickness, book.CategoryID,
        book.CreatedAt, book.CreatedBy, book.ModifiedAt, book.ModifiedBy,
    ).Scan(&book.ID)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Book added successfully", "book_id": book.ID})
}


func GetBooks(c *gin.Context) {
    var books []models.Book

    rows, err := config.DB.Query("SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books")
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

    c.JSON(http.StatusOK, books)
}


func GetBookByID(c *gin.Context) {
    id := c.Param("id")
    var book models.Book

    query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books WHERE id=$1`
    err := config.DB.QueryRow(query, id).Scan(
        &book.ID, &book.Title, &book.Description, &book.ImageURL,
        &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness,
        &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query book"})
        }
        return
    }

    c.JSON(http.StatusOK, book)
}


func DeleteBook(c *gin.Context) {
    id := c.Param("id")

    
    var exists bool
    err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE id=$1)", id).Scan(&exists)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check book existence"})
        return
    }

    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

 
    query := `DELETE FROM books WHERE id=$1`
    _, err = config.DB.Exec(query, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
