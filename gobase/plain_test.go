package gobase

import (
	"testing"
)

func TestPlain(t *testing.T) {

	var location1, location2, location3, location4 string
	var value1, value2, value3, value4 string
	var retr1, retr2, retr3, retr4 string
	var status1, status2, status3, status4 bool

	location1, location2, location3, location4 = "testPlain", "testPlain/test1", "testPlainSecond/test2/test3/test4", "testPlain2"
	value1, value2, value3, value4 = "some test data", "some test data1", "@some special test data with sym", "@#$%#@(*&^%$)"

	// save operations

	Save(location1, value1)
	Save(location2, value2)
	Save(location3, value3)
	Save(location4, value4)

	// retrive operations

	retr1, status1 = Retrive(location1)
	retr2, status2 = Retrive(location2)
	retr3, status3 = Retrive(location3)
	retr4, status4 = Retrive(location4)

	// checks

	if !(value1 == retr1) || !status1 {
		t.Errorf("Tests failed for location: %s | value: %s | received status: %t | received value: %s", location1, value1, status1, retr1)
	}

	if !(value2 == retr2) || !status2 {
		t.Errorf("Tests failed for location: %s | value: %s | received status: %t | received value: %s", location2, value2, status2, retr2)
	}

	if !(value3 == retr3) || !status3 {
		t.Errorf("Tests failed for location: %s | value: %s | received status: %t | received value: %s", location3, value3, status3, retr3)
	}

	if !(value4 == retr4) || !status4 {
		t.Errorf("Tests failed for location: %s | value: %s | received status: %t | received value: %s", location4, value4, status4, retr4)
	}
}
