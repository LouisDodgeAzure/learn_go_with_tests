package main

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	g := Hello("louis")
	fmt.Println(g)
}
