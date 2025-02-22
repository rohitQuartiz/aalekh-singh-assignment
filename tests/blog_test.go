package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aalekh12/Blog-Post-Api/database"
	"github.com/aalekh12/Blog-Post-Api/models"
	"github.com/aalekh12/Blog-Post-Api/routes"
	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup a test database in memory
func setupTestDB() {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&models.Blog{})
	database.DB = db
}

// Setup Fiber app with test routes
func setupTestApp() *fiber.App {
	app := fiber.New()
	routes.RegisterRoutes(app)
	return app
}

// Test creating a blog post
func TestCreateBlog(t *testing.T) {
	setupTestDB()
	app := setupTestApp()

	blogData := models.Blog{
		Title: "Test Blog",
		Body: models.BlogBody{
			Name:        "Test Name",
			Image:       "test.jpg",
			Description: "Test Description",
		},
	}

	body, _ := json.Marshal(blogData)
	req := httptest.NewRequest(http.MethodPost, "/api/blogs/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

// Test getting all blogs
func TestGetAllBlogs(t *testing.T) {
	setupTestDB()
	app := setupTestApp()

	// Add sample data
	database.DB.Create(&models.Blog{
		Title: "Blog 1",
		Body: models.BlogBody{
			Name:        "Test 1",
			Image:       "img1.jpg",
			Description: "Description 1",
		},
	})
	database.DB.Create(&models.Blog{
		Title: "Blog 2",
		Body: models.BlogBody{
			Name:        "Test 2",
			Image:       "img2.jpg",
			Description: "Description 2",
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/api/blogs/", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// Test getting a blog by ID
func TestGetBlogByID(t *testing.T) {
	setupTestDB()
	app := setupTestApp()

	blog := models.Blog{
		Title: "Test Blog",
		Body: models.BlogBody{
			Name:        "Test Name",
			Image:       "test.jpg",
			Description: "Test Description",
		},
	}
	database.DB.Create(&blog)

	req := httptest.NewRequest(http.MethodGet, "/api/blogs/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// Test updating a blog
func TestUpdateBlog(t *testing.T) {
	setupTestDB()
	app := setupTestApp()

	blog := models.Blog{
		Title: "Original Blog",
		Body: models.BlogBody{
			Name:        "Original Name",
			Image:       "orig.jpg",
			Description: "Original Description",
		},
	}
	database.DB.Create(&blog)

	updatedData := models.Blog{
		Title: "Updated Blog",
		Body: models.BlogBody{
			Name:        "Updated Name",
			Image:       "updated.jpg",
			Description: "Updated Description",
		},
	}
	body, _ := json.Marshal(updatedData)

	req := httptest.NewRequest(http.MethodPut, "/api/blogs/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// Test deleting a blog
func TestDeleteBlog(t *testing.T) {
	setupTestDB()
	app := setupTestApp()

	blog := models.Blog{
		Title: "Blog to Delete",
		Body: models.BlogBody{
			Name:        "Delete Name",
			Image:       "delete.jpg",
			Description: "To be deleted",
		},
	}
	database.DB.Create(&blog)

	req := httptest.NewRequest(http.MethodDelete, "/api/blogs/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
