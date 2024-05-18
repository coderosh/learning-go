package handlers

import (
	"net/http"
	"time"

	"gotodo/database"
	"gotodo/views"
)

type TodoHandler struct{}

func (h *TodoHandler) GetTodos() func(w http.ResponseWriter, r *http.Request) {
	tpl := views.NewTemplate("todos.html")

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, map[string]any{
			"Todos": database.GlobalDB.GetAll(),
		})
	}
}

func (h *TodoHandler) CreateTodo() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		todo := r.Form.Get("todo")

		curDate := time.Now().String()

		database.GlobalDB.Set(curDate, todo)

		http.Redirect(w, r, "/todos", http.StatusSeeOther)
	}
}

func (h *TodoHandler) DeleteTodo() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		database.GlobalDB.Delete(id)

		http.Redirect(w, r, "/todos", http.StatusSeeOther)
	}
}
