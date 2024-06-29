package validator

import (
	"reflect"
	"strings"
)

// ValidateField validates field of a struct. Returns Error and true if the field is exists
func ValidateField(st any, field string, opts ...Option) (*Error, bool) {
	err := &Error{
		errs: map[string][]string{},
	}

	rv := reflect.ValueOf(st)

	sf, ok := reflect.Indirect(rv).Type().FieldByName(field)
	if !ok {
		return nil, false
	}

	fieldTagName := strings.TrimSuffix(sf.Tag.Get("json"), ",omitempty")
	if fieldTagName == "" {
		fieldTagName = field
	}

	for _, opt := range opts {
		if msg, ok := opt(rv.FieldByName(field)); !ok {
			err.errs[fieldTagName] = append(err.errs[fieldTagName], msg)
		}
	}

	if err.Success() {
		return nil, true
	}

	return err, true
}

func ValidateStruct(st any, fieldsOpts map[string][]Option) *Error {
	err := &Error{
		errs: map[string][]string{},
	}

	rv := reflect.ValueOf(st)

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i).Name

		fieldTagName := strings.TrimSuffix(rv.Type().Field(i).Tag.Get("json"), ",omitempty")
		if fieldTagName == "" {
			fieldTagName = field
		}

		for _, opt := range fieldsOpts[field] {
			if msg, ok := opt(rv.Field(i)); !ok {
				err.errs[fieldTagName] = append(err.errs[fieldTagName], msg)
			}
		}
	}

	return err
}
