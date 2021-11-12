package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

var ErrNotPointer = errors.New("the argument is not a pointer")

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

func unmarshal(src map[string]interface{}, dst interface{}) {
	// v := reflect.ValueOf(dst).Elem()
	// if v.Kind() == reflect.Struct {
	// 	for i := 0; i < v.NumField(); i++ {
	// 		tag := v.Type().Field(i).Tag.Get("json")

	// 		if tag == "" || tag == "-" {
	// 			continue
	// 		}

	// 		if tag ==
	// 	}
	// }

	res := reflect.ValueOf(dst)

	typ := reflect.TypeOf(dst)
	fmt.Println(typ)

	fmt.Println("RES KIND IS", res.Kind())

	if res.Kind() == reflect.Ptr {
		res = res.Elem()
	}
	// source := reflect.ValueOf(src)

	// capacity := source.Len() + 1

	t := reflect.New(reflect.TypeOf(dst))

	fmt.Println(t.Type())

	v := reflect.ValueOf(t)

	// if v.Kind() == reflect.Ptr {
	// 	v = v.Elem()
	// }

	// res := reflect.ValueOf(dst)

	fmt.Println(v.Kind())

	if v.Kind() == reflect.Struct {
		fmt.Println("is struct")
		for i := 0; i < v.NumField(); i++ {
			tag := v.Type().Field(i).Tag.Get("json")
			if tag == "" || tag == "-" {
				continue
			}

			field := v.Field(i)
			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}

			value, exists := src[tag]
			if !exists {
				fmt.Println("DOES NOT EXIST")
				continue
			}

			if reflect.ValueOf(value).Kind() == reflect.Map {
				unmarshal(value.(map[string]interface{}), field.Addr().Addr())
				continue
			}

			if field.IsValid() {
				if field.CanSet() {
					if reflect.TypeOf(value) == nil {
						field.Set(reflect.Zero(field.Type()))
					} else {
						field.Set(reflect.ValueOf(value).Convert(field.Type()))
					}
				}
			}

		}
	}

	// res.Set(reflect.Append(res, v))
	// for i := 0; i < v.NumField(); i++ {
	// 	field := v.Field(i)
	// 	if field.Kind() == reflect.Ptr {
	// 		field = field.Elem()
	// 	}

	// 	if field.Kind() == reflect.Struct {
	// 		removeCyrillic(field.Addr().Interface())
	// 	} else if field.Kind() == reflect.String {

	// 		if field.IsValid() {

	// 			if field.CanSet() {
	// 				value := field.Interface()

	// 				newValue := strings.Map(
	// 					func(r rune) rune {
	// 						if unicode.Is(unicode.Cyrillic, r) {
	// 							return -1
	// 						}
	// 						return r
	// 					}, value.(string))

	// 				field.SetString(newValue)
	// 			}
	// 		}
	// 	}
	// }
}

func Unmarshal(data []byte, dst interface{}) error {
	// if dst is not a pointer, return the error
	if reflect.ValueOf(dst).Kind() != reflect.Ptr {
		return fmt.Errorf("%w", ErrNotPointer)
	}
	// create a map to unmarshal json into
	jsonMap := make([]map[string]interface{}, 0, 100)
	err := json.Unmarshal(data, &jsonMap)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	for i := range jsonMap {
		fmt.Println(jsonMap[i])
	}

	v := reflect.ValueOf(dst).Elem()

	typ := reflect.TypeOf(dst).Elem().Elem()

	t := reflect.MakeSlice(reflect.SliceOf(typ), 0, len(jsonMap))

	for i := range jsonMap {
		unmarshal(jsonMap[i], t)
		fmt.Println(dst)
	}
	// unmarshal(jsonMap, dst)
	reflect.Copy(v, t)
	return nil
}

type ola struct {
	E alo
	F string
}

type alo struct {
	A string
	B string
	C *string
	D string
}

var rawJson = []byte(`[
  {
    "id": 1,
    "address": {
      "city_id": 5,
      "street": "Satbayev"
    },
    "Age": 20
  },
  {
    "id": 1,
    "address": {
      "city_id": "6",
      "street": "Al-Farabi"
    },
    "Age": "32"
  }
]`)

type User struct {
	ID      int64   `json:"id"`
	Address Address `json:"address"`
	Age     int     `json:"age"`
}

type Address struct {
	CityID int64  `json:"city_id"`
	Street string `json:"street"`
}

func main() {
	c := "HELLO ЩЩЩЩЩССФЫЩ ahhahahha"
	temp := alo{
		A: "HELLO ЩЩЩЩЩССФЫЩ ahhahahha",
		B: "HELLO ЩЩЩЩЩССФЫЩ ahhahahha",
		C: &c,
		D: "HELLO ЩЩЩЩЩССФЫЩ ahhahahha",
	}

	temp1 := ola{
		E: temp,
		F: "ddo туттутрара лфыл askl",
	}

	err := RemoveCyrillic(&temp1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%-v\n", *temp1.E.C)

	var users []User
	if err := Unmarshal(rawJson, &users); err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}
}
