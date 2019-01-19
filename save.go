package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"os/exec"
	"log"
)

const separator = ".##."

func filter(text string) (string, bool) {

	if strings.Contains(text, separator) {
		return strings.Replace(text, separator, " ", -1), true
	}
	return "", false
}

func collectionStatus(collectionName string, expAddress string) bool {
	// checking current files in warehouse

	expAddress = "warehouse/" + expAddress
	result, err := exec.Command("ls", expAddress).Output()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	var existingFiles = strings.Split(string(result), "\n")
	fmt.Println("Existing files in warehouse/")
	fmt.Println(existingFiles)
	var checkStatus bool

	for _, element := range existingFiles {
		if element == collectionName + ".data" {
			checkStatus = true
			break
		} else {
			checkStatus = false
		}
	}

	return checkStatus
}

func createCollection(address string) bool {

	address = "warehouse/" + address
	file := ioutil.WriteFile(address, []byte(""), 0777)
	return true
}

func readCollection(address string) (string, bool) {

	address = "warehouse/" + address
	file, err := ioutil.ReadFile(address)
	if err != nil {
		panic("Error while reading the collection at Address: " + address)
		return "", false
	}
	return string(file), true
}

//Save ...
func Save(path string, data string) {

	exists := collectionStatus(path, path)
}