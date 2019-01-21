package main

import (
	"fmt"
	"github.com/Harkishen-Singh/Go-db/goBase"
)

func main() {
	fmt.Println("GO-db is running")
	var names = "harkishen"
	gobase.Save("names", names)
	// b, _ := gobase.GetCollections("names")
	// fmt.Println(b)
}