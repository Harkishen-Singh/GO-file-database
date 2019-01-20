package main

import (
	"io/ioutil"

)

func main() {
	ioutil.WriteFile("warehouse/extra/seom.data", []byte("some default text"), 0777)

}