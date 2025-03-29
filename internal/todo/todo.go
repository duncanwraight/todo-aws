package todo

import (
	"fmt"
	"strings"
	"time"
)

type Enum uint8

const (
	Low Enum = iota
	Medium
	High
	Urgent
)

type ListItem struct {
	Text        string
	Priority    Enum
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type ListItems []ListItem

func (listItems *ListItems) Add(items []string) {
	listItem := ListItem{
		Text:        strings.Join(items, " "),
		Priority:    Low,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*listItems = append(*listItems, listItem)
	fmt.Printf("New item added: %s\n", listItem.Text)
}

func (listItems *ListItems) List() {
	for _, listItem := range *listItems {
		fmt.Printf("%s\n", listItem.Text)
	}
}
