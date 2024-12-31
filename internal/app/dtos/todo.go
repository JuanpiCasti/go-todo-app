package dtos

type TodoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoCreateRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
