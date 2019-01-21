package main

import (
	"fmt"
	"github.com/Harkishen-Singh/Go-db/gobase"
)

func main() {
	fmt.Println("GO-db is running")
	var names = "harkishen"
	gobase.Save("names/firstOfAll/one", names)
	// b, _ := gobase.GetCollections("names")
	// fmt.Println(b)
}