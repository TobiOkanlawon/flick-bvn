package main

import (
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("does not accept invalid values", func (t *testing.T) {
		validator := Validator{}
		value := "jlsakdfjasdf"

		_, err := validator.Validate(value)

		if err == nil {
			t.Fatal("expected an error got none")
		}

		if err != ErrInvalidValue {
			t.Errorf("expected ErrInvalidValue error, got %s", err)
		}
	})

	t.Run("accepts valid BVN values", func (t *testing.T) {
		validator := Validator{}
		value := "123456789"

		_, err := validator.Validate(value)

		if err != nil {
			t.Fatal("expected no error")
		}
	})
}
