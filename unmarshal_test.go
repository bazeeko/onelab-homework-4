package main

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"testing"
)

func TestUnmarshalJson(t *testing.T) {
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

	var users []User

	expected := []User{
		{
			ID: 1,
			Address: Address{
				CityID: 5,
				Street: "Satbayev",
			},
			Age: 20,
		},
		{
			ID: 1,
			Address: Address{
				CityID: 6,
				Street: "Al-Farabi",
			},
			Age: 32,
		},
	}

	if err := json.Unmarshal(rawJson, &users); err != nil {
		t.Errorf("TestUnmarshalJson: %s", err)
	}

	if !reflect.DeepEqual(users, expected) {
		t.Errorf("TestUnmarshalXml: expected(%v), got(%v)", expected, users)
	}
}

func TestUnmarshalXml(t *testing.T) {
	var rawXml = []byte(`
<users>
  <user>
    <id>1</id>
    <address>
      <city_id>5</city_id>
      <street>Satbayev</street>
    </address>
    <age>20</age>
  </user>
  <user>
    <id>1</id>
    <address>
      <city_id>6</city_id>
      <street>Al-Farabi</street>
    </address>
    <age>32</age>
  </user>
</users>
`)

	var users Users

	expected := Users{
		Users: []User{{ID: 1,
			Address: Address{
				CityID: 5,
				Street: "Satbayev",
			},
			Age: 20}, {
			ID: 1,
			Address: Address{
				CityID: 6,
				Street: "Al-Farabi",
			},
			Age: 32,
		}},
	}

	if err := xml.Unmarshal(rawXml, &users); err != nil {
		t.Errorf("TestUnmarshalXml: %s", err)
	}

	if !reflect.DeepEqual(users, expected) {
		t.Errorf("TestUnmarshalXml: expected(%v), got(%v)", expected, users)
	}
}
