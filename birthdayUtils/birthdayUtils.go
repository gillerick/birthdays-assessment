package birthdayUtils

import "time"

// CheckIfBirthday /*This method signature is for the purpose of enabling testing.
//We are deliberately passing the currentDate as opposed to instantiating it within the function
//The method returns a birthday in the []string format or an error
//*/
func CheckIfBirthday(dateOfBirth []string, currentDate time.Time) ([]string, error) {
	//Here, we are interested in the date of birth which is the last element in our array
	birthDate, err := time.Parse("2006/01/02", dateOfBirth[len(dateOfBirth)-1])

	if err != nil {
		return nil, err
	}

	currentDay := currentDate.Day()
	currentMonth := currentDate.Month()
	currentYear := currentDate.Year()

	//	3.1. Check if the birthday falls on 29th of February and perform a leap year check
	if checkIsLeapYear(currentYear) {
		if birthDate.Month() == time.February && birthDate.Day() == 29 && currentDay == 28 {
			return dateOfBirth, nil
		}
	} else {
		if currentMonth == birthDate.Month() && currentDay == birthDate.Day() {
			return dateOfBirth, nil
		} else {
			return nil, nil
		}
	}
	return nil, nil

}

/*This function name begins in lowercase for access control; we are only using it within this package*/
func checkIsLeapYear(year int) bool {
	return year%4 == 0
}
