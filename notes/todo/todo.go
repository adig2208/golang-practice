package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type todo struct {
	Text string `json: text`
}

func New(content string) (*todo, error) {
	if content == "" {
		return nil, errors.New("missing value, invalid input")
	}
	return &todo{
		Text: content}, nil
}

func (t *todo) Display() {
	fmt.Println("Text:", t.Text)
}

func (t *todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
