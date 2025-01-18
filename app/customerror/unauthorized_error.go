package customerror

type unauthorizedTodoError struct {
	ID string
}

func (e *unauthorizedTodoError) Error() string {
	return "Can not access a todo with the ID: " + e.ID
}
