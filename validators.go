package main

import (
	// "errors"
	"fmt"
	// "regexp"
	// "strconv"
	"strings"
	// "github.com/hashicorp/terraform/helper/resource"
	// "github.com/hashicorp/terraform/helper/schema"
	// "github.com/hashicorp/terraform/helper/structure"
	// "github.com/hashicorp/terraform/helper/validation"
)

func validateOnOff(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	acceptedVals := []string{"ON", "OFF"}

	if !stringInSlice(value, acceptedVals) {
		errors = append(errors, fmt.Errorf(
			"%q must be one of [%q]", k, strings.Join(acceptedVals, "|")))
	}

	return
}

func validateContainment(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	acceptedVals := []string{"NONE", "PARTIAL"}

	if !stringInSlice(value, acceptedVals) {
		errors = append(errors, fmt.Errorf(
			"%q must be one of [%q]", k, strings.Join(acceptedVals, "|")))
	}

	return
}

func validateCompatibilityLevel(v interface{}, k string) (ws []string, errors []error) {
	value := v.(int)
	acceptedVals := []int{90, 100, 110, 120, 130, 140, 150}

	if !intInSlice(value, acceptedVals) {
		errors = append(errors, fmt.Errorf(
			"%q must be one of [%q]", k, intSliceToStr(acceptedVals, "|")))
	}

	return
}
