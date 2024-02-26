package main

import (
	"errors"
	"fmt"
	"reflect"
)

func Pop(arr interface{}, flag bool) (interface{}, error) {

	var res interface{}

	arrVal := reflect.ValueOf(arr)
	if arrVal.Kind() != reflect.Slice{

		return res,errors.New("arr is not a slice")
	}

	if slice,ok := arr.([]interface{}); ok {
		if flag {
			res = slice[1:]
		}else{
			res = slice[: len(slice)-1 ]
		}
	}else if sliceOfSlice, ok := arr.([][]interface{}); ok {

		if flag {
			res = sliceOfSlice[1:]
		}else{
			res = sliceOfSlice[:len(sliceOfSlice)-1]
		}
	}

	return res,nil
}

func main() {
	flag := true
	//arr := [][]interface{}{{1,2},{2,3},{3,4}}
	//arr := []interface{}{1,2,3}
	arr:="abc"
	fmt.Println( Pop(arr,flag) )
}
