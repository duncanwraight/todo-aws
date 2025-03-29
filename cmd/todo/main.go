package main

import (
	"fmt"
	"github.com/duncanwraight/todo-aws/internal/todo"
	"os"
)

var cmd []string

func main() {
	if len(os.Args) > 1 {
		cmd = os.Args[1:]

		switch cmd[0] {
		case "add":
			todoListItems := todo.ListItems{}
			todoListItems.Add(os.Args[1:])
			todoListItems.List()
		default:
			fmt.Println("Usage: add <item>")
		}
	}

}
