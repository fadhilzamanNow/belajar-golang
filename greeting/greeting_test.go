package greeting

import (
	"regexp"
	"testing"
)

// TestHelloName calls greeting.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Gladys") = %q %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greeting.Hello with an empty string,
// checking for an empty return value.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q %v, want "", error`, msg, err)
	}
}
