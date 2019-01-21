package main

import (
	"fmt"
	"github.com/Harkishen-Singh/GO-db/gobase"
)

func main() {
	fmt.Println("GO-db is running")
	gobase.Save("TestBase", "Random Test")
	fmt.Println(gobase.CollectionsAvailable("/"))
}