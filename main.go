package main

import "fmt"

func main() {

	var shortGolang = "Watch Go crash course"

	var fullGolang = "Watch Nana's Golang Full course"

	var rewardDessert = "Reward myself with a donut"

	taskItems := []string{shortGolang, fullGolang, rewardDessert}

	fmt.Println(taskItems)
	fmt.Println(taskItems[0])
}
