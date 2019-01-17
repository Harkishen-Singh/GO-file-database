package main

import (
	"fmt"
	"os/exec"
	"log"
	"strings"
)

//StorageStatus ...
func StorageStatus() bool {
	// checking current files in warehouse
	result, err := exec.Command("ls", "warehouse").Output()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	var existingFiles = strings.Split(string(result), "\n")
	fmt.Println("Existing files in warehouse/")
	fmt.Println(existingFiles)
	var checkStore bool

	for _, element := range existingFiles {
		if element == "MainStore.json" {
			checkStore = true
			break
		} else {
			checkStore = false
		}
	}

	if checkStore {
		return true
	}

}

// func SaveString(addr, message string) bool {

// }

func main() {
	Storage()
}