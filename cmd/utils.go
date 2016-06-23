package cmd

import (
	"reflect"
	"strings"
)

func sanitizeResponse(response string) string {
	response = strings.Replace(response, "\r", "", -1)
	response = strings.Replace(response, "\n", "", -1)
	response = strings.Replace(response, "\\", "", -1)
	response = strings.Replace(response, " ", "", -1)
	return response
}

func checkSlicesContainValue(slices []string, value string) bool {
	set := make(map[string]bool)
	for _, v := range slices {
		set[v] = true
	}
	return set[value]
}

func compareUnknownObject(obj1, obj2 interface{}) bool {
	return reflect.DeepEqual(obj1, obj2)
}
