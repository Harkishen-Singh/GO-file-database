package gobase

import "strconv"

//SaveUint8 ...
func SaveUint8(path string, data *uint8) bool {

	preCondition(path)

	var dataInString string
	dataInString = strconv.Itoa(int(*data))

	return saveCustom(path, dataInString, 1)
}

//SaveInt8 ...
func SaveInt8(path string, data *int8) bool {

	preCondition(path)

	var dataInString string
	dataInString = strconv.Itoa(int(*data))

	return saveCustom(path, dataInString, 2)
}

//SaveUint16 ...
func SaveUint16(path string, data *uint16) bool {

	preCondition(path)

	var dataInString string
	dataInString = strconv.Itoa(int(*data))

	return saveCustom(path, dataInString, 3)
}

//SaveInt16 ...
func SaveInt16(path string, data *int16) bool {

	preCondition(path)

	var dataInString string
	dataInString = strconv.Itoa(int(*data))

	return saveCustom(path, dataInString, 4)
}

//SaveUint32 ...
func SaveUint32(path string, data *uint32) bool {

	preCondition(path)

	var dataInString string
	dataInString = strconv.Itoa(int(*data))

	return saveCustom(path, dataInString, 5)
}

//SaveInt32 ...
func SaveInt32(path string, data *int32) bool {

	preCondition(path)

	var dataInString string
	dataInString = strconv.Itoa(int(*data))

	return saveCustom(path, dataInString, 6)
}

//SaveUint64 ...
func SaveUint64(path string, data *uint64) bool {

	preCondition(path)

	var dataInString string
	dataInString = strconv.Itoa(int(*data))

	return saveCustom(path, dataInString, 7)
}


//SaveInt64 ...
func SaveInt64(path string, data *int64) bool {

	preCondition(path)

	var dataInString string
	dataInString = strconv.Itoa(int(*data))

	return saveCustom(path, dataInString, 8)
}

//SaveInt ...
func SaveInt(path string, data *int) bool {

	preCondition(path)

	var dataInString string
	dataInString = strconv.Itoa(int(*data))

	return saveCustom(path, dataInString, 9)
}

//SaveFloat32 ...
func SaveFloat32(path string, data *float32) bool {

	preCondition(path)

	var dataInString string
	dataInString = strconv.FormatFloat(float64(*data), 'E', -1, 32)

	return saveCustom(path, dataInString, 10)
}

//SaveFloat64 ...
func SaveFloat64(path string, data *float64) bool {

	preCondition(path)

	var dataInString string
	dataInString = strconv.FormatFloat(*data, 'E', -1, 32)

	return saveCustom(path, dataInString, 11)
}