package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello-go", helloUser)

	http.ListenAndServe(":8080", nil)

}

func helloUser(w http.ResponseWriter, r *http.Request) {

	greeting := "Hello User! Welcome to TodoList App"
	fmt.Fprintln(w, greeting)

}
func printTask(taskItems2 []string) {

	fmt.Println("List of my Todos")

	for index, task := range taskItems2 {
		index = index + 1
		fmt.Println(index, ":", task)
	}

}

func addTask(taskItems1 []string, newTask string) []string {

	var updatedTaskItems = append(taskItems1, newTask)
	return updatedTaskItems

}
