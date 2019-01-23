package tests

import (
	"github.com/Harkishen-Singh/GO-file-database/gobase"
	"testing"
	"reflect"
	// "fmt"
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

func TestSaveArr_tests(t *testing.T) {

	var testObjects = []intermediatePlainType {
		{"testIntermediateArr",[]string{"first1", "second1", "third1", "fourth1", "five1"}, nil, "", false},
		{"testIntermediateArr/somelocation/random", []string{"first2", "second2", "third2", "fourth2", "five2"}, nil, "", false},
		{"testIntermediateArray/somelocation/random", []string{"first3", "second3", "third3", "fourth3", "five3"}, nil, "", false},
		{"testIntermediateArr2", []string{"first4", "second4", "third4", "fourth4", "five4"}, nil, "", false},
	}

	for _, ele := range testObjects {

		// save operation
		gobase.SaveArr(&ele.location, &ele.data)

		// retrive operation
		ele.retri, ele.typeStatus, ele.status = gobase.RetriveArr(&ele.location)
		// fmt.Println("Original: ", ele)
		// fmt.Println("Received: ", ele)
		if !(reflect.DeepEqual(ele.retri, ele.data) || ele.status) {
			t.Errorf("Tests failed for location: %s | value: %s | received status: %t | received value: %s", ele.location, ele.data, ele.status, ele.retri)
		}

	}
	gobase.Delete(testObjects[0].location)
	gobase.Delete(testObjects[1].location)
	gobase.Delete(testObjects[2].location)
	gobase.Delete(testObjects[3].location)

}