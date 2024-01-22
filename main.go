package main

import "fmt"


func main() {
	var input_list [5]string
	result_map := map[string]int{}
	input_list[0] = "cat"
	input_list[1] = "bat"
	input_list[2] = "cat"
	input_list[3] = "cat"
	input_list[4] = "rat"

	var word string = ""
	for i := 0; i < 5; i++ {
		// fmt.Println("word: ", input_list[i])
		word = input_list[i]
		_, ok := result_map[word]
		if !ok{
			result_map[word] = 0
		}
		result_map[word] += 1
	}

	fmt.Println(result_map)

	// ======
	var input_list2 [10] string
	var counter int = 0
	input_list2[0] = "c"
	input_list2[1] = "h"
	input_list2[2] = "a"
	input_list2[3] = "t"
	input_list2[4] = "c"
	input_list2[5] = "h"
	input_list2[6] = "a"
	input_list2[7] = "w"
	input_list2[8] = "a"
	input_list2[9] = "n"

	filtermap := map[string]int{
		"a":0,
		"e":0,
		"i":0,
		"o":0,
		"u":0,
	}

	for i := 0; i < 10; i++ {
		// fmt.Println(input_list2[i])
		_, ok := filtermap[input_list2[i]]
		if ok {
			counter += 1
		}
	}

	fmt.Println("counter [a/e/i/o/u] = ",counter)


}