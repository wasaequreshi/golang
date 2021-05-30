package greetings

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {
	name := "Wasae"
	want := regexp.MustCompile("\\b" + name + "\\b")
	msg, err := Hello("Wasae")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf("Hello(\"Wasae\" = %q, %v, want a match for %#q, nil", msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf("Hello(\"\") = %q, %v, want a match \"\" error", msg, err)
	}
}