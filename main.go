package main

import "fmt"

var shortGolang = "Watch Go crash course"

var fullGolang = "Watch Nana's Golang Full course"

var rewardDessert = "Reward myself with a donut"

var taskItems = []string{shortGolang, fullGolang, rewardDessert}

func main() {

	fmt.Println("Welcome to Our TodoList App!!")
	fmt.Println()

	printTask()

}

func printTask() {

	fmt.Println("List of my Todos")

	for index, task := range taskItems {
		index = index + 1
		fmt.Println(index, ":", task)
	}

}
