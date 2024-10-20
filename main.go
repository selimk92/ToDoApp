package main

import "fmt"

func main() {

	var shortGolang = "Watch Go crash course"

	var fullGolang = "Watch Nana's Golang Full course"

	var rewardDessert = "Reward myself with a donut"

	var taskItems = []string{shortGolang, fullGolang, rewardDessert}

	fmt.Println("Welcome to Our TodoList App!!")
	fmt.Println()

	printTask(taskItems)
	fmt.Println()
	taskItems = addTask(taskItems, "Go for a run")
	fmt.Println()
	printTask(taskItems)

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
