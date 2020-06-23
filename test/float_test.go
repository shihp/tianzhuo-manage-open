package test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Float(t *testing.T) {
	//s := strconv.FormatFloat(3.8499506416584404, 'f', -1, 64)
	s := "100"
	fmt.Println(s)
	strArr := strings.Split(s, ".")
	if len(strArr) == 1 {
		println(strArr[0])
	} else {
		s1 := strArr[0] + "." + strArr[1][0:2]
		fmt.Println(s1)
	}
}
