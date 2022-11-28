package handler

import (
	"context"
	"fmt"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func Handler(ctx context.Context, person Person) error {
	fmt.Println("Hello", person.FirstName, person.LastName)
	return nil
}
