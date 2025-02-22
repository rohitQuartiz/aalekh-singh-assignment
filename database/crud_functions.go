package database

import (
	"log"

	"github.com/aalekh12/Blog-Post-Api/models"
)

func CreateBlog(blog *models.Blog) (*models.Blog, error) {
	if err := DB.Create(blog).Error; err != nil {
		log.Printf("Error creating blog: %v", err)
		return nil, err
	}
	return blog, nil
}

func GetBlogByID(id uint16) (*models.Blog, error) {
	var blog models.Blog
	if err := DB.First(&blog, id).Error; err != nil {
		log.Printf("Error fetching blog with ID %d: %v", id, err)
		return nil, err
	}
	return &blog, nil
}

func GetAllBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	if err := DB.Find(&blogs).Error; err != nil {
		log.Printf("Error fetching blogs: %v", err)
		return nil, err
	}
	return blogs, nil
}

func UpdateBlog(id uint16, updatedBlog *models.Blog) (*models.Blog, error) {
	var blog models.Blog
	if err := DB.First(&blog, id).Error; err != nil {
		log.Printf("Blog with ID %d not found: %v", id, err)
		return nil, err
	}

	blog.Title = updatedBlog.Title
	blog.Body = updatedBlog.Body

	if err := DB.Save(&blog).Error; err != nil {
		log.Printf("Error updating blog with ID %d: %v", id, err)
		return nil, err
	}
	return &blog, nil
}

func DeleteBlog(id uint16) error {
	var blog models.Blog
	if err := DB.First(&blog, id).Error; err != nil {
		log.Printf("Blog with ID %d not found: %v", id, err)
		return err
	}

	if err := DB.Delete(&blog).Error; err != nil {
		log.Printf("Error deleting blog with ID %d: %v", id, err)
		return err
	}
	return nil
}
