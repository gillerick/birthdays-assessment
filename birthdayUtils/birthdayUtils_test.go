package birthdayUtils

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckIfBirthday(t *testing.T) {
	type args struct {
		dateOfBirth []string
		currentDate time.Time
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"Valid date format: Non Leap Year Birthday", args{[]string{"John", "Doe", "1996/03/23"}, time.Date(2022, 03, 23, 03, 34, 45, 23, time.Local)}, []string{"John", "Doe", "1996/03/23"}, false},
		{"Invalid date format: Non Leap Year Birthday", args{[]string{"John", "Doe", "1996/23/03"}, time.Date(2022, 03, 23, 03, 34, 45, 23, time.Local)}, nil, true},
		{"Valid date format: Leap Year Birthday", args{[]string{"John", "Doe", "1996/02/29"}, time.Date(2020, 02, 28, 03, 34, 45, 23, time.Local)},[]string{"John", "Doe", "1996/02/29"} , false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckIfBirthday(tt.args.dateOfBirth, tt.args.currentDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckIfBirthday() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckIfBirthday() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkIsLeapYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Non Leap Year", args{1998}, false},
		{"Leap Year", args{2024}, true},
		{"Leap Year", args{2020}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIsLeapYear(tt.args.year); got != tt.want {
				t.Errorf("checkIsLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
