package main

import (
	"birthday/birthday"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	//1. Read birthdays from source file (JSON)
	birthdaysByte, _ := ioutil.ReadFile("birthdays.json")

	type multiStringArray [][]string

	var birthdaysArray multiStringArray

	//2. Map the birthdays to the expected interface
	if err := json.Unmarshal(birthdaysByte, &birthdaysArray); err != nil {
		log.Printf("failed to map birthdays to interface: %v", err)
	}

	//Initialize empty array of type multiStringArray for storing today's birthdays
	todayBirthdays := make(multiStringArray, 0)

	//3. Loop through the mapped values and perform a birthDate check
	for _, birthDate := range birthdaysArray {
		dates, err := birthday.CheckIfBirthday(birthDate, time.Now())

		if err != nil {
			log.Printf("There was an error checking the birth date: %v", err)
		}

		//		3.2. Collect the current birthdays into an array
		if dates != nil {
			todayBirthdays = append(todayBirthdays, dates)
		}

	}

	//3.3. Print out all the birthdays
	fmt.Printf("Today's Birthdays: %s", todayBirthdays)
}
