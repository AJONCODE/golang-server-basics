package route

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// HomeRoute returns home route and prints "home route!" on console and browser
func HomeRoute() func(w http.ResponseWriter, r *http.Request) {
	home := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("home route!")

		fmt.Fprint(w, "home route!")
	}

	return home
}

// Todo contains basic info
type Todo struct {
	Title, Content string
}

// PageVariables are variables sent to the html template
type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

var todo []Todo

// GetTodos returns todo route and prints on browser
func GetTodos() func(w http.ResponseWriter, r *http.Request) {
	getTodo := func(w http.ResponseWriter, r *http.Request) {
		pageVariables := PageVariables{
			PageTitle: "Get Todos",
			PageTodos: todo,
		}

		t, err := template.ParseFiles("todo.html")

		if err != nil {
			fmt.Fprint(w, err.Error(), http.StatusBadRequest)
			log.Println("Template parsing error: ", err)
		} else {
			// t.Execute(w, nil)
			t.Execute(w, pageVariables)
		}
	}

	return getTodo
}

// AddTodo add todo to the existing todo list
func AddTodo() func(w http.ResponseWriter, r *http.Request) {
	addTodo := func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("Request parsing error: ", err)
		}

		todoToAdd := Todo{
			Title:   r.FormValue("title"),
			Content: r.FormValue("content"),
		}

		todo = append(todo, todoToAdd)
		fmt.Println("todo: ", todo)

		http.Redirect(w, r, "/todo", http.StatusSeeOther)
	}

	return addTodo
}
