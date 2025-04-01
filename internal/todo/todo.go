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

type List []ListItem

func (listItems *List) Add(items []string) {
	listItem := ListItem{
		Text:        strings.Join(items, " "),
		Priority:    Low, // TODO
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*listItems = append(*listItems, listItem)
	fmt.Printf("New item added: %s\n", listItem.Text)
}

func (listItems *List) List() {
	for _, listItem := range *listItems {
    completedAt := ""
    if listItem.CompletedAt.IsZero() {
      completedAt = listItem.CompletedAt.Format(time.Kitchen)
    }
		fmt.Printf(
			"%s - [%s] %s - Completed? [%s] (%s)\n",
			listItem.CreatedAt.Format(time.Kitchen),
			listItem.Priority,
			listItem.Text,
			listItem.Completed,
			completedAt,
		)
	}
}
