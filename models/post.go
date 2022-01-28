package models

type Post struct {
	ID      uint
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Author  string `json:"author" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostInput struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
