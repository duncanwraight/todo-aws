package todo

import "fmt"

func List() {
	fmt.Println("You have a list")
}

func Add(item string) {
	fmt.Printf("New item added: %s\n", item)
}
