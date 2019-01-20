package main

import (
	"fmt"
	"github.com/Harkishen-Singh/Go-db/funcs"
)

func main() {
	fmt.Println("GO-db is running")
	// funcs.Save("newTests", "this the the first form of package")
	funcs.Delete("newTests")
}