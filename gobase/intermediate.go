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
func RetriveArr(address *string) ([]string, bool) {

	var documentAvailable = collectionStatus(*address)
	var data []string
	*address = "warehouse/" + *address + ".data"
	if documentAvailable {
		openfile, err := ioutil.ReadFile(*address)
		if err != nil {
			return []string{}, false
		}
		temp := string(openfile)
		data = strings.Split(temp[1: len(temp) -1], ",")
		return data, true
	}
	return []string{}, false

}

//CollectionsAvailableArr ...
func CollectionsAvailableArr(address *string) ([]string, bool) {

	var existingCollections []string
	if *address != "/" {
		path := "warehouse/" + *address
		response, err := exec.Command("ls", path).Output()
		if err != nil {
			fmt.Println("Error while looking for Collections, at Address: "+*address)
			log.Fatal(err)
		}
		existingCollections = strings.Split(string(response), "\n")
	} else {
		response, err := exec.Command("ls", "warehouse/").Output()
		if err != nil {
			fmt.Println("Error while looking for Collections, at Address: "+*address)
			log.Fatal(err)
		}
		existingCollections = strings.Split(string(response), "\n")
	}
	return existingCollections, true

}

//SaveArr ...
func SaveArr(path *string, dataArr []string) bool {

	exists := collectionStatus(*path)
	if exists == false {
		fmt.Println("No Collection existing at the specified datapath. Creating one ...")
		createCollection(*path)
	}
	var address = "warehouse/" + *path + ".data"
	file, err := os.OpenFile(address, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var dataString = "["
	dataString += strings.Join(dataArr, ",")
	dataString += "]"
	_, err = file.WriteString(dataString)
	if err != nil {
		fmt.Println("Error occured while writing the following data:")
		fmt.Println("\n" + dataString)
		fmt.Println("Address: warehouse/"+address)
		fmt.Println(err)
		return false
	}
	return true

}
