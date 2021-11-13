package main

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
)

type Int int

type Address struct {
	CityID Int    `json:"city_id" xml:"city_id"`
	Street string `json:"street" xml:"street"`
}

type User struct {
	ID      Int     `json:"id" xml:"id"`
	Address Address `json:"address" xml:"address"`
	Age     Int     `json:"age" xml:"age"`
}

type Users struct {
	Users []User `xml:"user"`
}

func (i *Int) UnmarshalJSON(b []byte) error {
	var item interface{}

	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}

	switch v := item.(type) {
	case int:
		*i = Int(v)

	case int32:
		*i = Int(v)

	case int64:
		*i = Int(v)

	case float32:
		*i = Int(v)

	case float64:
		*i = Int(v)

	case string:
		num, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*i = Int(num)
	}
	return nil
}

func (i *Int) UnmarshalXML(d *xml.Decoder, s xml.StartElement) error {
	var item string

	err := d.DecodeElement(&item, &s)
	if err != nil {
		return err
	}

	num, err := strconv.Atoi(item)
	if err != nil {
		return err
	}

	*i = Int(num)

	return nil
}
