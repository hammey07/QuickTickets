package main

import (
	"fmt"
)

func main() {
	type Event struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Category string `json:"category"`
		Seat     string `json:"seat"`
	}

	type Seats struct {
		ID       int    `json:"id"`
		IsBooked string `json:"isBooked"`
	}

	fmt.Println("Init application backend")
}
