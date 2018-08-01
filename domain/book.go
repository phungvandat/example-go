package domain

// Book describe book in system
type Book struct {
	Model
	Name        string `json:"name"`
	CategoryID  UUID   `json:"category_id"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
