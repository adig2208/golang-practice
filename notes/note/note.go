package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json: "content"`
	CreatedAt time.Time `json: "created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("missing value, invalid input")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now()}, nil
}

func (n *Note) Display() {
	fmt.Println("Title:", n.Title)
	fmt.Println("Content:", n.Content)
}

func (n *Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
