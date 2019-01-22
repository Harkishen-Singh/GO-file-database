package main

import (
	"testing"
	"../gobase"
)

type SaveArrResult struct{
	path string
	data []string
	expected bool
	checked_data string
} 

var SaveArrResults = []SaveArrResult{
	{"a", {"The data is saved", "Hello"}, true, "[The data is saved,Hello]"},
}

func TestSaveArr(t *testing.T){
		for _,test := range SaveArrResults{
			result := gobase.SaveArr(&(test.path), test.data)
			if result != test.expected{
				t.Fatal("Expected Result Not Given")
			}

		}
}