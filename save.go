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

	res := strings.Index(collectionPath, "/")
	var subPath string
	if res != -1 {
		subPath = collectionPath[0: res]
	}
	expAddr := "warehouse/" + subPath
	result, err := exec.Command("ls", expAddr).Output()
	if err != nil {
		fmt.Println("Path not found!")
		return false
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

	address = "warehouse/" + address + ".data"
	fmt.Println("Address :: "+address)
	err := ioutil.WriteFile(address, []byte("default statement"), 0777)
	if err != nil {
		fmt.Println("Error in createCollection Address: "+address)
		fmt.Println(err)
		return false
	}
	return true

}

func readCollection(address string) (string, bool) {

	address = "warehouse/" + address + ".data"
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
	address = "warehouse/" + address + ".data"
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

//GetCollections ...
func GetCollections(address string) ([]string, bool) {

	path := "warehouse/" + address
	response, err := exec.Command("ls", path).Output()
	if err != nil {
		fmt.Println("Error while looking for Collections, at Address: "+address)
		log.Fatal(err)
	}
	var existingCollections = strings.Split(string(response), "\n")
	fmt.Println("Present collections at Address: ", address)
	fmt.Println(existingCollections)
	return existingCollections, true

}

//Save ...
func Save(path string, data string) bool {

	exists := collectionStatus(path)
	if exists == false {
		fmt.Println("No Collection existing at the specified datapath. Creating one ...")
		createCollection(path)
	}
	var address = "warehouse/" + path + ".data"
	file, err := os.OpenFile(address, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("second err here")
		panic(err)
	}
	data = "\n" + data
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error occured while writing the following data:")
		fmt.Println("\n" + data)
		fmt.Println("Address: warehouse/"+address)
		fmt.Println(err)
		return false
	}
	return true

}

func main() {
	Save("something" ,"Harkishen singh is the best")
}