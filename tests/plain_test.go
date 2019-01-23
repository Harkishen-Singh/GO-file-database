package tests

import (
	"testing"
	"github.com/Harkishen-Singh/GO-file-database/gobase"
)

type plainContainer struct {
	location string
	value string
	retr string
	status bool
}

func TestPlain(t *testing.T) {

	var test1, test2, test3, test4 plainContainer

	test1 = plainContainer{"testPlain", "some test data", "", false}
	test2 = plainContainer{"testPlain/test1", "some test data1", "", false}
	test3 = plainContainer{"testPlainSecond/test2/test3/test4", "@some special test data with sym", "", false}
	test4 = plainContainer{"testPlain2", "@#$%#@(*&^%$)", "", false}

	// save operations

	gobase.Save(test1.location, test1.value)
	gobase.Save(test2.location, test2.value)
	gobase.Save(test3.location, test3.value)
	gobase.Save(test4.location, test4.value)

	// retrive operations

	test1.retr, test1.status = gobase.Retrive(test1.location)
	test2.retr, test2.status = gobase.Retrive(test2.location)
	test3.retr, test3.status = gobase.Retrive(test3.location)
	test4.retr, test4.status = gobase.Retrive(test4.location)

	// checks

	if !(test1.value == test1.retr) || !test1.status {
		t.Errorf("Tests failed for location: %s | value: %s | received status: %t | received value: %s", test1.location, test1.value, test1.status, test1.retr)
	}

	if !(test2.value == test2.retr) || !test2.status {
		t.Errorf("Tests failed for location: %s | value: %s | received status: %t | received value: %s", test2.location, test2.value, test2.status, test2.retr)
	}

	if !(test3.value == test3.retr) || !test3.status {
		t.Errorf("Tests failed for location: %s | value: %s | received status: %t | received value: %s", test3.location, test3.value, test1.status, test3.retr)
	}

	if !(test4.value == test4.retr) || !test4.status {
		t.Errorf("Tests failed for location: %s | value: %s | received status: %t | received value: %s", test4.location, test4.value, test4.status, test4.retr)
	}
}
