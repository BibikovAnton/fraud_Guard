package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Printf("uuid.Nil: %q\n", uuid.Nil.String())
	fmt.Printf("Empty string: %q\n", "")
}
