package bvn

import (
	"strconv"
	"net/http"
	"log"
	"net/url"
	"errors"
)

type Validator struct{
	URL string
}

func newValidator(uri string) (Validator, error) {
	_, err := url.ParseRequestURI(uri)

	if err != nil {
		return Validator{}, errors.New("passed an illegal URI")
	}
	
	return Validator{URL: uri}, nil
}

func (v *Validator) Validate(value string) (info Info, err error) {
	// Basically, you call the validate function with the BVN value and it hits the validation API

	_, err = strconv.Atoi(value)

	if err != nil {
		return Info{}, ErrInvalidValue
	}

	_, err = http.Get(v.URL)
	
	if err != nil {
		log.Fatalln(err)
	}

	// TODO: Next step is to hit the Flick API to get the information for the thing

	return Info{}, nil
}
