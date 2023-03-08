package main

import "fmt"

func main() {
	//currentpresident := "4"
	array := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	array = append(array[:5], array[6:]...)
	//	match := Match{playernames: array}
	//	match.LaunchGame()
	// step 1 filter alive people
	//step 1.5 get index of current president
	// step 2 // give presidency to next guy by index
	// if last guy
}

func main3() {
	s1 := []int{1, 2, 3}
	s2 := []int{99, 100}
	s1 = append(s1, s2...)

	fmt.Println(s1) // [1 2 3 99 100]
}

type T struct {
	a string
}

func main2() {

	players := make(map[string]*T)

	array := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	//	match := Match{playernames: array}
	//	match.LaunchGame()
	for i := 0; i < len(array); i++ {
		players[array[i]] = &T{a: array[i]}
	}

	for i := 0; i < len(players); i++ {
		players[array[i]].a = "5"
	}

	delete(players, "2")

	for s := range players {
		fmt.Println(s)
	}

	for i := 0; i < len(players); i++ {
		fmt.Print(players[array[i]])
	}
}

func main1() {

	//match := Match{}
	//match.LaunchGame()
	//fmt.Print(match.players)
	//fmt.Println(" hitler is ", match.hitler)
	//fmt.Println(match.players["1"].hasVoted)
}
