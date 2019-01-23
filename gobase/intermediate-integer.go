package gobase

import "strconv"

//SaveArrUint8 ...
func SaveArrUint8(path *string, data *[]uint8) bool{

	preCondition(path)

	dataInString := func(data *[]uint8) *[]string {

		var temp *[]string
		final := []string{""}
		temp = &final
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		final = *temp
		final = final[1:]
		return &final
	}(data)

	return saveArrCustom(path, *dataInString,1)
}

//SaveArrUint16 ...
func SaveArrUint16(path *string, data *[]uint16) bool{

	preCondition(path)

	dataInString := func(data *[]uint16) *[]string {

		var temp *[]string
		final := []string{""}
		temp = &final
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		final = *temp
		final = final[1:]
		return &final
	}(data)

	return saveArrCustom(path, *dataInString,3)
}

//SaveArrUint32 ...
func SaveArrUint32(path *string, data *[]uint32) bool{

	preCondition(path)

	dataInString := func(data *[]uint32) *[]string {

		var temp *[]string
		final := []string{""}
		temp = &final
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		final = *temp
		final = final[1:]
		return &final
	}(data)

	return saveArrCustom(path, *dataInString,5)
}

//SaveArrUint64 ...
func SaveArrUint64(path *string, data *[]uint64) bool{

	preCondition(path)

	dataInString := func(data *[]uint64) *[]string {

		var temp *[]string
		final := []string{""}
		temp = &final
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		final = *temp
		final = final[1:]
		return &final
	}(data)

	return saveArrCustom(path, *dataInString,7)
}

//SaveArrInt8 ...
func SaveArrInt8(path *string, data *[]int8) bool{

	preCondition(path)

	dataInString := func(data *[]int8) *[]string {

		var temp *[]string
		final := []string{""}
		temp = &final
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		final = *temp
		final = final[1:]
		return &final
	}(data)

	return saveArrCustom(path, *dataInString,2)
}

//SaveArrInt16 ...
func SaveArrInt16(path *string, data *[]int16) bool{

	preCondition(path)

	dataInString := func(data *[]int16) *[]string {

		var temp *[]string
		final := []string{""}
		temp = &final
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		final = *temp
		final = final[1:]
		return &final
	}(data)

	return saveArrCustom(path, *dataInString,4)
}

//SaveArrInt32 ...
func SaveArrInt32(path *string, data *[]int32) bool{

	preCondition(path)

	dataInString := func(data *[]int32) *[]string {

		var temp *[]string
		final := []string{""}
		temp = &final
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		final = *temp
		final = final[1:]
		return &final
	}(data)

	return saveArrCustom(path, *dataInString,6)
}

//SaveArrInt64 ...
func SaveArrInt64(path *string, data *[]int64) bool{

	preCondition(path)

	dataInString := func(data *[]int64) *[]string {

		var temp *[]string
		final := []string{""}
		temp = &final
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(int(ele)))
		}
		final = *temp
		final = final[1:]
		return &final
	}(data)

	return saveArrCustom(path, *dataInString,8)
}

//SaveArrInt ...
func SaveArrInt(path *string, data *[]int) bool{

	preCondition(path)

	dataInString := func(data *[]int) *[]string {

		var temp *[]string
		final := []string{""}
		temp = &final
		for _,ele := range *data{
			*temp = append(*temp, strconv.Itoa(ele))
		}
		final = *temp
		final = final[1:]
		return &final
	}(data)

	return saveArrCustom(path, *dataInString,9)
}