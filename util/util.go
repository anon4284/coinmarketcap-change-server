package util

import (
	"fmt"
	"strings"
)

func FormatString(str string) string {
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, "%", "", -1)
	str = strings.Replace(str, ")", "", -1)
	return strings.Replace(str, " ", "", -1)
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
