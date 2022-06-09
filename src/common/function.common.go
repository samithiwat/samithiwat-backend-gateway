package common

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IsExisted(e map[string]struct{}, key string) bool {
	_, ok := e[key]
	if ok {
		return true
	}
	return false
}

func FormatPath(method string, path string, id int32) string {
	if id > 0 {
		path = strings.Replace(path, strconv.Itoa(int(id)), ":id", 1)
	}

	return fmt.Sprintf("%v %v", method, path)
}

func FindIntFromStr(s string) []int32 {
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(s, -1)

	var result []int32

	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		result = append(result, int32(n))
	}

	return result
}
