package gobase

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

//RetriveArr ...
func RetriveArr(address string) ([]string, string, bool) {

	var documentAvailable = collectionStatus(address)
	var data []string
	address = EnvironmentPath + "warehouse/" + address + ".data"
	if documentAvailable {
		openfile, err := ioutil.ReadFile(address)
		if err != nil {
			return []string{}, "", false
		}
		temp := string(openfile)
		var typeData string
		typeData = temp[:6]
		temp = temp[6:]
		data = strings.Split(temp[1:len(temp)-1], ",")
		return data, typeData, true
	}
	return []string{}, "", false

}

//CollectionsAvailableArr ...
func CollectionsAvailableArr(address string) ([]string, bool) {

	var existingCollections []string
	if address != "/" {
		path := EnvironmentPath + "warehouse/" + address
		response, err := exec.Command("ls", path).Output()
		if err != nil {
			fmt.Println("Error while looking for Collections, at Address: " + address)
			log.Fatal(err)
		}
		existingCollections = strings.Split(string(response), "\n")
	} else {
		response, err := exec.Command("ls", EnvironmentPath+"warehouse/").Output()
		if err != nil {
			fmt.Println("Error while looking for Collections, at Address: " + address)
			log.Fatal(err)
		}
		existingCollections = strings.Split(string(response), "\n")
	}
	return existingCollections, true

}

//SaveArr ...
func SaveArr(path string, dataArr *[]string) bool {

	return saveArrCustom(path, *dataArr, 12)

}

func saveArrCustom(path string, dataArr []string, pass uint16) bool {

	exists := collectionStatus(path)
	if exists == false {
		createCollection(path)
	}
	var address = EnvironmentPath + "warehouse/" + path + ".data"
	file, err := os.OpenFile(address, os.O_WRONLY, 0600)
	file.Seek(0, 0)
	file.Truncate(0)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var dataString = "["
	dataString += strings.Join(dataArr, ",")
	dataString += "]"

	dataString = CheckType(pass) + dataString

	_, err = file.WriteString(dataString)
	if err != nil {
		fmt.Println("Error occurred while writing the following data:")
		fmt.Println("\n" + dataString)
		fmt.Println("Address: warehouse/" + address)
		fmt.Println(err)
		return false
	}
	return true

}
