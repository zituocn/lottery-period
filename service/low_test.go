package service

import (
	"fmt"
	"testing"
	"time"
)

func Test_GetOneByPeriod(t *testing.T) {
	low := new(Low)
	result, err := low.GetOneByPeriod(1, 2018281)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v \n", result)
}

func Test_GetOneByTime(t *testing.T) {
	low := new(Low)
	result, err := low.GetOneByTime(1, time.Now().Unix(), -1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v \n", result)
}
