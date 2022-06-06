package main

import (
	"encoding/json"
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
		currentDate := time.Now()
		dateOfBirth, err := time.Parse("2006/01/02", birthday[len(birthday)-1])

		if err != nil{
			panic(err)
		}

		if dateOfBirth.Day() == currentDate.Day() && dateOfBirth.Month() == currentDate.Month() {
			todayBirthdays = append(todayBirthdays, dateOfBirth)
		}
	}

}
