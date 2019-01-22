package main

import (
	"fmt"
	// "github.com/Harkishen-Singh/GO-db/gobase"
	"./gobase"
)

func main() {
	fmt.Println("GO-db is running")
	gobase.Save("TestBase", "Random Test")
	fmt.Println(gobase.CollectionsAvailable("/"))
	fmt.Println(gobase.Retrive("TestBase"))
	var app = "someArr/Test"
	b := []float64{1, 2, 3, 4, 5}
	gobase.SaveArr(&app, []string{"Harkishen", "Singh"})
	gobase.SaveArrFloat64(&app, &b)
}