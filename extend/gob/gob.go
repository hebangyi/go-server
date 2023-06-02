package sGob

import (
	"encoding/gob"
	spool "go-server/extend/pool"
	"reflect"
)

func Decode(data []byte, params []reflect.Type) ([]reflect.Value, error) {
	buffer := spool.GetBuffer(data)
	decoder := gob.NewDecoder(buffer)

	defer spool.PutBuffer(buffer)

	valueList := make([]reflect.Value, len(params))
	for i, param := range params {
		newValue := reflect.New(param)
		err := decoder.DecodeValue(newValue)
		if err != nil {
			return nil, err
		}

		valueList[i] = newValue.Elem()
	}

	return valueList, nil
}

func DecodeFunc(data []byte, paramsType reflect.Type) ([]reflect.Value, error) {
	paramsLen := paramsType.NumIn()
	if paramsLen < 1 {
		return nil, nil
	}

	buffer := spool.GetBuffer(data)
	decoder := gob.NewDecoder(buffer)

	defer spool.PutBuffer(buffer)

	valueList := make([]reflect.Value, paramsLen)
	for i := 0; i < paramsLen; i++ {
		params := reflect.New(paramsType.In(i))
		err := decoder.DecodeValue(params)
		if err != nil {
			return nil, err
		}

		valueList[i] = params.Elem()
	}

	return valueList, nil
}

func Encode(values ...interface{}) ([]byte, error) {
	buffer := spool.GetBuffer(nil)
	encoder := gob.NewEncoder(buffer)

	defer spool.PutBuffer(buffer)

	var err error
	for _, value := range values {
		err = encoder.Encode(value)
		if err != nil {
			return nil, err
		}
	}

	data := buffer.Bytes()
	return data, nil
}
