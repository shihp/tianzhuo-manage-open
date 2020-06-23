package test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_Time(t *testing.T) {

	s := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(s)
}
