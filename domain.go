package main

//Todo é a struct da lista de tarefas
type Todo struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
