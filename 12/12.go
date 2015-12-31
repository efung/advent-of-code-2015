package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var sum float64

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	check(err)

	var f interface{}
	json := json.NewDecoder(bufio.NewReader(file))
	json.Decode(&f)
	switch f.(type) {
	case map[string]interface{}:
		obj := f.(map[string]interface{})
		if ShouldDecodeObject(obj) {
			DecodeObject(obj)
		}
	case []interface{}:
		DecodeArray(f.([]interface{}))
	}

	fmt.Printf("Sum: %d\n", int(sum))
}

func DecodeObject(obj map[string]interface{}) {
	for _, v := range obj {
		switch vv := v.(type) {
		case float64:
			sum += vv
		case []interface{}:
			DecodeArray(vv)
		case map[string]interface{}:
			if ShouldDecodeObject(vv) {
				DecodeObject(vv)
			}
		}
	}
}

func DecodeArray(arr []interface{}) {
	for _, e := range arr {
		switch ee := e.(type) {
		case float64:
			sum += ee
		case []interface{}:
			DecodeArray(ee)
		case map[string]interface{}:
			if ShouldDecodeObject(ee) {
				DecodeObject(ee)
			}
		}
	}
}

func ShouldDecodeObject(obj map[string]interface{}) bool {
	// Uncomment below for Part 2 of problem
	for _, v := range obj {
		switch vv := v.(type) {
		case string:
			if vv == "red" {
				return false
			}
		}
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
