package slices

import (
	"fmt"
	"strings"
)

func ExampleEvery() {
	fmt.Println(Every([]int{1, 30, 39, 29, 10, 13}, func(e int) bool {
		return e < 40
	}))
	// Output: true
}

func ExampleFilter() {
	fmt.Println(Filter([]string{"spray", "elite", "exuberant", "destruction", "present"}, func(e string) bool {
		return len(e) > 6
	}))
	// Output: [exuberant destruction present]
}

func ExampleFind() {
	fmt.Println(Find([]int{5, 12, 8, 130, 44}, func(e int) bool {
		return e > 10
	}))
	// Output: 12 true
}

func ExampleFindIndex() {
	fmt.Println(FindIndex([]int{5, 12, 8, 130, 44}, func(e int) bool {
		return e > 13
	}))
	// Output: 3
}

func ExampleForEach() {
	ForEach([]string{"a", "b", "c"}, func(e string) {
		fmt.Printf("%s ", strings.ToUpper(e))
	})
	// Output: A B C
}

func ExampleMap() {
	fmt.Println(Map([]int{1, 4, 9, 16}, func(e int) int {
		return e * 2
	}))
	// Output: [2 8 18 32]
}

func ExampleReduce() {
	fmt.Println(Reduce([]int{1, 2, 3, 4}, func(a int, e int) int {
		return a + e
	}, 0))
	// Output: 10
}

func ExampleSome() {
	fmt.Println(Some([]int{1, 2, 3, 4, 5}, func(e int) bool {
		return e%2 == 0
	}))
	// Output: true
}
