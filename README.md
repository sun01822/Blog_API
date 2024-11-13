
# 🌐 Blog API - Golang Echo Framework 📖

Welcome to the **Blog API**! Built with the Golang Echo framework, this API facilitates CRUD operations for blogs, users, and comments, allowing users to create, retrieve, update, and delete blog posts with ease. Designed for managing blog, user, and comment data efficiently, the Blog API is ready for seamless integration into blogging platforms and content management systems.


---

## 📚 Table of Contents

- [Getting Started](#-getting-started)
- [API Endpoints](#-api-endpoints)
  - [User Endpoints](#-user-endpoints)
  - [Blog Endpoints](#-blog-endpoints)
- [Schema Definitions](#-schema-definitions)
- [License](#-license)

---

## 🚀 Getting Started

### Prerequisites

- **Golang** - Make sure Golang is installed on your system.
- **Echo Framework** - This project is built on Echo. You can install it by running:
  ```bash
  go get -u github.com/labstack/echo/v4
  ```

### Installation

1. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/sun01822/Blog_API.git
   ```
2. Move to the project directory:
   ```bash
   cd Blog_API
   ```
3. Create app.env file, copy, paste and add values to it:
   ```bash DBUSER=
      DBPASS=
      DBIP=
      DBNAME=
      PORT=
      JWTSECRET=
   ```
4. Start the API server:
   ```bash
   go run main.go
   ```

The Blog API server will run at `http://localhost:8080`.

---

## 📖 API Endpoints

### 🔹 User Endpoints

- **User Signup** - `POST /user/create`
  - **Request Body:** Should follow the `SignUpRequest` schema.
  - **Response:** Confirmation of user creation or error.

- **User Login** - `POST /user/login`
  - **Request Body:** Should follow the `LoginRequest` schema.
  - **Response:** A JWT token upon successful login.

- **Delete a User** - `DELETE /user/delete`
  - Requires Bearer token for authorization.
  - **Query parameter:** `user_id` (string) - ID of the user.
  - **Response:** Confirmation of user deletion or error.

- **Update a User** - `UPDATE /user/delete`
  - Requires Bearer token for authorization.
   - **Query parameter:** `user_id` (string) - ID of the user.
  - **Response:** Confirmation of user update or error.

- **Get All Users** - `GET /user/getAll`
  - Optional query parameters for pagination: `offset` and `limit`.
  - **Response:** Returns all users with pagination.

- **Get User by ID** - `GET /user/get`
  - **Query parameter:** `user_id` (string) - ID of the user.
  - **Response:** Returns all users with pagination.

<br/>

### 🔹 Blog Endpoints

- **Create a Blog Post** - `POST /blog/create`
  - Requires Bearer token for authorization.
  - **Request Body:** Must follow the `BlogPostRequest` schema.
  - **Response:** A message confirming creation or an error.

- **Get a Blog Post** - `GET /blog/get`
  - **Query Parameter:** `blog_id` (string) - ID of the blog post to fetch.
  - **Response:** Returns blog details or an error message.

- **Get All Blog Posts by User** - `GET /blog/get/user`
  - **Query Parameter:** `user_id` (string) - ID of the user.
  - **Response:** Returns all blog posts by the user.

- **Update a Blog Post** - `PUT /blog/update`
  - Requires Bearer token for authorization.
  - **Query Parameter:** `blog_id` (string) - ID of the blog post to update.
  - **Request Body:** Follows the `UpdateBlogPostRequest` schema.
  - **Response:** Update confirmation or error.

- **Delete a Blog Post** - `DELETE /blog/delete`
  - Requires Bearer token for authorization.
  - **Query Parameter:** `blog_id` (string) - ID of the blog post to delete.
  - **Response:** Deletion confirmation or an error.

- **Add, Remove, and Update Reaction on a Blog Post** - `POST /blog/reaction`
  - Requires Bearer token for authorization.
  - **Query Parameter:** `blog_id` (string) - ID of the blog post to add reaction and
                         `reaction_id` (uint) - ID of the reaction {1: like, 2: love, 3: care, 4: haha, 5: wow, 6: sad, 7: angry}
  - **Response:** Reaction confirmation with BlogResp or an error.

- **Add Comment on a Blog Post** - `POST /blog/comment`
  - Requires Bearer token for authorization.
  - **Query Parameter:** `blog_id` (string) - ID of the blog post to add comment
  - **Request Body:** Should follow the `Comment` schema.
  - **Response:** Comment confirmation with BlogResp or an error.
 
- **Get Comments on a Blog Post** - `POST /blog/comments`
  - Requires Bearer token for authorization.
  - **Query Parameter:** `blog_id` (string) - ID of the blog post and 
                         `comment_ids` (string) - ID's of comment
  - **Response:** Comment confirmation with CommentResp or an error.

- **Update Comment on a Blog Post** - `PUT /blog/comment`
  - Requires Bearer token for authorization.
  - **Query Parameter:** `blog_id` (string) - ID of the blog post and 
                         `comment_id` (string) - ID of the comment
  - **Response:** Comment confirmation with BlogResp or an error.

- **Delete Comment on a Blog Post** - `DELETE /blog/comment`
  - Requires Bearer token for authorization.
  - **Query Parameter:** `blog_id` (string) - ID of the blog post and 
                         `comment_id` (string) - ID of the comment
  - **Response:** Comment confirmation or an error.


---

## 📦 Schema Definitions

### BlogPostRequest
```json
{
  "category": "string",
  "content_text": "string",
  "description": "string",
  "is_published": "boolean",
  "photo_url": "string",
  "title": "string"
}
```

### SignUpRequest
```json
{
  "email": "string",
  "password": "string",
  "first_name": "string",
  "last_name": "string",
  "phone": "string"
}
```

### LoginRequest
```json
{
  "email": "string",
  "password": "string"
}
```

### Comment
```json
{
  "content": "string"
}
```

### CommentResp
```json
[
  {
    "blog_post_id": "string",
    "content": "string",
    "id": "string",
    "user_id": "string"
  }
]
```

### BlogResp
```json
{
  "category": "string",
  "comments": [
    {
      "blog_post_id": "string",
      "content": "string",
      "id": "string",
      "user_id": "string"
    }
  ],
  "comments_count": 0,
  "content_text": "string",
  "created_at": "string",
  "deleted_at": "string",
  "description": "string",
  "id": "string",
  "is_published": true,
  "photo_url": "string",
  "published_at": "string",
  "reactions": [
    {
      "blog_post_id": "string",
      "id": "string",
      "type": 0,
      "user_id": "string"
    }
  ],
  "reactions_count": 0,
  "title": "string",
  "updated_at": "string",
  "user_id": "string",
  "views": 0
}```



---

## 📝 License

This project is licensed under the Apache 2.0 License. See the [LICENSE](http://www.apache.org/licenses/LICENSE-2.0.html) file for details.

For any issues or questions, please contact [API Support](mailto:support@swagger.io).

---

**Thank you for using Blog API! Happy Coding!** 🎉
