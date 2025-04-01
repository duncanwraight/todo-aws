package main

import (
	"fmt"
	"github.com/duncanwraight/todo-aws/internal/todo"
	"os"
)

var cmd []string

func usage() {
	fmt.Println("Usage: todo <list|add|del> [arguments]")
	os.Exit(1)
}

func main() {
	numArgs := len(os.Args)
	if numArgs > 1 {
		cmd = os.Args[1:]
		todoList := todo.List{}

		switch cmd[0] {
		case "list":
			todoList.Add(os.Args[2:])
			todoList.List()
		case "add":
			if numArgs == 2 {
				usage()
			}
			todoList.Add(os.Args[1:])
		case "del":
			fmt.Println("Implement deletion")
		default:
			usage()
		}
		todoList.List()
	} else {
		fmt.Println("Usage: todo <list|add|del>")
	}

}
