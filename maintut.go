package main

import (
	"fmt"
	// "github.com/Harkishen-Singh/GO-db/gobase"
	"./gobase"
)

func main() {
	fmt.Println("GO-db is running")
	var app = "testbase"
	var details uint16
	details = 234
	gobase.SaveUint16(&app, &details)
	// fmt.Println(gobase.Retrive(&app))
	// fmt.Println(gobase.CollectionsAvailable("/"))
	// fmt.Println(gobase.Retrive("TestBase"))
	// var app = "someArr/Test"
	// gobase.SaveArr(&app, []string{"Harkishen", "Singh"})
	// var addr = "someArrays/one"
	// var vals = []float64{2.3, 5.64, 96.5}
	// gobase.SaveArrFloat64(&addr, &vals)
	// fmt.Println(gobase.RetriveArr(&addr))
	

}


