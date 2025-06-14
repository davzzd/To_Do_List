package main

import (
	"net/http"
	"strings"
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
	t.items = append(t.items, components.Todo{
		Text:      item,
		Completed: false,
	})
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

func (t *TodoList) Toggle(item string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for i := range t.items {
		if t.items[i].Text == item {
			t.items[i].Completed = !t.items[i].Completed
			break
		}
	}
}

func (t *TodoList) GetAll() []components.Todo {
	t.mu.RLock()
	defer t.mu.RUnlock()
	items := make([]components.Todo, len(t.items))
	copy(items, t.items)
	return items
}

func main() {
	todos := &TodoList{}

	//Serves the static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//API routes
	api := http.NewServeMux()

	//handles /todos endpoint for both GET and POST
	api.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			templ.Handler(components.Home(todos.GetAll())).ServeHTTP(w, r)
		case http.MethodPost:
			item := r.FormValue("item")
			if item != "" {
				todos.Add(item)
			}
			http.Redirect(w, r, "/api/todos", http.StatusSeeOther)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	//handles /todos/{item} endpoints
	api.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/todos/")
		if path == "" {
			http.NotFound(w, r)
			return
		}

		//handles toggle endpoint
		if strings.HasSuffix(path, "/toggle") {
			item := strings.TrimSuffix(path, "/toggle")
			if item != "" {
				todos.Toggle(item)
			}
		} else {
			//delete endpoint
			todos.Delete(path)
		}

		http.Redirect(w, r, "/api/todos", http.StatusSeeOther)
	})

	http.Handle("/api/", http.StripPrefix("/api", api))

	//redirects to home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, "/api/todos", http.StatusSeeOther)
	})

	http.ListenAndServe(":8080", nil)
}
