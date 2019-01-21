package main

import (
	"fmt"
	"github.com/Harkishen-Singh/Go-db/gobase"
)

func main() {
	fmt.Println("GO-db is running")
	gobase.Save("TestBase", "Random Test")
}