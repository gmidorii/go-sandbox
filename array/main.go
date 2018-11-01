package main

import "fmt"

var arr = [...]int{1, 2, 3}

func main() {
	// append not use array
	//arr = append(arr, 4)

	// compile error (array size check in compile)
	//arr[3] = 4

	fmt.Println(arr)
}
