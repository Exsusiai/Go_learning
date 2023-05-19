package greetings

import (
	"testing"
	"regexp"
	"fmt"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// TestHellosNames tests greetings.Hellos function with a slice of names.
func TestHellosNames(t *testing.T) {
	names := []string{"Alice", "Bob", "Eve"}
	messages, err := Hellos(names)

	if err != nil {
		t.Fatalf("Hellos(%v) encountered an error: %v", names, err)
	}

	if len(messages) != len(names) {
		t.Fatalf("Hellos(%v) returned a different number of messages than names, expected: %d, got: %d", names, len(names), len(messages))
	}

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	for _, name := range names {
		message, ok := messages[name]
		if !ok {
			t.Fatalf("Hellos(%v) did not return a greeting for %s", names, name)
		}

		matched := false
		for _, format := range formats {
			want := fmt.Sprintf(format, name)
			if message == want {
				matched = true
				break
			}
		}

		if !matched {
			t.Errorf("Hellos(%v) returned an incorrect greeting for %s: %q", names, name, message)
		}
	}
}

// TestHellosEmpty tests greetings.Hellos function with an empty slice.
func TestHellosEmpty(t *testing.T) {
	names := []string{}
	messages, err := Hellos(names)

	if err != nil {
		t.Fatalf("Hellos(%v) encountered an error: %v", names, err)
	}

	if len(messages) != 0 {
		t.Fatalf("Hellos(%v) expected to return empty messages, but got %d messages", names, len(messages))
	}
}

// TestHellosError tests greetings.Hellos function with a slice containing an empty name.
func TestHellosError(t *testing.T) {
	names := []string{"Alice", ""}
	_, err := Hellos(names)

	if err == nil {
		t.Fatalf("Hellos(%v) expected to return an error, but no error was returned", names)
	}
}