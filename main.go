package main

import (
	"encoding/json"
	"io/ioutil"
)

func main()  {
	//1. Read birthdays from source file (JSON)
	birthdaysByte, _ := ioutil.ReadFile("birthdays.json")

	type multiStringArray [][]string

	var birthdaysArray multiStringArray

	if err := json.Unmarshal(birthdaysByte, &birthdaysArray); err != nil{
		panic(err)
	}
}
