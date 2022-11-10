package validator

import (
	"fmt"
	reflect "reflect"

	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
)

type ValidateErrors []ValidateError

func (ve ValidateErrors) Error() string {
	ret := ""
	for _, v := range ve {
		ret += fmt.Sprintf("%v\n", v.Error())
	}

	return ret
}

type ValidateError struct {
	name     string
	orgError v10.FieldError
}

func (fe *ValidateError) Tag() string {
	return fe.orgError.Tag()
}
func (fe *ValidateError) ActualTag() string {
	return fe.ActualTag()
}
func (fe *ValidateError) Namespace() string {
	return fe.name
}
func (fe *ValidateError) StructNamespace() string {
	return fe.orgError.StructNamespace()
}
func (fe *ValidateError) Field() string {
	return fe.name
}
func (fe *ValidateError) StructField() string {
	return fe.orgError.StructField()
}
func (fe *ValidateError) Value() interface{} {
	return fe.orgError.Value()
}
func (fe *ValidateError) Param() string {
	return fe.orgError.Param()
}
func (fe *ValidateError) Kind() reflect.Kind {

	return fe.orgError.Kind()
}
func (fe *ValidateError) Type() reflect.Type {
	return fe.orgError.Type()
}
func (fe *ValidateError) Translate(ut ut.Translator) string {
	return fe.orgError.Translate(ut)
}
func (fe *ValidateError) Error() string {
	return fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", fe.Field(), fe.Tag())
}

var _ v10.FieldError = new(ValidateError)

func WrapValidatorError(name string, err error) error {
	orgErrors, ok := err.(v10.ValidationErrors)

	if !ok {
		return err
	}

	var ves ValidateErrors
	for _, v := range orgErrors {
		ves = append(ves, ValidateError{name, v})
	}

	return ves
}

func DoValidate(v interface{}, name string) error {
	val, ok := v.(interface{ Validate() error })

	if !ok {
		return nil
	}

	if err := val.Validate(); err != nil {
		return WrapValidatorError(name, err)
	}

	return nil
}
