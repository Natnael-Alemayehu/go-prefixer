package main

import "fmt"

func preFixer(prefix string) func(string) string {
	return func(pre string) string {
		return prefix + " " + pre
	}
}

func main() {
	helloPrefix := preFixer("hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
