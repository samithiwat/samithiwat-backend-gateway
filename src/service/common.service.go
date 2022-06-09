package service

import "fmt"

func FormatErr(errors []string) string {
	result := ""
	if len(errors) > 0 {
		result = errors[0]
		for i := 1; i < len(errors); i++ {
			result = fmt.Sprintf("%v, %v", result, errors[i])
		}
	}
	return result
}
