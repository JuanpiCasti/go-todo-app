package customerror

import "strconv"

type TodoNotFoundError struct {
	ID int
}

func (e *TodoNotFoundError) Error() string {
	return "Todo with the given ID not found: " + strconv.Itoa(e.ID)
}

func NewTodoNotFoundError(ID int) error {
	return &TodoNotFoundError{ID}
}
