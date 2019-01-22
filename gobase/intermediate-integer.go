package gobase

import(
	// "fmt"
	"strconv"
)

// func preCondition(address *string){

// 	exists := collectionStatus(*address)
// 	if exists == false {
// 		fmt.Println("No Collection existing at the specified datapath. Creating one ...")
// 		createCollection(*address)
// 	}
// }

//SaveArrUint8 ...
func SaveArrUint8(path *string, data *[]uint8) bool{

	preCondition(path)

	dataInString := func(data *[]uint8) *[]string {

		var temp *[]string
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		return temp
	}(data)

	return SaveArr(path, *dataInString)
}

//SaveArrUint16 ...
func SaveArrUint16(path *string, data *[]uint16) bool{

	preCondition(path)

	dataInString := func(data *[]uint16) *[]string {

		var temp *[]string
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		return temp
	}(data)

	return SaveArr(path, *dataInString)
}

//SaveArrUint32 ...
func SaveArrUint32(path *string, data *[]uint32) bool{

	preCondition(path)

	dataInString := func(data *[]uint32) *[]string {

		var temp *[]string
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		return temp
	}(data)

	return SaveArr(path, *dataInString)
}

//SaveArrUint64 ...
func SaveArrUint64(path *string, data *[]uint64) bool{

	preCondition(path)

	dataInString := func(data *[]uint64) *[]string {

		var temp *[]string
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		return temp
	}(data)

	return SaveArr(path, *dataInString)
}

//SaveArrInt8 ...
func SaveArrInt8(path *string, data *[]int8) bool{

	preCondition(path)

	dataInString := func(data *[]int8) *[]string {

		var temp *[]string
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		return temp
	}(data)

	return SaveArr(path, *dataInString)
}

//SaveArrInt16 ...
func SaveArrInt16(path *string, data *[]int16) bool{

	preCondition(path)

	dataInString := func(data *[]int16) *[]string {

		var temp *[]string
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		return temp
	}(data)

	return SaveArr(path, *dataInString)
}

//SaveArrInt32 ...
func SaveArrInt32(path *string, data *[]int32) bool{

	preCondition(path)

	dataInString := func(data *[]int32) *[]string {

		var temp *[]string
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		return temp
	}(data)

	return SaveArr(path, *dataInString)
}

//SaveArrInt64 ...
func SaveArrInt64(path *string, data *[]int64) bool{

	preCondition(path)

	dataInString := func(data *[]int64) *[]string {

		var temp *[]string
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		return temp
	}(data)

	return SaveArr(path, *dataInString)
}

//SaveArrInt ...
func SaveArrInt(path *string, data *[]int) bool{

	preCondition(path)

	dataInString := func(data *[]int) *[]string {

		var temp *[]string
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(ele))
		}
		return temp
	}(data)

	return SaveArr(path, *dataInString)
}