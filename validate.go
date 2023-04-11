package validate

import (
	"fmt"
	"reflect" 
	"errors"
)

var ErrNotStruct = errors.New("wrong argument given, should be a struct")
var ErrInvalidValidatorSyntax = errors.New("invalid validator syntax")
var ErrValidateForUnexportedFields = errors.New("validation for unexported field is not allowed")


type ValidationError struct {
	Err error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var result string
	for _, i := range v {
		result += fmt.Sprint(i.Err, "\n")
	}
	return result
}

func Validate(v any) error {
	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Struct {
		 return ErrNotStruct
	}

	var errs ValidationErrors
	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get("json")

		if len(tag) == 0 {
			 continue
		}

		if !validateTag(tag) {
			errs = append(errs, ValidationError{Err: ErrInvalidValidatorSyntax})
			continue
		}

		if !val.Field(i).CanInterface() {
			errs = append(errs, ValidationError{Err: ErrValidateForUnexportedFields})
			continue
		}

		if !validateValue(val.Field(i), tag) {
			errs = append(errs, ValidationError{Err: ErrValidateForUnexportedFields})
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func validateTag(tag string) bool {
	fields := []string{"text", "title", "user_id", "id", "author_id", "published"}
	for _, field := range fields {
		if field == tag {
			return true
		}
	}
	return false
}

func validateValue(v any, tag string) bool {
	switch tag {
	case "title":
		lenTitle := len(fmt.Sprint(v))
		return 0 < lenTitle && lenTitle <= 100
	case "text":
		lenText := len(fmt.Sprint(v))
		return 0 < lenText && lenText <= 500
	case "id":
		return true
	case "author_id":
		return true
	case "user_id":
		return true
	case "published":
		return true
	}
	return false
}
