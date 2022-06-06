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
		// TODO: Add test cases.
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
