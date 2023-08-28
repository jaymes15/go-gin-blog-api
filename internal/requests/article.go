package requests

// Binding from JSON
type CreateArticle struct {
	Title   string `json:"title" binding:"required,min=2"`
	Content string `json:"content" binding:"required"`
}
