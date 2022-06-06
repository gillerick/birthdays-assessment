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

	if currentMonth == birthDate.Month() && currentDay == birthDate.Day() {
		return dateOfBirth, nil
	} else {
		return nil, nil
	}
}
