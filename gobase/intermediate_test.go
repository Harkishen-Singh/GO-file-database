package gobase

import (
	"testing"
	"reflect"
	"fmt"
)

type intermediatePlainType struct {
	location string
	data []string
	retri []string
	typeStatus string
	status bool
}

type intermediateDecimalType32 struct {
	location string
	data []float32
	retri []string
	typeStatus string
	status bool
}

type intermediateDecimalType64 struct {
	location string
	data []float64
	retri []string
	typeStatus string
	status bool
}

func TestSaveArr(t *testing.T) {

	var testObjects = []intermediatePlainType {
		{"testIntermediateArr",[]string{"first1", "second1", "third1", "fourth1", "five1"},[]string{}, "", false},
		{"testIntermediateArr/somelocation/random", []string{"first2", "second2", "third2", "fourth2", "five2"}, []string{}, "", false},
		{"testIntermediateArray/somelocation/random", []string{"first3", "second3", "third3", "fourth3", "five3"}, []string{}, "", false},
		{"testIntermediateArr2", []string{"first4", "second4", "third4", "fourth4", "five4"}, []string{}, "", false},
	}

	for _, ele := range testObjects {

		// save operation
		SaveArr(ele.location, &ele.data)

		// retrive operation
		ele.retri, ele.typeStatus, ele.status = RetriveArr(ele.location)

		fmt.Println("Original: ", ele)
		fmt.Println("Received: ", ele)
		if !(reflect.DeepEqual(ele.retri, ele.data) && ele.status) {
			t.Errorf("Tests failed for location: %s | value: %s | received status: %t | received value: %s", ele.location, ele.data, ele.status, ele.retri)
		}

	}
	//Delete Operation
	Delete(testObjects[0].location)
	Delete(testObjects[1].location)
	Delete(testObjects[2].location)
	Delete(testObjects[3].location)

}