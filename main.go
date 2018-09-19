package main

import (
	"time"
	"fmt"
)

func main() {

	layout := "2006/01/02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	str2 := "2017-12-13T19:21:04+09:00"
	t, err := time.Parse(layout, str)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	t2, err := time.Parse(time.RFC3339,str2)
	DateFormat := "2006/01/02"
	path := t2.Format(DateFormat)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t2)
	fmt.Println(path)

}
