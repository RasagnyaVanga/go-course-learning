package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewBlogDB initializes and returns a new database connection for managing blog posts.
// It loads environment variables from the .env file to configure the database connection,
// builds a MySQL DSN (Data Source Name), and connects using GORM.
//
// The function performs automatic migration for the BlogPost model to ensure the database
// schema is up to date. It returns a pointer to a dbBlog instance if successful, or an error
// if any step fails (e.g., missing environment variables, connection failure, or migration error).
func NewBlogDB() (*dbBlog, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || password == "" || host == "" || port == "" || name == "" {
		return nil, errors.New("database environment variables not set")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	if err := db.AutoMigrate(&BlogPost{}); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return &dbBlog{db: db}, nil
}

// CreateBlog creates a new blog post with the given title, author, and content.
// It inserts the post into the database and returns an error if the operation fails.
func (dbBlog *dbBlog) CreateBlog(title string, author string, content string) error {
	blogpost := BlogPost{Title: title, Author: author, Content: content}
	res := (*dbBlog).db.Create(&blogpost)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// ReadAllBlogs retrieves all blog posts from the database.
// It returns a slice of BlogPost structs and an error if the query fails.
func (blog dbBlog) ReadAllBlogs() ([]BlogPost, error) {
	var blogs []BlogPost
	res := blog.db.Model(&BlogPost{}).Find(&blogs)
	if res.Error != nil {
		return nil, res.Error
	}
	return blogs, nil
}

// UpdateTable updates the title and content of a blog post identified by its ID.
// It returns an error if the update fails or if no blog post is found with the given ID.
func (blog *dbBlog) UpdateTable(id uint, newtitle string, content string) error {
	res := blog.db.Model(&BlogPost{}).
		Where("id =?", id).
		Updates(map[string]interface{}{
			"Title":   newtitle,
			"Content": content,
		})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("no blog post found with the given ID")
	}
	return nil
}

// DeleteBlog removes a blog post from the database by its title.
// It returns an error if the deletion fails.
func (blog *dbBlog) DeleteBlog(title string) error {
	res := blog.db.Where("title=?", title).Delete(&BlogPost{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// SearchBlog finds blog posts matching the given author or title.
// It returns a slice of BlogPost structs and an error if the query fails.
func (blog dbBlog) SearchBlog(author string, title string) ([]BlogPost, error) {
	var posts []BlogPost
	res := blog.db.Where("author = ?", author).Or("title = ?", title).Find(&posts)
	if res.Error != nil {
		return nil, res.Error
	}
	return posts, nil
}
