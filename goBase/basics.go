package funcs

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

func checkExistingDir(name string) bool {

	name = "warehouse/" + name
	_, err := exec.Command("ls", name).Output()
	if err != nil {
		return false
	}
	return true
}

func makeDir(name string) bool {

	if !checkExistingDir(name) {
		name = "warehouse/" + name
		fmt.Println("hisssss:  " + name)
		_, err := exec.Command("mkdir", name).Output()
		if err != nil {
			fmt.Println("Err makeDir")
			panic(err)
		}
	}
	return true

}

func collectionStatus(collectionPath string) bool {

	res := strings.LastIndex(collectionPath, "/")
	var subPath string
	if res != -1 {
		subPath = collectionPath[0: res]

		if !makeDir(subPath) {
			os.Exit(50005)
		}
	}
	expAddr := "warehouse/" + subPath
	result, err := exec.Command("ls", expAddr).Output()
	if err != nil {
		fmt.Println("Path not found!")
		return false
	}
	var existingFiles = strings.Split(string(result), "\n")
	var checkStatus bool

	for _, element := range existingFiles {
		if element == collectionPath[res+1:] + ".data" {
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
	err := ioutil.WriteFile(address, []byte(""), 0777)
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
			return "ERROR", false
		}
		data = string(openfile)
		return data, true
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

//Delete ...
func Delete(path string) bool {

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

//Warehouse ...
func Warehouse() {

	db := "warehouse"
	resp, err := exec.Command("ls").Output()
	if err != nil {
		panic(err)
	}
	var checkWarehouse bool
	var respStringArr = strings.Split(string(resp), "\n")
	for _, ele := range respStringArr {
		if ele == "warehouse" {
			checkWarehouse = true
			break
		}
	}
	if !checkWarehouse {
		exec.Command("mkdir", db).Output()
	}

}