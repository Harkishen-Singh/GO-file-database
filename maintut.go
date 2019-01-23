package main

import (
	"fmt"
	"github.com/Harkishen-Singh/GO-file-database/gobase"
	// "gobase"
)

func main() {
	fmt.Println("GO-db is running")
	fmt.Println(gobase.CollectionsAvailable("/"))
	var addr = "someArrays/one"
	var vals = []float64{2.3, 5.64, 96.5}
	gobase.SaveArrFloat64(&addr, &vals)
	fmt.Println(gobase.RetriveArr(&addr))
	var path = "tests/muskan/details"
	var details = "Muskan Khedia"
	gobase.Save(&path, &details)
}

