package gobase

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"os/exec"
	"log"
)


//RetriveArr ...
func RetriveArr(address string) (string, bool) {

	var documentAvailable = collectionStatus(address)
	var data string
	address = "warehouse/" + address + ".data"
	if documentAvailable {
		openfile, err := ioutil.ReadFile(address)
		if err != nil {
			return "ERROR", false
		}
		data = string(openfile)
		return data, true
	}
	return "DOCUMENT_UNAVAILABLE", false

}

//CollectionsAvailableArr ...
func CollectionsAvailableArr(address string) ([]string, bool) {

	var existingCollections []string
	if address != "/" {
		path := "warehouse/" + address
		response, err := exec.Command("ls", path).Output()
		if err != nil {
			fmt.Println("Error while looking for Collections, at Address: "+address)
			log.Fatal(err)
		}
		existingCollections = strings.Split(string(response), "\n")
	} else {
		response, err := exec.Command("ls", "warehouse/").Output()
		if err != nil {
			fmt.Println("Error while looking for Collections, at Address: "+address)
			log.Fatal(err)
		}
		existingCollections = strings.Split(string(response), "\n")
	}
	return existingCollections, true

}

//SaveArr ...
func SaveArr(path string, data string) bool {

	exists := collectionStatus(path)
	if exists == false {
		fmt.Println("No Collection existing at the specified datapath. Creating one ...")
		createCollection(path)
	}
	var address = "warehouse/" + path + ".data"
	file, err := os.OpenFile(address, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
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

//DeleteArr ...
func DeleteArr(path string) bool {

	path = "warehouse/" + path
	_, err := exec.Command("rm", "-R", path).Output()
	if err != nil {
		_, err2 := exec.Command("rm", "-R", path + ".data").Output()
		if err2 != nil {
			panic("Deletion not possible ''Type 2''. Collection doesnot exists in Address: " + path)
		}
	}
	return true

}
