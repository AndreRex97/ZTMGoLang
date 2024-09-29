package main

import "fmt"

func main() {
	var myName = "Andre"
	fmt.Println("My name is", myName)

	var name = "Tammy"
	fmt.Println("Name =", name)

	username := "admin"
	fmt.Println("Username =", username)

	var sum int
	fmt.Println("Sum is", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is", part1, "other is", other)

	part2, other := 2, 0
	fmt.Println("part 2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("Sum =", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName=", lessonName)
	fmt.Println("lessonType=", lessonType)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1, word2)
}
