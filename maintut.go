package main

import (
	"fmt"
	"github.com/Harkishen-Singh/GO-file-database/gobase"
	// "gobase"
)

func main() {
	fmt.Println("GO-db is running")
	gobase.Save("TestBase", "Rand")
	fmt.Println(gobase.CollectionsAvailable("/"))
	fmt.Println(gobase.Retrive("TestBase"))
	var app = "someArr/Test"
	gobase.SaveArr(&app, []string{"Hahen", "Singh"})
	var addr = "someArrays/one"
	var vals = []float64{2.3, 5.64, 96.5}
	gobase.SaveArrFloat64(&addr, &vals)
	fmt.Println(gobase.RetriveArr(&addr))
}

