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

func collectionStatus(collectionPath string) bool {
	// checking current files in warehouse

	collectionPath = "warehouse/" + collectionPath
	result, err := exec.Command("ls", collectionPath).Output()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	var existingFiles = strings.Split(string(result), "\n")
	fmt.Println("Existing files in warehouse/")
	fmt.Println(existingFiles)
	var checkStatus bool

	for _, element := range existingFiles {
		if element == collectionPath + ".data" {
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
	ioutil.WriteFile(address, []byte(""), 0777)
	return true

}

func readCollection(address string) (string, bool) {

	address = "warehouse/" + address
	file, err := ioutil.ReadFile(address)
	if err != nil {
		fmt.Println("Error while reading the collection at Address: " + address)
		return "", false
	}
	return string(file), true

}

//GetDocuments ...
func GetDocuments(address string) (string, bool) {

	var documentAvailable = collectionStatus(address)
	var data string
	address = "warehouse/" + address
	if documentAvailable {
		openfile, err := ioutil.ReadFile(address)
		if err != nil {
			data = string(openfile)
			return data, true
		}
		return "ERROR", false
	}
	return "DOCUMENT_UNAVAILABLE", false

}

//Save ...
func Save(path string, data string) bool {

	exists := collectionStatus(path)
	if exists == false {
		fmt.Println("No Collection existing at the specified datapath. Creating one ...")
		createCollection(path)
	}
	file, err := os.OpenFile(path, os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error occured while writing")
		fmt.Println(data)
		fmt.Println("Address: warehouse/"+path)
		return false
	}
	return true

}