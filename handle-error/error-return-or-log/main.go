package main

import (
	"fmt"
	"log"
)

func main() {
	if err := topFunction(); err != nil {
		log.Printf("Error: %v", err)
	}
}

func topFunction() error {
	err := level1Function()
	if err != nil {
		return fmt.Errorf("topFunction: %w", err)
	}
	return nil
}

func level1Function() error {
	err := level2Function()
	if err != nil {
		return fmt.Errorf("level1Function: %w", err)
	}
	return nil
}

func level2Function() error {
	err := level3Function()
	if err != nil {
		return fmt.Errorf("level2Function: %w", err)
	}
	return nil
}

func level3Function() error {
	err := level4Function()
	if err != nil {
		return fmt.Errorf("level3Function: %w", err)
	}
	return nil
}

func level4Function() error {
	err := fmt.Errorf("something went wrong")
	return fmt.Errorf("level4Function: %w", err)
}
