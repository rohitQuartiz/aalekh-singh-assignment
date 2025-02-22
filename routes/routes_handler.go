package routes

import (
	"log"
	"strconv"

	"github.com/aalekh12/Blog-Post-Api/database"
	"github.com/aalekh12/Blog-Post-Api/models"
	"github.com/gofiber/fiber/v2"
)

// CreateBlog creates a new blog post
// @Summary Create a new blog post
// @Description Create a new blog post with title, description, and body
// @Tags Blogs
// @Accept json
// @Produce json
// @Param blog body models.Blog true "Blog Post"
// @Success 201 {object} models.Blog
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/blogs/ [post]
func CreateBlog(c *fiber.Ctx) error {
	blog := new(models.Blog)

	// Parse request body
	if err := c.BodyParser(blog); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	blog, err := database.CreateBlog(blog)

	// Save to DB
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create blog"})
	}

	return c.Status(fiber.StatusCreated).JSON(blog)
}

// GetBlogByID retrieves a single blog post by ID
// @Summary Get a blog post by ID
// @Description Retrieve a blog post using its ID
// @Tags Blogs
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200 {object} models.Blog
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/blogs/{id} [get]
func GetBlogByID(c *fiber.Ctx) error {
	id := c.Params("id")
	blogID, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid blog ID"})
	}

	var blog *models.Blog
	blog, err = database.GetBlogByID(uint16(blogID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Blog not found"})
	}

	return c.JSON(blog)
}

// GetAllBlogs retrieves all blog posts
// @Summary Get all blog posts
// @Description Retrieve a list of all blog posts
// @Tags Blogs
// @Produce json
// @Success 200 {array} models.Blog
// @Failure 404 {object} map[string]string
// @Router /api/blogs/ [get]
func GetAllBlogs(c *fiber.Ctx) error {
	var blogs []models.Blog
	blogs, err := database.GetAllBlogs()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Blogs not found"})
	}

	return c.JSON(blogs)
}

// UpdateBlog updates an existing blog post
// @Summary Update a blog post
// @Description Update a blog post's title, description, or body
// @Tags Blogs
// @Accept json
// @Produce json
// @Param id path int true "Blog ID"
// @Param blog body models.Blog true "Updated Blog Data"
// @Success 200 {object} models.Blog
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/blogs/{id} [put]
func UpdateBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Received ID:", id)

	blogID, err := strconv.ParseUint(id, 10, 16)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid blog ID"})
	}

	var blog models.Blog // Initialize the blog variable

	// Parse request body into the blog struct
	if err := c.BodyParser(&blog); err != nil {
		log.Println("Req JSON Error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	log.Println("Parsed blog data:", blog)

	updatedBlog, err := database.UpdateBlog(uint16(blogID), &blog)
	if err != nil {
		log.Println("DB Error:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Blog not found"})
	}

	return c.JSON(updatedBlog)
}

// DeleteBlog deletes a blog post
// @Summary Delete a blog post
// @Description Delete a blog post by ID
// @Tags Blogs
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/blogs/{id} [delete]
func DeleteBlog(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := database.DB.Delete(&models.Blog{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete blog"})
	}

	return c.JSON(fiber.Map{"message": "Blog deleted successfully"})
}

// RegisterRoutes registers all blog routes
func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/blogs")
	api.Post("/", CreateBlog)
	api.Get("/", GetAllBlogs)
	api.Get("/:id", GetBlogByID)
	api.Put("/:id", UpdateBlog)
	api.Delete("/:id", DeleteBlog)
}
