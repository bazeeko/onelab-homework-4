package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

var ErrNotPointer = errors.New("the argument is not a pointer")

type Ola struct {
	E Alo
	F string
}

type Alo struct {
	A string
	B string
	C *string
	D string
}

func removeCyrillic(src interface{}) {
	v := reflect.ValueOf(src)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}

			if field.Kind() == reflect.Struct {
				removeCyrillic(field.Addr().Interface())
			} else if field.Kind() == reflect.String {

				if field.IsValid() {

					if field.CanSet() {
						value := field.Interface()

						newValue := strings.Map(
							func(r rune) rune {
								if unicode.Is(unicode.Cyrillic, r) {
									return -1
								}
								return r
							}, value.(string))

						field.SetString(newValue)
					}
				}
			}
		}
	}
}

func RemoveCyrillic(src interface{}) error {
	if reflect.ValueOf(src).Kind() != reflect.Ptr {
		return fmt.Errorf("removeCyrillic: %w", ErrNotPointer)
	}

	removeCyrillic(src)

	return nil
}
