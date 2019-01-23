package gobase

import (
	"testing"
)

type SaveArrResult struct{
	path string
	data []string
	expected bool
	checkeddata string
} 

var SaveArrResults = []SaveArrResult{
	// {"a", {"The data is saved", "Hello"}, true, "[The data is saved,Hello]"},
}

func TestSaveArr(t *testing.T){
		for _,test := range SaveArrResults{
			result := SaveArr(&(test.path), &(test.data))
			if result != test.expected{
				t.Fatal("Expected Result Not Given")
			}

		}
}