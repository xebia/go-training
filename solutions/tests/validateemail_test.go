package tests

import (
	"fmt"
	"testing"
)

// START OMIT
func TestValidateEmail(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: "", want: false},
		{input: "@", want: false},
		{input: "@xebia", want: false},
		{input: "m/grol@xebia..com", want: false},
		{input: "mgrol@xebia..com", want: false},
		{input: "mgrol@xebia", want: true},
		{input: "mgrol+dev@xebia", want: true},
		{input: "marc.grol@gmail.com", want: true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Testcase: %s", tc.input), func(t *testing.T) {
			got := IsValidEmailAddress(tc.input)
			if got != tc.want {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

// END OMIT

// TestIsValidEmailAddress boilerplate is created by the IDE
func TestIsValidEmailAddress(t *testing.T) {
	type args struct {
		emailAddress string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{emailAddress: ""}, want: false},
		{args: args{emailAddress: "@"}, want: false},
		{args: args{emailAddress: "@xebia"}, want: false},
		{args: args{emailAddress: "m/grol@xebia..com"}, want: false},
		{args: args{emailAddress: "mgrol@xebia..com"}, want: false},
		{args: args{emailAddress: "mgrol@xebia"}, want: true},
		{args: args{emailAddress: "mgrol+dev@xebia"}, want: true},
		{args: args{emailAddress: "marc.grol@gmail.com"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmailAddress(tt.args.emailAddress); got != tt.want {
				t.Errorf("IsValidEmailAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
