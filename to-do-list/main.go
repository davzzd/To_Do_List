package main

import (
	"net/http"
	"sync"

	"to-do-list/components"

	"github.com/a-h/templ"
)

type TodoList struct {
	items []components.Todo
	mu    sync.RWMutex
}

func (t *TodoList) Add(item string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.items = append(t.items, components.Todo{Text: item})
}

func (t *TodoList) Delete(item string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for i, todo := range t.items {
		if todo.Text == item {
			t.items = append(t.items[:i], t.items[i+1:]...)
			break
		}
	}
}

func (t *TodoList) GetAll() []components.Todo {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.items
}

func main() {
	todos := &TodoList{}

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		templ.Handler(components.Home(todos.GetAll())).ServeHTTP(w, r)
	})

	// Add todo
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		item := r.FormValue("item")
		if item != "" {
			todos.Add(item)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// Delete todo
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		item := r.FormValue("item")
		if item != "" {
			todos.Delete(item)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.ListenAndServe(":8080", nil)
}
