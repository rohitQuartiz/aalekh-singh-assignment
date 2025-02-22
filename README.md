# Blog Post API

## Overview
Blog Post API is a RESTful API designed to manage blog posts. It provides CRUD (Create, Read, Update, Delete) functionality, allowing users to create, update, fetch, and delete blog posts. This API is built using Go and the Fiber framework and is hosted on Render.

### Hosted URL
The API is deployed on Render and can be accessed at:
**[https://blog-post-api-chmy.onrender.com/api/blogs/](https://blog-post-api-chmy.onrender.com/api/blogs/)**

### Swagger URL
The Swagger documentation for the API can be accessed at:
**[https://blog-post-api-chmy.onrender.com/swagger/index.html#/](https://blog-post-api-chmy.onrender.com/swagger/index.html#/)**

## Features
- **Create Blog Posts**: Add new blog posts with a title, body, image, and description.
- **Retrieve Blog Posts**: Fetch all blogs or a specific blog by ID.
- **Update Blog Posts**: Modify the title, body, and other attributes of a blog.
- **Delete Blog Posts**: Remove a blog post by ID.
- **Database Integration**: Uses PostgreSQL for storing blog data.
- **Validation**: Ensures proper request formats and data integrity.
- **Error Handling**: Provides detailed error messages for invalid operations.

## Technologies Used
- **Go**: Primary language used for backend development.
- **Fiber**: A fast, lightweight web framework for Go.
- **GORM**: ORM library used for database interactions.
- **SQLite**: Relational database for storing blog posts.
- **JSON**: Used for data exchange between client and server.
- **Render**: Cloud platform used for hosting the API.

## Libraries
- **Fiber v3**: Web framework ([github.com/gofiber/fiber/v3](https://github.com/gofiber/fiber))
- **GORM**: ORM for Go ([gorm.io/gorm](https://gorm.io/gorm))
- **PostgreSQL Driver**: ([gorm.io/driver/postgres](https://gorm.io/driver/postgres))
- **Testify**: Testing framework ([github.com/stretchr/testify](https://github.com/stretchr/testify))

## Endpoints
### 1. Create a Blog Post
**POST** `/api/blogs/`
```json
{
    "title": "My First Blog",
    "body": {
        "name": "Author Name",
        "image": "blog-image.jpg",
        "description": "This is a sample blog post."
    }
}
```

### 2. Get All Blogs
**GET** `/api/blogs/`

### 3. Get a Blog by ID
**GET** `/api/blogs/{id}`

### 4. Update a Blog Post
**PUT** `/api/blogs/{id}`
```json
{
    "title": "Updated Blog Title",
    "body": {
        "name": "Updated Name",
        "image": "updated-image.jpg",
        "description": "Updated description."
    }
}
```

### 5. Delete a Blog Post
**DELETE** `/api/blogs/{id}`

## Installation & Setup
1. Clone the repository:
   ```sh
   git clone https://github.com/your-repo/blog-post-api.git
   cd blog-post-api
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the application:
   ```sh
   go run main.go
   ```

## Testing
Run unit tests using:
```sh
go test ./tests
```

## Deployment
The API is hosted on Render and automatically updates from the repository upon new commits.

## License
This project is open-source.

---
**Author**: Aalekh Singh 


