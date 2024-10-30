package bvn

import (
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("does not accept invalid URL", func(t *testing.T) {
		_, err := newValidator("")

		if err == nil {
			t.Errorf("expected an error while passing in an invalid URL")
		}
	})

	t.Run("accepts a valid URL", func(t *testing.T) {
		_, err := newValidator("https://google.com")

		if err != nil {
			t.Errorf("did not expect and error while passing a valid URL %s", err)
		}
	})

	t.Run("does not accept invalid values", func(t *testing.T) {
		validator, _ := newValidator("https://google.com")

		value := "jlsakdfjasdf"

		_, err := validator.Validate(value)

		if err == nil {
			t.Fatal("expected an error got none")
		}

		if err != ErrInvalidValue {
			t.Errorf("expected ErrInvalidValue error, got %s", err)
		}
	})

	t.Run("accepts valid BVN values", func(t *testing.T) {
		validator, _ := newValidator("https://google.com")
		value := "123456789"

		_, err := validator.Validate(value)

		if err != nil {
			t.Fatal("expected no error")
		}
	})
}
