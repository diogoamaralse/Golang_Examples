package pratices

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func getUser(id int) (string, error) {
	if id != 1 {
		return "", fmt.Errorf("user %d: %w", id, ErrNotFound)
	}
	return "Alice", nil
}

func ErrorHandling() {
	_, err := getUser(2)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("User not found!")
	}
}
