package gobase

import (
	"testing"
	"fmt"
)

type plainContainer struct {
	location string
	value string
	retr string
	status bool
	typedata string
	checkedData string
}

func TestPlain(t *testing.T) {

	var plainTests = []plainContainer{
		{"testPlain", "some test data", "", false, "string",""},
		{"testPlain/test1", "some test data1", "", false,"string",""},
		{"testPlainSecond/test2/test3/test4", "@some special test data with sym", "", false,"string",""},
		{"testPlain2", "@#$%#@(*&^%$)", "", false,"string",""},
	}


	for _,ele := range plainTests{

		//Save Operations
		Save(ele.location, &ele.value)
		//Retrive Operations
		ele.retr, ele.typedata, ele.status = Retrive(ele.location)
		fmt.Println("Original: ", ele)
		fmt.Println("Received: ", ele)

		if (ele.value == ele.retr) && ele.status && (ele.typedata == ele.checkedData){
			t.Errorf("Tests failed for location: %s | value: %s | received status: %t | received value: %s", ele.location, ele.value, ele.status, ele.retr)
		}

		//Delete Operation
		Delete(ele.location)
	}

}
