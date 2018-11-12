package tools

import (
	"fmt"
	"strings"
)

func ArrToString(arr []string) string{
	str := strings.Trim(fmt.Sprint(arr), "[]")
	str = strings.Replace(str, " ", "", -1)
	return str
}