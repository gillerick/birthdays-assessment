package main

import (
	"birthdayUtils/birthdayUtils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func main()  {
	//1. Read birthdays from source file (JSON)
	birthdaysByte, _ := ioutil.ReadFile("birthdays.json")

	type multiStringArray [][]string

	var birthdaysArray multiStringArray

	//2. Map the birthdays to the expected interface
	if err := json.Unmarshal(birthdaysByte, &birthdaysArray); err != nil{
		panic(err)
	}

	//Initialize empty array of type multiStringArray for storing today's birthdays
	todayBirthdays := make(multiStringArray, 0)

	//3. Loop through the mapped values and perform a birthday check
	for _, birthday := range birthdaysArray {
		dates, err := birthdayUtils.CheckIfBirthday(birthday, time.Now())

		if err != nil{
			panic(err)
		}

		todayBirthdays = append(todayBirthdays, dates)


	}

	fmt.Printf("Today's Birthdays: %s", todayBirthdays)
}
