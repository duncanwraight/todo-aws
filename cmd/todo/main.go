package main

import (
	"github.com/duncanwraight/todo-aws/internal/todo"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		todo.Add(os.Args[1])
	} else {
		todo.List()
	}
}
