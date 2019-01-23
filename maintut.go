package main

import (
	"fmt"
	"github.com/Harkishen-Singh/GO-file-database/gobase"
	// "./gobase"
)

func main() {
	fmt.Println("GO-db is running")
	var x = "TestBase"
	var y = "Rand"
	gobase.Save(&x, &y)
	fmt.Println(gobase.CollectionsAvailable("/"))
	fmt.Println(gobase.Retrive(&x))
	var app = "someArr/Test"
	name := []string{"Harkishen", "Singh"}
	gobase.SaveArr(&app, &name)
	var addr = "someArrays/one"
	var vals = []float64{2.3, 5.64, 96.5}
	gobase.SaveArrFloat64(&addr, &vals)
	fmt.Println(gobase.RetriveArr(&addr))
	var details uint64
	details = 534532545342
	var addr2 = "someArrays/two"
	gobase.SaveUint64(&addr2, &details)

}

