package tests

import (
	"testing"
	"github.com/Harkishen-Singh/GO-file-database/gobase"
	// "../gobase"
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

	var plainTests = []plainContainer {
		{"testPlain", "some test data", "", false, "string",""},
		{"testPlain/test1", "some test data1", "", false,"string",""},
		{"testPlainSecond/test2/test3/test4", "@some special test data with sym", "", false,"string",""},
		{"testPlain2", "@#$%#@(*&^%$)", "", false,"string",""},
	}


	for _,ele := range plainTests {

		// Save Operation
		gobase.Save(&(ele.location), &ele.value)
 		// Retrive Operation
		ele.retr, ele.typedata, ele.status = gobase.Retrive(&(ele.location))

		if (ele.value == ele.retr) && ele.status && (ele.typedata == ele.checkedData) {
			t.Errorf("Tests failed for location: %s | value: %s | received status: %t | received value: %s", ele.location, ele.value, ele.status, ele.retr)
		}

	}

	// gobase.Delete(plainTests[0].location)
	// gobase.Delete(plainTests[2].location)
	// gobase.Delete(plainTests[3].location)

}

