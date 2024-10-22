package main

import (
	"strconv"
)

type Validator struct{}

func (v *Validator) Validate(value string) (info Info, err error) {
	// Basically, you call the validate function with the BVN value and it hits the validation API

	_, err = strconv.Atoi(value)

	if err != nil {
		return Info{}, ErrInvalidValue
	}

	// TODO: Next step is to hit the Flick API to get the information for the thing

	return Info{}, nil
}
