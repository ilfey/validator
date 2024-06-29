package validator

import (
	"fmt"
	"reflect"
)

// Option is a validator option. It returns true if the value is valid
type Option func(reflect.Value) (string, bool)

// Required returns true if the value is nil
func Required() Option {
	return func(v reflect.Value) (string, bool) {
		switch v.Kind() {
		case reflect.Ptr:
			return "Field \"%s\" is required", !v.IsNil()
		default:
			return "", true
		}
	}
}

func NegativeInt() Option {
	return func(v reflect.Value) (string, bool) {
		return "Field \"%s\" must be negative", v.Int() < 0
	}
}

func NegativeFloat() Option {
	return func(v reflect.Value) (string, bool) {
		return "Field \"%s\" must be negative", v.Float() < 0
	}
}

func PositiveInt() Option {
	return func(v reflect.Value) (string, bool) {
		return "Field \"%s\" must be positive", v.Int() > 0
	}
}

func PositiveFloat() Option {
	return func(v reflect.Value) (string, bool) {
		return "Field \"%s\" must be positive", v.Float() > 0
	}
}

func LessThanInt(num int64) Option {
	return func(v reflect.Value) (string, bool) {
		return "Field \"%s\" must be less than " + fmt.Sprint(num), v.Int() < num
	}
}

func LessThanFloat(num float64) Option {
	return func(v reflect.Value) (string, bool) {
		return "Field \"%s\" must be less than " + fmt.Sprint(num), v.Float() < num
	}
}

func GreaterThanInt(num int64) Option {
	return func(v reflect.Value) (string, bool) {
		return "Field \"%s\" must be greater than " + fmt.Sprint(num), v.Int() > num
	}
}

func GreaterThanFloat(num float64) Option {
	return func(v reflect.Value) (string, bool) {
		return "Field \"%s\" must be greater than " + fmt.Sprint(num), v.Float() > num
	}
}

func LenLessThan(num int) Option {
	return func(v reflect.Value) (string, bool) {
		return "Field \"%s\" must be less than " + fmt.Sprint(num), v.Len() < num
	}
}

func LenGreaterThan(num int) Option {
	return func(v reflect.Value) (string, bool) {
		return "Field \"%s\" must be greater than " + fmt.Sprint(num), v.Len() < num
	}
}
