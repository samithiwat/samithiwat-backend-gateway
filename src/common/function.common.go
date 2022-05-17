package common

import (
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

func FormatPathID(path string, id int32) string {
	if id > 0 {
		path = strings.Replace(path, strconv.Itoa(int(id)), ":id", 1)
	}

	return path
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
