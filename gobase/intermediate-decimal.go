package gobase

import (
	"fmt"
	"strconv"
)

func preCondition(address *string) {

	exists := collectionStatus(*address)
	if exists == false {
		fmt.Println("No Collection existing at the specified datapath. Creating one ...")
		createCollection(*address)
	}

}

//SaveArrFloat32 ...
func SaveArrFloat32(address *string, data *[]float32) bool {

	preCondition(address)

	increasePrecision := func (data *[]float32) *[]float64 {

		var temp *[]float64
		for _, ele := range *data {
			*temp = append(*temp, float64(ele))
		}
		return temp
	}
	dataInString := func (data *[]float64) *[]string {

		var temp *[]string
		final := []string{""}
		temp = &final

		for _, ele := range *data {
			*temp = append(*temp, strconv.FormatFloat(ele, 'E', -1, 32))
		}
		final = *temp
		final = final[1:]
		return &final
	} (increasePrecision(data))

	return saveArrCustom(address, *dataInString, 10)

}

//SaveArrFloat64 ...
func SaveArrFloat64(address *string, data *[]float64) bool {

	preCondition(address)

	dataInString := func (data *[]float64) *[]string {

		var temp *[]string
		final := []string{""}
		temp = &final
		for _, ele := range *data {
			*temp = append(*temp, strconv.FormatFloat(ele, 'E', -1, 32))
		}
		final = *temp
		final = final[1:]
		return &final
	} (data)

	return saveArrCustom(address, *dataInString, 11)

}