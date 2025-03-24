package main

import (
	"BeamDB/internal/storage"
	"fmt"
	"os"
)

func main() {
	//  ==== Checkpoint 1 completed ====
	// Test saving some data
	data := []byte("Hello, Deep!")
	err := storage.SaveData("test_save.db", data)
	if err != nil {
		fmt.Println("Error saving data:", err)
		os.Exit(1)
	}

	fmt.Printf("Data saved successfully\n")

	// read back data to verify
	readData, readErr := os.ReadFile("test_save.db")
	if readErr != nil {
		fmt.Println("Error reading data:", err)
		os.Exit(1)
	}

	fmt.Printf("Data read: %s\n", string(readData))
}
